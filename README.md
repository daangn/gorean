# gorean
korean analyzer utility tools

- basically [ruby 'korean-string'](https://github.com/bhumphreys/korean-string) is transforming using golang
- main reference
    - https://www.bada-ie.com/board/view/?page=9&uid=1782&category_code=&code=all
    - http://www.w3c.or.kr/i18n/hangul-i18n/ko-code.html

# Getting started

``` go
// See `korean_test.go` and `util_test.go` files

- gorean.SplitKorean // 문자열 한 글자를 토크나이징

- gorean.JoinTokens // 토큰 최대 3개를 활용해 한글 조합문자 만들기

- gorean.IsAbleToComposeAlphabetsForSingleCharacter// 문자열 한 글자에 대한 토큰들로 한글 조합문자가 가능한지 확인

- gorean.FindNoneKoreanAlphabetsForSingleCharacter // 문자열 한 글자에 대한 토큰들 중에 한글 외 문자열이 있는지 체크

- gorean.GenerateEdgeNGramTokens: 한글 문자열 edge ngram 토크나이징 // 외국어, 숫자, 특문 토크나이징 안하고 토큰화

- gorean.SortStringArray: 한글 문자열 정렬 가능
```

# PR
- 메서드 네이밍이 구리니 제발 좋은 이름 추천부탁드려요. 흑 ; ㅅ ;
I'll appreciate to any PR, Issues or Suggestions

# Release note