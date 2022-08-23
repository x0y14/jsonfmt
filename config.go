package jsonfmt

import "github.com/x0y14/jsonfmt/format"

type Config struct {
	Overwrite        bool
	OriginalFilePath string
	OutputFilePath   string
	FormatterConfig  *format.Config
	PrintToTerminal  bool
}
