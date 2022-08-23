package tokenize

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *Token
	}{
		{
			"empty object",
			"{}",
			&Token{
				TkReserved,
				NewPosition(1, 0, 0),
				0,
				"{",
				&Token{
					TkReserved,
					NewPosition(1, 1, 1),
					0,
					"}",
					&Token{
						TkEOF,
						NewPosition(1, 2, 2),
						0,
						"",
						nil,
					},
				},
			},
		},
		{
			"string value",
			`{"key":"value"}`,
			&Token{
				Kind:     TkReserved,
				Position: NewPosition(1, 0, 0),
				Num:      0,
				Str:      "{",
				Next: &Token{
					Kind:     TkString,
					Position: NewPosition(1, 1, 1),
					Num:      0,
					Str:      "\"key\"",
					Next: &Token{
						Kind:     TkReserved,
						Position: NewPosition(1, 1+len("\"key\""), 1+len("\"key\"")),
						Num:      0,
						Str:      ":",
						Next: &Token{
							Kind:     TkString,
							Position: NewPosition(1, 1+len("\"key\"")+1, 1+len("\"key\"")+1),
							Num:      0,
							Str:      "\"value\"",
							Next: &Token{
								Kind:     TkReserved,
								Position: NewPosition(1, 1+len("\"key\"")+1+len("\"value\""), 1+len("\"key\"")+1+len("\"value\"")),
								Num:      0,
								Str:      "}",
								Next: &Token{
									Kind:     TkEOF,
									Position: NewPosition(1, 1+len("\"key\"")+1+len("\"value\"")+1, 1+len("\"key\"")+1+len("\"value\"")+1),
									Num:      0,
									Str:      "",
									Next:     nil,
								},
							},
						},
					},
				},
			},
		},
		{
			"float value",
			`{"key":123}`,
			&Token{
				Kind:     TkReserved,
				Position: NewPosition(1, 0, 0),
				Num:      0,
				Str:      "{",
				Next: &Token{
					Kind:     TkString,
					Position: NewPosition(1, 1, 1),
					Num:      0,
					Str:      "\"key\"",
					Next: &Token{
						Kind:     TkReserved,
						Position: NewPosition(1, 1+len("\"key\""), 1+len("\"key\"")),
						Num:      0,
						Str:      ":",
						Next: &Token{
							Kind:     TkNumber,
							Position: NewPosition(1, 1+len("\"key\"")+1, 1+len("\"key\"")+1),
							Num:      123,
							Str:      "",
							Next: &Token{
								Kind:     TkReserved,
								Position: NewPosition(1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123)), 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))),
								Num:      0,
								Str:      "}",
								Next: &Token{
									Kind:     TkEOF,
									Position: NewPosition(1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))+1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))+1),
									Num:      0,
									Str:      "",
									Next:     nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := Tokenize(tt.in)
			assert.Equal(t, tt.expect, tok)
		})
	}
}
