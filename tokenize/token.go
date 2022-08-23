package tokenize

type Token struct {
	Kind     TokenKind
	Position *Position

	Num float64 // numberの時に使う
	Str string  // string, reservedの時に使う

	Next *Token
}

func NewToken(kind TokenKind, position *Position, num float64, str string) *Token {
	return &Token{
		Kind:     kind,
		Position: position,
		Num:      num,
		Str:      str,
	}
}

func NewEOFToken(cur *Token, position *Position) *Token {
	tok := NewToken(TkEOF, position, 0, "")
	cur.Next = tok
	return tok
}

func NewReservedToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkReserved, position, 0, str)
	cur.Next = tok
	return tok
}

func NewIdentToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkIdent, position, 0, str)
	cur.Next = tok
	return tok
}

func NewNumberToken(cur *Token, position *Position, num float64) *Token {
	tok := NewToken(TkNumber, position, num, "")
	cur.Next = tok
	return tok
}

func NewStringToken(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkString, position, 0, str)
	cur.Next = tok
	return tok
}
