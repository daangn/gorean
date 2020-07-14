package gorean

import (
	"fmt"
)

// GenerateEdgeNGramTokens is for creating forward match at service
func GenerateEdgeNGramTokens(str string) ([]string, error) {
	var tokens []string
	var stackedKeywords []string
	argRunes := []rune(str)
	for alphabetPos := range argRunes {
		var stack []rune
		for i := 0; i <= alphabetPos; i++ {
			stack = append(stack, argRunes[i])
		}
		stackedKeywords = append(stackedKeywords, string(stack))
	}
	splits, err := Split(str, SplitOptBasic)
	if err != nil {
		return tokens, err
	}
	var previousSplit []string
	for splitIndex, split := range splits {
		stackedKeyword := ""
		if splitIndex >= 1 {
			stackedKeyword = stackedKeywords[splitIndex-1]
			previousSplit = splits[splitIndex-1]
		}

		lenSplit := len(split)
		for splitPos := 0; splitPos < lenSplit; splitPos++ {
			// 초성 일 때, whitespace 가 아니며, 한글 일 경우
			if splitPos == 0 && split[0] != " " && isKoreanHex(int64(argRunes[splitIndex])) {
				chosung := split[0]
				if coupleJaumFirst := koreanCoupleJaumFirstMap[chosung]; coupleJaumFirst != "" {
					chosung = coupleJaumFirst
					tokens = append(tokens, stackedKeyword+chosung)
				}

				lenPreviousSplit := len(previousSplit)
				if lenPreviousSplit == 2 {
					alphabets := append(previousSplit, chosung)
					if joinAlphabet, joinAlphabetErr := JoinTokens(alphabets); joinAlphabetErr != nil {
						// Allow to skip
						fmt.Printf("이전글자 초성 중성과 현재 글자의 초성 조합 중 토큰조합 실패: %s, %s\n%v\n", str, alphabets, joinAlphabetErr)
					} else {
						tokens = append(tokens, string(argRunes[0:splitIndex-1])+joinAlphabet)
					}
				} else if lenPreviousSplit == 3 {
					if coupleJaum := koreanCoupleJaumMap[previousSplit[2]+chosung]; coupleJaum != "" {
						alphabets := []string{previousSplit[0], previousSplit[1], coupleJaum}
						if koreanDoubleJaumMap[alphabets[2]] != true {
							if joinAlphabet, joinAlphabetErr := JoinTokens(alphabets); joinAlphabetErr != nil {
								// Allow to skip
								fmt.Printf("이전글자의 종성과 현재글자의 초성 조합 중 토큰조합 실패: %s, %s\n%v\n", str, alphabets, joinAlphabetErr)
							} else {
								tokens = append(tokens, string(argRunes[0:splitIndex-1])+joinAlphabet)
							}
						}
					}
				}
			}

			// Jungsung - add single Moum when it couple Moum token
			if splitPos == 1 {
				jungsung := split[1]
				if coupleMoumFirst := koreanCoupleMoumFirstMap[jungsung]; coupleMoumFirst != "" {
					alphabets := []string{split[0], coupleMoumFirst}
					if joinAlphabet, joinAlphabetErr := JoinTokens(alphabets); joinAlphabetErr != nil {
						// Allow skip
						fmt.Printf("조합 중성에 대한 분리 중 토큰조합 실패: %s, %s\n%v\n", str, alphabets, joinAlphabetErr)
					} else {
						tokens = append(tokens, stackedKeyword+joinAlphabet)
					}
				}
			}

			// Jongsung - add single Jaum when it couple Jaum token
			if splitPos == 2 {
				jongsung := split[2]
				if coupleJaumFirst := koreanCoupleJaumFirstMap[jongsung]; coupleJaumFirst != "" {
					alphabets := []string{split[0], split[1], coupleJaumFirst}
					if joinAlphabet, joinAlphabetErr := JoinTokens(alphabets); joinAlphabetErr != nil {
						// Allow to skip
						fmt.Printf("조합 종성에 대한 분리로직 중 토큰조합 실패: %s, %s\n%v\n", str, alphabets, joinAlphabetErr)
					} else {
						tokens = append(tokens, stackedKeyword+joinAlphabet)
					}
				}
			}

			// base - generate token
			var stack []string
			token := stackedKeyword
			for subSplitPos := 0; subSplitPos <= splitPos; subSplitPos++ {
				stack = append(stack, split[subSplitPos])
			}
			if splitPos == 0 {
				token += stack[0]
			} else {
				if joinToken, joinTokenErr := JoinTokens(stack); joinTokenErr != nil {
					// Allow to skip
					fmt.Printf("조건 없는 split을 한 token의 조합 중 토큰조합 실패: %s, %s\n%v\n", str, stack, joinTokenErr)
				} else {
					token += joinToken
				}
			}
			tokens = append(tokens, token)
		}
	}
	return tokens, nil
}
