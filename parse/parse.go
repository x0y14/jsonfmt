package parse

import (
	"github.com/x0y14/jsonfmt/tokenize"
	"log"
)

// 着目しているトークン
var token *tokenize.Token

// 現在着目しているトークンが記号でないとおかしい場合に使用、正しければ読み進める
func expectReserved(symbol string) {
	if token.Kind != tokenize.TkReserved {
		log.Fatalf("[%d行目の%d文字目] 記号ではありません: TokenKind(%d)", token.Position.LineNo, token.Position.LpBegin, token.Kind)
	}
	if symbol != token.Str {
		log.Fatalf("[%d行目の%d文字目] 想定された記号ではありません: %s", token.Position.LineNo, token.Position.LpBegin, token.Str)
	}
	token = token.Next
}

// 現在着目しているトークンが文字列出ないとおかしい場合に使用、正しければ読み進め、トークンを返却する
func expectString() *tokenize.Token {
	if token.Kind != tokenize.TkString {
		log.Fatalf("[%d行目の%d文字目] 文字列ではありません: TokenKind(%d)", token.Position.LineNo, token.Position.LpBegin, token.Kind)
	}
	tok := token
	token = token.Next
	return tok
}

// 現在着目しているトークンが一致する記号であればトークンを読み進め、読んだトークンを返却する
func consumeReserved(symbol string) *tokenize.Token {
	if token.Kind != tokenize.TkReserved || token.Str != symbol {
		return nil
	}

	tok := token
	token = token.Next
	return tok
}

// 現在着目しているトークンが一致する識別子であれば読み進め、読んだトークンを返却する
func consumeIdent(ident string) *tokenize.Token {
	if token.Kind != tokenize.TkIdent || token.Str != ident {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

// 現在着目しているトークンが数字であれば読み進め、読んだトークンを返却する
func consumeNumber() *tokenize.Token {
	if token.Kind != tokenize.TkNumber {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

// 現在着目しているトークンが文字列であれば読み進め、読んだトークンを返却する
func consumeString() *tokenize.Token {
	if token.Kind != tokenize.TkString {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func kv() *Node {
	key := expectString()
	expectReserved(":")
	return NewKVNode(NewStringNode(key.Str), value())
}

func value() *Node {
	// string
	if v := consumeString(); v != nil {
		return NewStringNode(v.Str)
	}
	// number
	if v := consumeNumber(); v != nil {
		return NewNumberNode(v.Num)
	}
	// true
	if v := consumeIdent("true"); v != nil {
		return NewTrueNode()
	}
	// false
	if v := consumeIdent("false"); v != nil {
		return NewFalseNode()
	}
	// null
	if v := consumeIdent("null"); v != nil {
		return NewNullNode()
	}
	// object
	if v := consumeReserved("{"); v != nil {
		return object()
	}
	// array
	if v := consumeReserved("["); v != nil {
		return array()
	}

	log.Fatalf("[%d行目の%d文字目] 予期せぬトークンを発見しました: TokenKind(%d)", token.Position.LineNo, token.Position.LpBegin, token.Kind)
	return nil
}

func object() *Node {
	var children []*Node
	// nilであれば、別の種類のトークンに着目している
	for consumeReserved("}") == nil {
		children = append(children, kv())
		_ = consumeReserved(",")
	}
	return NewObjectNode(children)
}

func array() *Node {
	var children []*Node
	for consumeReserved("]") == nil {
		children = append(children, value())
		_ = consumeReserved(",")
	}
	return NewArrayNode(children)
}

func Parse(tok *tokenize.Token) *Node {
	token = tok
	// toplevel object
	if v := consumeReserved("{"); v != nil {
		return object()
	}
	// toplevel array
	if v := consumeReserved("["); v != nil {
		return array()
	}
	log.Fatalf("[%d行目の%d文字目] 予期せぬトークンを発見しました: TokenKind(%d)", token.Position.LineNo, token.Position.LpBegin, token.Kind)
	return nil
}
