package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dummyJoinTokensQuestionEmpty = []string{}
	dummyJoinTokensAnswerEmpty   = ""

	dummyJoinTokensQuestionWhitespace = []string{" "}
	dummyJoinTokensAnswerWhitespace   = ""

	dummyJoinTokensQuestionTripleWhitespace = []string{" ", " ", " "}
	dummyJoinTokensAnswerTripleWhitespace   = ""

	dummyJoinTokensQuestionSixthWhitespace = []string{" ", " ", " ", " ", " ", " "}
	dummyJoinTokensAnswerSixthWhitespace   = ""

	dummyJoinTokensQuestionEnglishAlphabets = []string{"e", "c", "o"}
	dummyJoinTokensAnswerEnglishAlphabets   = ""

	dummyJoinTokensQuestionOneKoreanAlphabets = []string{"ㄱ"}
	dummyJoinTokensAnswerOneKoreanAlphabets   = ""

	dummyJoinTokensQuestionTwoKoreanAlphabets = []string{"ㄱ", "ㅏ"}
	dummyJoinTokensAnswerTwoKoreanAlphabets   = "가"

	dummyJoinTokensQuestionThreeKoreanAlphabets = []string{"ㄱ", "ㅏ", "ㅇ"}
	dummyJoinTokensAnswerThreeKoreanAlphabets   = "강"

	dummyJoinTokensQuestionFourKoreanAlphabets = []string{"ㅁ", "ㅏ", "ㅇ", "ㅏ"}
	dummyJoinTokensAnswerFourKoreanAlphabets   = ""
)

func Test_JoinTokens(t *testing.T) {
	strEmpty, strEmptyErr := JoinTokens(dummyJoinTokensQuestionEmpty)
	assert.NotEmpty(t, strEmptyErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerEmpty, strEmpty, "empty")

	strWhitespace, strWhitespaceErr := JoinTokens(dummyJoinTokensQuestionWhitespace)
	assert.NotEmpty(t, strWhitespaceErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerWhitespace, strWhitespace, "empty")

	strTripleWhitespace, strTripleWhitespaceErr := JoinTokens(dummyJoinTokensQuestionTripleWhitespace)
	assert.NotEmpty(t, strTripleWhitespaceErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerTripleWhitespace, strTripleWhitespace, "empty")

	strSixthWhitespace, strSixthWhitespaceErr := JoinTokens(dummyJoinTokensQuestionSixthWhitespace)
	assert.NotEmpty(t, strSixthWhitespaceErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerSixthWhitespace, strSixthWhitespace, "empty")

	strEnglishAlphabets, strEnglishAlphabetsErr := JoinTokens(dummyJoinTokensQuestionEnglishAlphabets)
	assert.NotEmpty(t, strEnglishAlphabetsErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerEnglishAlphabets, strEnglishAlphabets, "only accept korean alphabets")

	strOneKoreanAlphabets, strOneKoreanAlphabetsErr := JoinTokens(dummyJoinTokensQuestionOneKoreanAlphabets)
	assert.NotEmpty(t, strOneKoreanAlphabetsErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerOneKoreanAlphabets, strOneKoreanAlphabets, "empty")

	strTwoKoreanAlphabets, strTwoKoreanAlphabetsErr := JoinTokens(dummyJoinTokensQuestionTwoKoreanAlphabets)
	assert.Empty(t, strTwoKoreanAlphabetsErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerTwoKoreanAlphabets, strTwoKoreanAlphabets, "the word is 당")

	strThreeKoreanAlphabets, strThreeKoreanAlphabetsErr := JoinTokens(dummyJoinTokensQuestionThreeKoreanAlphabets)
	assert.Empty(t, strThreeKoreanAlphabetsErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerThreeKoreanAlphabets, strThreeKoreanAlphabets, "the word is 당")

	strFourKoreanAlphabets, strFourKoreanAlphabetsErr := JoinTokens(dummyJoinTokensQuestionFourKoreanAlphabets)
	assert.NotEmpty(t, strFourKoreanAlphabetsErr, "Have to occurred an error")
	assert.Equal(t, dummyJoinTokensAnswerFourKoreanAlphabets, strFourKoreanAlphabets, "empty")
}

func Test_IsAbleToComposeAlphabetsForSingleCharacter(t *testing.T) {
	startAlphabet := rune('가')
	endAlphabet := rune('힣')

	// test 가-힣
	for c := startAlphabet; c <= endAlphabet; c = c + 1 {
		s := string(c)
		token, err := Split(s, SplitOptBasic)
		isAble := IsAbleToComposeAlphabetsForSingleCharacter(token[0])
		noneKoreanOffset := FindNoneKoreanAlphabetsForSingleCharacter(token[0])
		assert.Empty(t, err)
		assert.Empty(t, noneKoreanOffset)
		assert.Equal(t, isAble, true)
	}
}

func Test_SplitAndJoinComplete(t *testing.T) {
	startAlphabet := rune('가')
	endAlphabet := rune('힣')

	// test 가-힣
	for c := startAlphabet; c <= endAlphabet; c = c + 1 {
		s := string(c)
		token, err := Split(s, SplitOptBasic)
		assert.Empty(t, err)
		tokenStr, err := JoinTokens(token[0])
		assert.Empty(t, err)
		assert.Equal(t, s, tokenStr)
	}
}
