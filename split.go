package gorean

import "fmt"

type SplitOpt int16

const (
	SplitOptBasic SplitOpt = iota
	SplitOptGetOnlyKorean
)

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
		} else if hex == 32 { // whitespace
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
