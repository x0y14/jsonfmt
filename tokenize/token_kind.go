package tokenize

type TokenKind int

const (
	TkIllegal TokenKind = iota
	TkEOF

	TkReserved // 記号
	TkIdent    // true, false, null
	TkNumber   // 数字
	TkString   // "abc"
)
