package main

import (
	"fmt"
	"strings"

	"github.com/daangn/gorean"
)

func main() {
	s := "  똠빵각하 5ri구이-  "
	sk, err := gorean.Split(s, gorean.SplitOptBasic)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sk)
		// second parameter[onlyKorean]: gorean.SplitOptBasic
		// [[ ] [ ] [ㄸ ㅗ ㅁ] [ㅃ ㅏ ㅇ] [ㄱ ㅏ ㄱ] [ㅎ ㅏ] [ ] [5] [r] [i] [ㄱ ㅜ] [ㅇ ㅣ] [-] [ ] [ ]]
		// second parameter[onlyKorean]: gorean.SplitOptGetOnlyKorean
		// [[ㄸ ㅗ ㅁ] [ㅃ ㅏ ㅇ] [ㄱ ㅏ ㄱ] [ㅎ ㅏ] [ㄱ ㅜ] [ㅇ ㅣ]]
	}

	var jt1 []string
	var jt2 []string
	for _, tokens := range sk {
		// case: 1
		if character, err := gorean.JoinTokens(tokens); err != nil {
			fmt.Println(err)
			/*
				( ) has been out-ranged tokens for JoinKorean
				( ) has been out-ranged tokens for JoinKorean
				( ) has been out-ranged tokens for JoinKorean
				(5) has been out-ranged tokens for JoinKorean
				(r) has been out-ranged tokens for JoinKorean
				(i) has been out-ranged tokens for JoinKorean
				(-) has been out-ranged tokens for JoinKorean
				( ) has been out-ranged tokens for JoinKorean
				( ) has been out-ranged tokens for JoinKorean
			*/
		} else {
			jt1 = append(jt1, character)
		}

		// case: 2
		if gorean.IsAbleToComposeAlphabetsForSingleCharacter(tokens) {
			character, _ := gorean.JoinTokens(tokens)
			jt2 = append(jt2, character)
		} else {
			noneKoreanToken := gorean.FindNoneKoreanAlphabetsForSingleCharacter(tokens)
			// you should write to something for exception existing none korean tokens
			fmt.Printf("Error! positions [%v] at [%v]\n", tokens, noneKoreanToken)
			/*
				Error! positions [[ ]] at [[0]]
				Error! positions [[ ]] at [[0]]
				Error! positions [[ ]] at [[0]]
				Error! positions [[5]] at [[0]]
				Error! positions [[r]] at [[0]]
				Error! positions [[i]] at [[0]]
				Error! positions [[-]] at [[0]]
				Error! positions [[ ]] at [[0]]
				Error! positions [[ ]] at [[0]]
			*/
		}
	}
	fmt.Printf("jt1 output => %s\n", strings.Join(jt1, "")) // 똠빵각하구이
	fmt.Printf("jt2 output => %s\n", strings.Join(jt2, "")) // 똠빵각하구이

	if edgeNGram, err := gorean.GenerateEdgeNGramTokens(s); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(edgeNGram)
		/*
			// Warning: Including whitespace each items at []string, It didn't Trim
			[       ㄷ   ㄸ   또   똠   똠ㅂ   똠ㅃ   똠빠   똠빵   똠빵ㄱ   똠빵가   똠빵각   똠빵각ㅎ   똠빵각하   똠빵각하    똠빵각하 5   똠빵각하 5r   똠빵각하 5ri   똠빵각하 5riㄱ   똠빵각하 5ri구   똠빵각하 5ri궁   똠빵각하 5ri구ㅇ   똠빵각하 5ri구이   똠빵각하 5ri구이-   똠빵각하 5ri구이-    똠빵각하 5ri구이-  ]
		*/
	}

	messKoreanSort := []string{
		"하기스",
		"김치볶음밥",
		"사자왕왕",
		"자루소바오이시",
		"왕초",
		"밥상머리",
		"까치꾸치",
		"마장동",
		"동백",
	}
	gorean.Sort(messKoreanSort, gorean.SortOptAscPivotFirst)
	fmt.Println(messKoreanSort) // [김치볶음밥 까치꾸치 동백 마장동 밥상머리 사자왕왕 왕초 자루소바오이시 하기스]
}
