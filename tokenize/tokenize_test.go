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
				kind:     TkReserved,
				position: NewPosition(1, 0, 0),
				num:      0,
				str:      "{",
				next: &Token{
					kind:     TkString,
					position: NewPosition(1, 1, 1),
					num:      0,
					str:      "\"key\"",
					next: &Token{
						kind:     TkReserved,
						position: NewPosition(1, 1+len("\"key\""), 1+len("\"key\"")),
						num:      0,
						str:      ":",
						next: &Token{
							kind:     TkString,
							position: NewPosition(1, 1+len("\"key\"")+1, 1+len("\"key\"")+1),
							num:      0,
							str:      "\"value\"",
							next: &Token{
								kind:     TkReserved,
								position: NewPosition(1, 1+len("\"key\"")+1+len("\"value\""), 1+len("\"key\"")+1+len("\"value\"")),
								num:      0,
								str:      "}",
								next: &Token{
									kind:     TkEOF,
									position: NewPosition(1, 1+len("\"key\"")+1+len("\"value\"")+1, 1+len("\"key\"")+1+len("\"value\"")+1),
									num:      0,
									str:      "",
									next:     nil,
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
				kind:     TkReserved,
				position: NewPosition(1, 0, 0),
				num:      0,
				str:      "{",
				next: &Token{
					kind:     TkString,
					position: NewPosition(1, 1, 1),
					num:      0,
					str:      "\"key\"",
					next: &Token{
						kind:     TkReserved,
						position: NewPosition(1, 1+len("\"key\""), 1+len("\"key\"")),
						num:      0,
						str:      ":",
						next: &Token{
							kind:     TkNumber,
							position: NewPosition(1, 1+len("\"key\"")+1, 1+len("\"key\"")+1),
							num:      123,
							str:      "",
							next: &Token{
								kind:     TkReserved,
								position: NewPosition(1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123)), 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))),
								num:      0,
								str:      "}",
								next: &Token{
									kind:     TkEOF,
									position: NewPosition(1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))+1, 1+len("\"key\"")+1+len(fmt.Sprintf("%v", 123))+1),
									num:      0,
									str:      "",
									next:     nil,
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
