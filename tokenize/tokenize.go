package tokenize

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

var reserves string
var userInput []rune

var l int  // line no        行番号
var lp int // line Position  行での位置
var wp int // whole Position 全体での位置

func init() {
	reserves = "{},:"
}

func current() rune {
	return userInput[wp]
}

func wpAdvance() int {
	wp++
	return wp
}

func wpAdvanceN(n int) int {
	wp += n
	return wp
}

func lpAdvance() int {
	lp++
	return lp
}

func lpAdvanceN(n int) int {
	lp += n
	return lp
}

func Tokenize(text string) *Token {
	// initialization
	userInput = []rune(text)
	l = 1
	wp = 0
	lp = 0
	var head Token // 仮の一番最初
	cur := &head   // 現在着目しているトークン

	for wp < len(userInput) {
		// newline
		if '\n' == current() {
			l++         // 行を移動する
			wpAdvance() // 全体としては一つ進んでいるので+する
			lp = 0      // 行を移動するので行番号は0になる
			continue
		}

		// white space
		if ' ' == current() || '\t' == current() {
			wpAdvance()
			lpAdvance()
			continue
		}

		// reserved
		if strings.ContainsRune(reserves, current()) {
			cur = NewReservedToken(cur, NewPosition(l, wp, lp), string(current()))
			wpAdvance()
			lpAdvance()
			continue
		}

		// true
		if 't' == current() {
			cur = NewIdentToken(cur, NewPosition(l, wp, lp), "true")
			wpAdvanceN(4)
			lpAdvanceN(4)
			continue
		}

		// false
		if 'f' == current() {
			cur = NewIdentToken(cur, NewPosition(l, wp, lp), "false")
			wpAdvanceN(5)
			lpAdvanceN(5)
			continue
		}

		// null
		if 'n' == current() {
			cur = NewIdentToken(cur, NewPosition(l, wp, lp), "null")
			wpAdvanceN(4)
			lpAdvanceN(4)
			continue
		}

		// string
		if '"' == current() {
			var str string
			pos := NewPosition(l, wp, lp)

			// "
			str += "\""
			wpAdvance()
			lpAdvance()

			for '"' != current() {
				// escaped double quotation
				if '\\' == current() && '"' == userInput[wp+1] {
					str += "\\\""
					wpAdvanceN(2)
					lpAdvanceN(2)
					continue
				}

				str += string(current())
				wpAdvance()
				lpAdvance()
			}

			// "
			str += "\""
			wpAdvance()
			lpAdvance()

			cur = NewStringToken(cur, pos, str)
			continue
		}

		// number
		if unicode.IsDigit(current()) || '-' == current() {
			var num string
			pos := NewPosition(l, wp, lp)
			for unicode.IsDigit(current()) || '-' == current() || '.' == current() {
				num += string(current())
				wpAdvance()
				lpAdvance()
			}
			// try parse
			f, err := strconv.ParseFloat(num, 64)
			if err != nil {
				log.Fatalf("[%d行目の%d文字目] 数字をパースできませんでした: %s", l, lp, num)
			}
			cur = NewNumberToken(cur, pos, f)
			continue
		}

		log.Fatalf("[%d行目の%d文字目] 予期しない文字を発見しました: %s", l, lp, string(current()))
	}

	// 最後だというトークンを追加してあげる
	NewEOFToken(cur, NewPosition(l, wp, lp))
	return head.Next
}
