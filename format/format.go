package format

import (
	"fmt"
	"github.com/x0y14/jsonfmt/parse"
	"log"
	"strings"
)

var deps int // 深度
var configuration *Config

func gen(node *parse.Node) string {
	switch node.Kind {
	case parse.NdKV:
		s := strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += gen(node.Key)
		s += ": "
		s += gen(node.Value)
		return s
	case parse.NdString:
		return node.Str
	case parse.NdNumber:
		return fmt.Sprintf("%v", node.Number)
	case parse.NdTrue:
		return "true"
	case parse.NdFalse:
		return "false"
	case parse.NdNULL:
		return "null"
	case parse.NdObject:
		s := strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += "{"
		// 改行をつけてあげる
		if len(node.Children) >= 1 {
			deps++
			defer func() {
				s += "\n"
				deps--
			}()
			s += "\n"
		}
		for i, kv := range node.Children {
			s += gen(kv)
			// if is not last one
			if i != len(node.Children)-1 {
				s += ", "
				s += "\n"
			}
		}
		s += strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += "}"
		return s
	case parse.NdArray:
		s := strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += "["
		// 要素があれば
		if len(node.Children) >= 1 {
			deps++
			defer func() {
				s += "\n"
				deps--
			}()
			s += "\n"
		}
		for i, item := range node.Children {
			s += gen(item)
			// if is not last one
			if i != len(node.Children)-1 {
				s += ", "
				s += "\n"
			}
		}
		s += strings.Repeat(strings.Repeat(" ", configuration.Indent), deps)
		s += "]"
		return s
	}

	log.Fatalf("予期せぬノードを発見しました: NodeKind(%d)", node.Kind)
	return ""
}

func Format(config *Config, node *parse.Node) string {
	deps = 0
	configuration = config
	return gen(node)
}
