package gorean

import (
	"fmt"
	"strconv"
)

// TokenType is type of korean token that ['chosung', 'jungsung', 'jongsung']
type TokenType string

// These are TokenTypes
const (
	typeChosung  TokenType = "chosung"
	typeJungsung TokenType = "jungsung"
	typeJongsung TokenType = "jongsung"
)

func isChosungHex(hex int64) bool {
	if koreanChosungHexToRuneMap[hex] != 0 {
		return true
	}
	return false
}

func isJungsungHex(hex int64) bool {
	if koreanJungsungHexToRuneMap[hex] != 0 {
		return true
	}
	return false
}

func isJongsungHex(hex int64) bool {
	if koreanJongsungHexToRuneMap[hex] != 0 {
		return true
	}
	return false
}

func isKoreanHex(hex int64) bool {
	return hex >= koreanStartHexCode && hex <= koreanEndHexCode
}

func isKoreanToken(r rune, ktt TokenType) bool {
	hex, err := runeToHexInt64(r)
	if err != nil {
		return false
	}
	if ktt == typeChosung && isChosungHex(hex) {
		return true
	} else if ktt == typeJungsung && isJungsungHex(hex) {
		return true
	} else if ktt == typeJongsung && isJongsungHex(hex) {
		return true
	}
	return false
}

func getDistanceKorean(hex int64) int64 {
	if isKoreanHex(hex) {
		return hex - koreanStartHexCode
	}
	return -1
}

func getChosungIndex(dist int64) int64 {
	return ((dist - (dist % koreanJongsungLen)) / koreanJongsungLen) / koreanJungsungLen
}
func getJungsungIndex(dist int64) int64 {
	return ((dist - (dist % koreanJongsungLen)) / koreanJongsungLen) % koreanJungsungLen
}
func getJongsungIndex(dist int64) int64 {
	return dist % koreanJongsungLen
}

func runeToHexInt64(r rune) (int64, error) {
	hex, err := strconv.ParseInt(fmt.Sprintf("0x%x", r), 0, 64)
	if err != nil {
		return 0, err
	}
	return hex, nil
}

func intArrayToString(intAry []int) string {
	if len(intAry) == 0 {
		return ""
	}

	estimate := len(intAry) * 4
	b := make([]byte, 0, estimate)

	for _, n := range intAry {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}
	b = b[:len(b)-1]
	return string(b)
}
