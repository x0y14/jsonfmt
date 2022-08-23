package tokenize

type Token struct {
	kind     TokenKind
	position *Position

	num float64 // numberの時に使う
	str string  // string, reservedの時に使う

	next *Token
}

func NewToken(kind TokenKind, position *Position, num float64, str string) *Token {
	return &Token{
		kind:     kind,
		position: position,
		num:      num,
		str:      str,
	}
}

func NewEOFToken(cur *Token, position *Position) *Token {
	tok := NewToken(TkEOF, position, 0, "")
	cur.next = tok
	return tok
}

func NewReservedToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkReserved, position, 0, str)
	cur.next = tok
	return tok
}

func NewIdentToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkIdent, position, 0, str)
	cur.next = tok
	return tok
}

func NewNumberToken(cur *Token, position *Position, num float64) *Token {
	tok := NewToken(TkNumber, position, num, "")
	cur.next = tok
	return tok
}

func NewStringToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkString, position, 0, str)
	cur.next = tok
	return tok
}
