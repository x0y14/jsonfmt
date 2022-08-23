package parse

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/jsonfmt/tokenize"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *Node
	}{
		{
			"empty",
			"{}",
			NewObjectNode(nil),
		},
		{
			"string value",
			`{"key": "value"}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewStringNode("\"value\"")),
			}),
		},
		{
			"number without dot value",
			`{"key": 123}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(123)),
			}),
		},
		{
			"number with dot value",
			`{"key": 123.4}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(123.4)),
			}),
		},
		{
			"minus number without dot value",
			`{"key": -123}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(-123)),
			}),
		},
		{
			"minus number with dot value",
			`{"key": -123.4}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(-123.4)),
			}),
		},
		{
			"number dot zero value",
			`{"key": 123.0}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(123.0)),
			}),
		},
		{
			"minus number dot zero value",
			`{"key": -123.0}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewNumberNode(-123.0)),
			}),
		},
		{
			"two string value pairs",
			`{"key1": "value1", "key2": "value2"}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key1\""), NewStringNode("\"value1\"")),
				NewKVNode(NewStringNode("\"key2\""), NewStringNode("\"value2\"")),
			}),
		},
		{
			"array value",
			`{"key": [123, "value", {"key in array": "value in array"}, [-123]]}`,
			NewObjectNode([]*Node{
				NewKVNode(NewStringNode("\"key\""), NewArrayNode([]*Node{
					NewNumberNode(123),
					NewStringNode("\"value\""),
					NewObjectNode([]*Node{
						NewKVNode(NewStringNode("\"key in array\""), NewStringNode("\"value in array\"")),
					}),
					NewArrayNode([]*Node{
						NewNumberNode(-123),
					}),
				})),
			}),
		},
		{
			"toplevel array",
			`[{}, [{}], [], [[]]]`,
			NewArrayNode([]*Node{
				NewObjectNode(nil),
				NewArrayNode([]*Node{
					NewObjectNode(nil),
				}),
				NewArrayNode(nil),
				NewArrayNode([]*Node{
					NewArrayNode(nil),
				}),
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := tokenize.Tokenize(tt.in)
			node := Parse(tok)
			assert.Equal(t, tt.expect, node)
		})
	}
}
