# gorean
한글 분석 유틸리티 golang 라이브러리.

- 기본적으로 [ruby 'korean-string'](https://github.com/bhumphreys/korean-string)라이브러리를 golang으로 포팅 했습니다.
- main reference
    - https://www.bada-ie.com/board/view/?page=9&uid=1782&category_code=&code=all
    - http://www.w3c.or.kr/i18n/hangul-i18n/ko-code.html

# Method Summary

``` go
// See `korean_test.go` and `util_test.go` files

- gorean.SplitKorean // 문자열 한 글자를 토크나이징

- gorean.JoinTokens // 토큰 최대 3개를 활용해 한글 조합문자 만들기

- gorean.IsAbleToComposeAlphabetsForSingleCharacter// 문자열 한 글자에 대한 토큰들로 한글 조합문자가 가능한지 확인

- gorean.FindNoneKoreanAlphabetsForSingleCharacter // 문자열 한 글자에 대한 토큰들 중에 한글 외 문자열이 있는지 체크

- gorean.GenerateEdgeNGramTokens: 한글 문자열 edge ngram 토크나이징 // 외국어, 숫자, 특문 토크나이징 안하고 토큰화

- gorean.SortStringArray: 한글 문자열 정렬 가능
```

#### gorean.SplitKorean
- 입력값으로 장문의 문자열을 받을 수 있으며, 각각 한글자에 따른 2개~3개의 elements로 되어있는 자모 배열이 나오게 되며, 결과값은 이중배열이 나오게 된다.

``` go
question := " karrot  마켓 "
answer   := [][]string{
    []string{" "},
    []string{"k"},
    []string{"a"},
    []string{"r"},
    []string{"r"},
    []string{"o"},
    []string{"t"},
    []string{" "},
    []string{" "},
    []string{"ㅁ", "ㅏ"},
    []string{"ㅋ", "ㅔ", "ㅅ"},
    []string{" "},
}
strKoreanAndEnglishWithWhitespace, err := SplitKorean(question, false)
assert.Equal(t, question, answer, "korean and english with whitespace") // success
```

#### gorean.JoinTokens
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열을 조합해서 하나의 한글 조합문자열을 만든다.


#### gorean.IsAbleToComposeAlphabetsForSingleCharacter
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열이 한국어 조합글자가 될 수 있는지 에 대한 검증코드
- 디버깅 목적으로 만듬, JoinTokens를 하기 이전에 체크해보고 넘어 가라고 만듬


#### gorean.FindNoneKoreanAlphabetsForSingleCharacter
- 입력값으로 주어질 배열 2개~3개의 string들이 되어있는 배열에 한글 자모가 아닌 글자가 포함되어있는지 확인하는 디버깅 코드
- 디버깅 목적으로 만듬, IsAbleToComposeAlphabetsForSingleCharacter에서 false일 때에 사용하도록 의도

``` go
tokens := []string{"ㅎ", "ㅣ", "g"}
var singleKorean string
if isKoreanToken := gorean.IsAbleToComposeAlphabetsForSingleCharacter(tokens); isKoreanToken == false {
offset := gorean.FindNoneKoreanAlphabetsForSingleCharacter(tokens)
return "", fmt.Errorf(
    "[%s] offset [%s] has problem. not korean or not right token type about position",
    strings.Join(tokens, ","),
    intArrayToString(offset),
)
// have to occur an error
} else {
singleKorean = gorean.JoinTokens(tokens)
fmt.Println(singleKorean)
}
```

#### gorean.GenerateEdgeNGramTokens
- 한글 EdgeNGram 및 전방일치 토크나이저를 얻기위해 만듬
- 강남역 => [ㄱ,가,강,강ㄴ,강나,강남,강남ㅇ,강남여,강남역]

#### gorean.SortStringArray
문자열 한 글자 토크나이징

``` go
dummyQuestionGenerateEdgeNGramTokensKorean = "까투리왕"
dummyAnswerGenerateEdgeNGramTokensKorean   = []string{
    "ㄱ",
    "ㄲ",
    "까",
    "깥",
    "까ㅌ",
    "까투",
    "까툴",
    "까투ㄹ",
    "까투리",
    "까투링",
    "까투리ㅇ",
    "까투리오",
    "까투리와",
    "까투리왕",
}
tokensKoreanWithEnglish, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensKoreanWithEnglish)
sortedKoreanWithEnglish := SortStringArray(tokensKoreanWithEnglish, false)
assert.Equal(t, sortedKoreanWithEnglish, dummyAnswerGenerateEdgeNGramTokensKoreanWithEnglish, "equal")
```

# PR
- 메서드 네이밍이 구리니 제발 좋은 이름 추천부탁드려요. 흑 ; ㅅ ;
I'll appreciate to any PR, Issues or Suggestions

# Release note