package jsonfmt

import (
	"fmt"
	"github.com/x0y14/jsonfmt/format"
	"github.com/x0y14/jsonfmt/parse"
	"github.com/x0y14/jsonfmt/tokenize"
	"log"
	"os"
)

func Formatting(config *Config, json string) {
	token := tokenize.Tokenize(json)
	node := parse.Parse(token)
	result := format.Format(config.FormatterConfig, node)

	// ターミナルに出力する
	if config.PrintToTerminal {
		fmt.Println(result)
		return
	}

	// ファイルに出力
	if config.OutputFilePath != "" {
		_, err := os.Stat(config.OutputFilePath)
		if err == nil {
			log.Fatalf("ファイルがすでに存在します: %s", config.OutputFilePath)
		}
		err = os.WriteFile(config.OutputFilePath, []byte(result), 0644)
		if err != nil {
			log.Fatalf("書き込みに失敗しました: %s", err)
		}
		return
	}

	if config.Overwrite {
		f, err := os.OpenFile(config.OriginalFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("ファイルを開けませんでした: %s", err)
		}

		_, err = f.WriteString(result)
		if err != nil {
			log.Fatalf("書き込みに失敗しました: %s", err)
		}

		if err = f.Close(); err != nil {
			log.Fatalf("ファイルを閉じれませんでした: %s", err)
		}
	}
}
