package main

import (
	"flag"
	"github.com/x0y14/jsonfmt"
	"github.com/x0y14/jsonfmt/format"
	"log"
	"os"
)

var overwrite bool

const overwriteDesc = "Whether to overwrite the original data with formatted data.\n" +
	"Defaults to false, which only displays the formatted data."

var outputFilePath string

const outputFilePathDesc = "If set, formatted data will be output on this path (0644).\n" +
	"It cannot be used with the Overwrite setting, which takes precedence.\n" +
	"If the file already exists in the path, an error is generated."

var indent int

const indentDesc = "Specifies the base indentation for outputting json. Default is 2."

var printToTerminal bool

const printToTerminalDesc = "Output to terminal, if both overwrite and output are not set, this is enabled."

func init() {

	flag.BoolVar(&overwrite, "overwrite", false, overwriteDesc)
	flag.BoolVar(&overwrite, "w", false, "shorthand of overwrite")

	flag.StringVar(&outputFilePath, "output", "", outputFilePathDesc)
	flag.StringVar(&outputFilePath, "o", "", "shorthand of output")

	flag.IntVar(&indent, "indent", 2, indentDesc)
	flag.IntVar(&indent, "i", 2, "shorthand of indent")

	flag.BoolVar(&printToTerminal, "print", true, printToTerminalDesc)
	flag.BoolVar(&printToTerminal, "p", true, "shorthand of print")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("整形対象とするファイルが設定されていません")
	}

	// 上書き許可と出力先が同時に設定されていたら、安全のため、上書きをオフにして、出力を有効にする
	if overwrite && outputFilePath != "" {
		overwrite = false
	}

	// 出力先が設定されていたらプリント機能をオフに
	if outputFilePath != "" {
		printToTerminal = false
	}

	// 上書き設定されていたらプリント機能をオフに
	if overwrite {
		printToTerminal = false
	}

	config := &jsonfmt.Config{
		Overwrite:        overwrite,
		OriginalFilePath: args[0],
		OutputFilePath:   outputFilePath,
		FormatterConfig: &format.Config{
			Indent: indent,
		},
		PrintToTerminal: printToTerminal,
	}

	f, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalf("ファイルの読み込みに失敗しました: %s", err)
	}

	jsonfmt.Formatting(config, string(f))
}
