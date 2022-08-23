package tokenize

type Position struct {
	LineNo  int
	WpBegin int
	LpBegin int
}

func NewPosition(line, wpBegin, lpBegin int) *Position {
	return &Position{
		LineNo:  line,
		WpBegin: wpBegin,
		LpBegin: lpBegin,
	}
}
