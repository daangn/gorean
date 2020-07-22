<p align="center">
    <a href="https://github.com/daangn/gorean"><img src="https://camo.githubusercontent.com/3e224d221eae1953ae073b0050197faf4b4428af/68747470733a2f2f636972636c6563692e636f6d2f67682f6461616e676e2f676f7265616e2e7376673f7374796c653d737667" alt="daangn" data-canonical-src="https://circleci.com/gh/daangn/gorean.svg?style=svg" style="max-width:100%;"></a>
    <a href="https://GitHub.com/daangn/gorean/releases/"><img src="https://camo.githubusercontent.com/de0de733bb357134d4a347556dd34a809a7a8ddf/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f72656c656173652f6461616e676e2f676f7265616e2e737667" alt="GitHub release" data-canonical-src="https://img.shields.io/github/release/daangn/gorean.svg" style="max-width:100%;"></a>
    <a href="https://GitHub.com/daangn/gorean/releases/"><img src="https://camo.githubusercontent.com/f3808af4419bfd709a7e5abb0bf3a7f0f5009b29/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f646f776e6c6f6164732f6461616e676e2f676f7265616e2f746f74616c2e737667" alt="Github all releases" data-canonical-src="https://img.shields.io/github/downloads/daangn/gorean/total.svg" style="max-width:100%;"></a>
    <a href="https:/
/GitHub.com/daangn/gorean/graphs/contributors/"><img src="https://camo.githubusercontent.com/68dfce75e28f1c6bc86d7c97f0b9ec4bbce59036/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f636f6e7472696275746f72732f6461616e676e2f676f7265616e2e737667" alt="GitHub contributors" data-canonical-src="https://img.shields.io/github/contributors/daangn/gorean.svg" style="max-width:100%;"></a>
    <a href="https://lbesson.mit-license.org/" rel="nofollow"><img src="https://camo.githubusercontent.com/311762166ef25238116d3cadd22fcb6091edab98/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d4d49542d626c75652e737667" alt="MIT license" data-canonical-src="https://img.shields.io/badge/License-MIT-blue.svg" style="max-width:100%;"></a>
</p>

<h1 align="center"><a id="user-content--gorean" class="anchor" aria-hidden="true" href="#-gorean"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a><g-emoji class="g-emoji" alias="kr" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f1f0-1f1f7.png">🇰🇷</g-emoji> Gorean</h1>

<p align="center">golang native로 작성된 한글 분석 유틸리티 라이브러리 입니다.
기본적으로 <a href="https://github.com/bhumphreys/korean-string">ruby 'korean-string'</a>라이브러리를 golang으로 포팅한 라이브러리 이며,
그 이외 한글 분석 유틸리티 도구들을 준비했습니다.
해당 도구는 한글검색에 사용되는 한글분석 유틸리티를 모아둘 예정입니다.</p>

<p align="center">라이브러리의 코드는 <a href="https://www.bada-ie.com/board/view/?page=9&amp;uid=1782&amp;category_code=&amp;code=all" rel="nofollow">bada-ie</a> 님의 코드를 참고하여 구현하였으며,
자세한 한글 인코딩 관련 정보는 <a href="http://www.w3c.or.kr/i18n/hangul-i18n/ko-code.html" rel="nofollow">w3c-hangul-i18n</a>에서 열람하실 수 있습니다.</p>


# 🍗 Speed Cheat Sheet

``` go
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
	gorean.Sort(messKoreanSort, gorean.SortOptAsc)
	fmt.Println(messKoreanSort) // [김치볶음밥 까치꾸치 동백 마장동 밥상머리 사자왕왕 왕초 자루소바오이시 하기스]

	koreanWithEnglish := "초성퀴즈 with English"
	korean := gorean.Korean(koreanWithEnglish, 10)
	fmt.Println(korean) // 초성퀴즈

	chosung, _ := gorean.Chosung(strings.Join(korean, " "))
	fmt.Println(strings.Join(chosung, "")) // ㅊㅅㅋㅈ
}

```

