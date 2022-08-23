package format

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/jsonfmt/parse"
	"github.com/x0y14/jsonfmt/tokenize"
	"testing"
)

var config *Config

func init() {
	config = &Config{
		Indent: 2,
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect string
	}{
		{
			"empty",
			"{}",
			"{}",
		},
		{
			"kv",
			`{"key":"value"}`,
			`{"key": "value"}`,
		},
		{
			"kv in kv",
			"{\"key\":{\"k\":\"v\"}}",
			"{\"key\": {\"k\": \"v\"}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := tokenize.Tokenize(tt.in)
			node := parse.Parse(token)
			out := Format(config, node)
			assert.Equal(t, tt.expect, out)
		})
	}
}
