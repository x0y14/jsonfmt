package format

import (
	"fmt"
	"github.com/x0y14/jsonfmt/parse"
	"strings"
)

var lp int   // 行中位置
var deps int // 深さ
var configuration *Config

func gen(node *parse.Node) string {
	var s string
	switch node.Kind {
	case parse.NdKV:
		// 深さ分インデントをつけてあげる
		s += strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += gen(node.Key)
		s += ": "
		s += gen(node.Value)
	case parse.NdString:
		s += node.Str
	case parse.NdNumber:
		s += fmt.Sprintf("%v", node.Number)
	case parse.NdTrue:
		s += "true"
	case parse.NdFalse:
		s += "false"
	case parse.NdNULL:
		s += "null"
	case parse.NdObject, parse.NdArray:
		// もしlpがゼロなら、行の先頭であるので、インデントをつけてあげる
		if lp == 0 {
			s += strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		}

		if node.Kind == parse.NdObject {
			s += "{"
		} else {
			s += "["
		}

		// 子要素の数が１以上であれば、深さを１追加
		// そうでなければ余分な空白が生まれてしまうので無視
		// {     } みたいな
		incremented := false
		if len(node.Children) >= 1 {
			deps++
			s += "\n"
			lp = 0 // 改行したら行中位置を0に
			defer func() {
				s += "\n"
				lp = 0 // 改行したら行中位置を0に
				// defer内でdeps--すると}の位置がずれる
			}()
			incremented = true // depsがマイナスにならないように増加させたよ通知
		}

		// 子要素出力
		for i, item := range node.Children {
			s += gen(item)
			// 最後の要素でなければ,を出力
			if i != len(node.Children)-1 {
				s += ", "
			}
			s += "\n"
			lp = 0
		}

		if incremented {
			deps--
		}

		s += strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		if node.Kind == parse.NdObject {
			s += "}"
		} else {
			s += "]"
		}
	}

	// 行中位置を更新してあげる
	lp = len(s)
	return s
}

func Format(config *Config, node *parse.Node) string {
	configuration = config
	return gen(node)
}