# 🍱 API Summary

``` go
- func gorean.Split([]string, SplitOpt)
    - type SplitOpt gorean.SplitOptBasic
    - type SplitOpt gorean.SplitOptGetOnlyKorean
- func gorean.JoinTokens([]string) // 2 <= len(ary) <= 3
- func gorean.IsAbleToComposeAlphabetsForSingleCharacter([]string) // 2 <= len(ary) <= 3
- func gorean.FindNoneKoreanAlphabetsForSingleCharacter([]string) // 2 <= len(ary) <= 3
- func gorean.GenerateEdgeNGramTokens(string)
- func gorean.Sort([]string, SortOpt)
    - type SortOpt gorean.SortOptAsc
    - type SortOpt gorean.SortOptDesc
```

#### gorean.Split
- 입력값으로 장문의 문자열을 받을 수 있으며, 각각 한글자에 따른 2개~3개의 elements로 되어있는 자모 배열이 나오게 되며, 결과값은 이중배열이 나오게 된다.

#### gorean.JoinTokens
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열을 조합해서 하나의 한글 조합문자열을 만든다.

#### gorean.IsAbleToComposeAlphabetsForSingleCharacter
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열이 한국어 조합글자가 될 수 있는지 에 대한 검증코드
- 디버깅 목적으로 만듬, JoinTokens를 하기 이전에 체크해보고 넘어 가라고 만듬

#### gorean.FindNoneKoreanAlphabetsForSingleCharacter
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열에 한글 자모가 아닌 글자가 포함되어있는지 확인하는 디버깅 코드
- 디버깅 목적으로 만듬, IsAbleToComposeAlphabetsForSingleCharacter에서 false일 때에 사용하도록 의도

#### gorean.GenerateEdgeNGramTokens
- 한글 EdgeNGram 및 전방일치 토크나이저를 얻기위해 만듬
- 강남역 => [ㄱ,가,강,강ㄴ,강나,강남,강남ㅇ,강남여,강남역]

#### gorean.Sort
- 문자열 정렬을 위해 존재함. 정렬의 다양한 옵션 제공

# 📝 Release note

- `v0.0.5` **[Latest Release]**
    1. 기능추가 초성 얻기, `Chosung()`
    2. 기능추가 한글 찾기, `Korean()`
- `v0.0.6` [ToDo]
    1. 기능추가 영어자판 to 한글 
    2. 기능추가 한글자판 to 영어자판
- `v0.0.7`
    1. 기능업그레이드: 한글 찾기, (flash-text)[https://medium.com/@jwyeom63/%EB%B2%88%EC%97%AD-%EC%A0%95%EA%B7%9C%ED%91%9C%ED%98%84%EC%8B%9D%EC%9C%BC%EB%A1%9C-5%EC%9D%BC-%EA%B1%B8%EB%A6%AC%EB%8A%94-%EC%9E%91%EC%97%85-15%EB%B6%84%EB%A7%8C%EC%97%90-%EB%81%9D%EB%82%B4%EA%B8%B0-2e615a907048] 알고리즘을 이용해 한글찾기 개선
- `v0.1.1` [ToDo]
    1. benchmark 검증 및 성능 최적화

# 👍 Contribute
- @drake-jin 은 부족한게 많습니다. 피드백은 언제나 환영이에요.
- 뭔가 추가 해줬으면 하는게 있거나, 변경했으면 하는게 있으면 언제든지 이슈 및 PR 보내주세요. 특별한 일 없으면 **2일안에** 반영해드릴게요.

만약 이 라이브러리 daangn/gorean을 잘 사용하셨다면...

1. [GitHub Star](https://github.com/daangn/gorean/stargazers)
2. 이직각을 재고 계시다면 🥕[당근마켓](https://www.notion.so/daangn/2c789a2c7b1a4cfca40b11afba678315)🥕