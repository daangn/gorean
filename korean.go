package gorean

import (
	"fmt"
	"strings"
)

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

// IsAbleToComposeAlphabetsForSingleCharacter check alphabets to compose for single character
func IsAbleToComposeAlphabetsForSingleCharacter(tokens []string) bool {
	offset := FindNoneKoreanAlphabetsForSingleCharacter(tokens)

	offsetLen := len(offset)
	if offsetLen == 0 {
		return true
	}
	return false
}

// FindNoneKoreanAlphabetsForSingleCharacter find offset about none korean alphabets for single character
func FindNoneKoreanAlphabetsForSingleCharacter(tokens []string) []int {
	runes := []rune(strings.Join(tokens, ""))

	var offset []int
	for i, r := range runes {
		if i == chosungIndex && r != 0 && isKoreanToken(r, typeChosung) == false {
			offset = append(offset, chosungIndex)
		}
		if i == jungsungIndex && r != 0 && isKoreanToken(r, typeJungsung) == false {
			offset = append(offset, jungsungIndex)
		}
		if i == jongsungIndex && r != 0 && isKoreanToken(r, typeJongsung) == false {
			offset = append(offset, jongsungIndex)
		}
	}
	return offset
}

// JoinTokens compose to word using only korean alphabets
func JoinTokens(tokens []string) (string, error) {
	// FIXME: Reviewing Point how to trim `tokens` and make `runes` more efficiently whole this method has illness with that
	tokensLen := len(tokens)
	if tokensLen == 2 {
		tokens = append(tokens, "")
	} else if tokensLen < 2 || tokensLen > 3 {
		return "", fmt.Errorf(
			"(%s) has been outranged tokens for JoinKorean",
			strings.Join(tokens, ","),
		)
	} else if isKoreanToken := IsAbleToComposeAlphabetsForSingleCharacter(tokens); isKoreanToken == false {
		offset := FindNoneKoreanAlphabetsForSingleCharacter(tokens)
		return "", fmt.Errorf(
			"[%s] offset [%s] has problem. not korean or not right token type about position",
			strings.Join(tokens, ","),
			intArrayToString(offset),
		)
	}

	tokensStr := strings.Trim(strings.Join(tokens, ""), " ")

	word := []int64{0xAC00}
	runes := []rune(tokensStr)

	for i, r := range runes {
		hex, err := runeToHexInt64(r)
		if err != nil {
			return "", fmt.Errorf("[%d] at [%s] cannot admit to hex value", r, tokensStr)
		}
		if i == chosungIndex && r != 0 {
			index := koreanChosungRuneToIndexMap[hex]
			word = append(word, 28*21*index)
		} else if i == jungsungIndex && r != 0 {
			index := koreanJungsungRuneToIndexMap[hex]
			word = append(word, 28*index)
		} else if i == jongsungIndex && r != 0 {
			index := koreanJongsungRuneToIndexMap[hex]
			word = append(word, index)
		}
	}
	sum := int64(0)
	for _, r := range word {
		sum += r
	}
	return string(sum), nil
}
