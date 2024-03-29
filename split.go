package gorean

import (
	"bytes"
	"fmt"
)

type SplitOpt int16

const (
	DecimalWhitespace        = 32
	DecimalExclamationMark   = 33
	DecimalSlash             = 47
	DecimalColon             = 58
	DecimalAt                = 64
	DecimalLeftSquareBracket = 91
	DecimalGraveAccent       = 96
	DecimalLeftCurlyBracket  = 123
	DecimalTilde             = 126

	SplitOptBasic SplitOpt = iota
	SplitOptGetOnlyKorean
)

// SplitToString disassemble sentence with string.
// if ignoreSeparators is true, it'll be ignored whitespaces. e.g. "Apple phone " => "Applephone"
// if ignoreCases is true, it'll be ignored cases. e.g. Apple => apple
// if ignoreSpecialChars is true, it'll be ignored special characters. e.g. !@#$%^&*() => ""
func SplitToString(str string, opt SplitOpt, ignoreSeparators, ignoreCases, ignoreSpecialChars bool) (string, error) {
	runes := []rune(str)
	var buf bytes.Buffer
	for _, r := range runes {
		decimal, err := runeToHexInt64(r)
		if err != nil {
			return "", fmt.Errorf("[%d] at [%s] cannot admit to decimal value", r, str)
		}

		// if alphabet and ignoreCases set true, change to lower case
		if decimal >= 65 && decimal <= 90 {
			if ignoreCases == true {
				buf.WriteString(string(r + 32))
			} else {
				buf.WriteString(string(r))
			}
			continue
		}
		/*
				ASCII
			Character 	Decimal
				!		33
				"		34
				#		35
				$		36
				%		37
				&		38
				'		39
				(		40
				)		41
				*		42
				+		43
				,		44
				-		45
				.		46
				/		47

				:		58
				;		59
				<		60
				=		61
				>		62
				?		63
				@		64

				[		91
				\		92
				]		93
				^		94
				_		95
				`		96

				{		123
				|		124
				}		125
				~		126
		*/
		if ignoreSpecialChars == true {
			if decimal >= DecimalExclamationMark && decimal <= DecimalSlash {
				continue
			}
			if decimal >= DecimalColon && decimal <= DecimalAt {
				continue
			}
			if decimal >= DecimalLeftSquareBracket && decimal <= DecimalGraveAccent {
				continue
			}
			if decimal >= DecimalLeftCurlyBracket && decimal <= DecimalTilde {
				continue
			}
		}

		if decimal == DecimalWhitespace {
			if ignoreSeparators == false {
				buf.WriteString(string(r))
			}
			continue
		}

		if dist := getDistanceKorean(decimal); dist != -1 {
			chosungIndex := getChosungIndex(dist)
			jungsungIndex := getJungsungIndex(dist)
			jongsungIndex := getJongsungIndex(dist)

			buf.WriteString(koranChosung[chosungIndex])
			buf.WriteString(koreanJungsung[jungsungIndex])
			if jongsungIndex != 0 {
				buf.WriteString(koreanJongsung[jongsungIndex])
			}
		} else {
			if opt != SplitOptGetOnlyKorean {
				buf.WriteString(string(r))
			}
		}
	}
	return buf.String(), nil
}

// Split disassemble sentence of korean
func Split(str string, opt SplitOpt) ([][]string, error) {
	runes := []rune(str)
	var words [][]string
	for _, r := range runes {
		hex, err := runeToHexInt64(r)
		if err != nil {
			return words, fmt.Errorf("[%d] at [%s] cannot admit to hex value", r, str)
		}
		var tokens []string
		if dist := getDistanceKorean(hex); dist != -1 {
			chosungIndex := getChosungIndex(dist)
			jungsungIndex := getJungsungIndex(dist)
			jongsungIndex := getJongsungIndex(dist)

			tokens = append(tokens, koranChosung[chosungIndex])
			tokens = append(tokens, koreanJungsung[jungsungIndex])
			if jongsungIndex != 0 {
				tokens = append(tokens, koreanJongsung[jongsungIndex])
			}
		} else if hex == DecimalWhitespace { // whitespace
			words = append(words, []string{string(r)})
		} else {
			if opt != SplitOptGetOnlyKorean {
				tokens = append(tokens, string(r))
			}
		}
		if len(tokens) > 0 {
			words = append(words, tokens)
		}
	}
	return words, nil
}
