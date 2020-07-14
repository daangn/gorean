package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dummySplitQuestionEmpty = ""
	dummySplitAnswerEmpty   [][]string

	dummySplitQuestionWhitespace = " "
	dummySplitAnswerWhitespace   = [][]string{[]string{" "}}

	dummySplitQuestionTripleWhitespace = "   "
	dummySplitAnswerTripleWhitespace   = [][]string{[]string{" "}, []string{" "}, []string{" "}}

	dummySplitQuestionSingleKoreanWord = "가"
	dummySplitAnswerSingleKoreanWord   = [][]string{
		[]string{"ㄱ", "ㅏ"},
	}

	dummySplitQuestionKoreanWord = "당근마켓"
	dummySplitAnswerKoreanWord   = [][]string{
		[]string{"ㄷ", "ㅏ", "ㅇ"},
		[]string{"ㄱ", "ㅡ", "ㄴ"},
		[]string{"ㅁ", "ㅏ"},
		[]string{"ㅋ", "ㅔ", "ㅅ"},
	}

	dummySplitQuestionKoreanWordWithWhitespace = " 당신의 근처 "
	dummySplitAnswerKoreanWordWithWhitespace   = [][]string{
		[]string{" "},
		[]string{"ㄷ", "ㅏ", "ㅇ"},
		[]string{"ㅅ", "ㅣ", "ㄴ"},
		[]string{"ㅇ", "ㅢ"},
		[]string{" "},
		[]string{"ㄱ", "ㅡ", "ㄴ"},
		[]string{"ㅊ", "ㅓ"},
		[]string{" "},
	}

	dummySplitQuestionEnglish = "karrot"
	dummySplitAnswerEnglish   = [][]string{
		[]string{"k"},
		[]string{"a"},
		[]string{"r"},
		[]string{"r"},
		[]string{"o"},
		[]string{"t"},
	}

	dummySplitQuestionKoreanAndEnglish = "karrot마켓"
	dummySplitAnswerKoreanAndEnglish   = [][]string{
		[]string{"k"},
		[]string{"a"},
		[]string{"r"},
		[]string{"r"},
		[]string{"o"},
		[]string{"t"},
		[]string{"ㅁ", "ㅏ"},
		[]string{"ㅋ", "ㅔ", "ㅅ"},
	}

	dummySplitQuestionKoreanAndEnglishWithWhitespace = " karrot  마켓 "
	dummySplitAnswerKoreanAndEnglishWithWhitespace   = [][]string{
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
)

func Test_Split(t *testing.T) {
	strEmpty, err := Split(dummySplitQuestionEmpty, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerEmpty, strEmpty, "emptyString")
	assert.Empty(t, err)

	strWhitespace, err := Split(dummySplitQuestionWhitespace, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerWhitespace, strWhitespace, "single white space")
	assert.Empty(t, err)

	strTripleWhitespace, err := Split(dummySplitQuestionTripleWhitespace, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerTripleWhitespace, strTripleWhitespace, "multiple white space")
	assert.Empty(t, err)

	strSingleKoreanWord, err := Split(dummySplitQuestionSingleKoreanWord, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerSingleKoreanWord, strSingleKoreanWord, "multiple white space")
	assert.Empty(t, err)

	strKoreanWord, err := Split(dummySplitQuestionKoreanWord, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerKoreanWord, strKoreanWord, "korean word")
	assert.Empty(t, err)

	strKoreanWordWithWhitespace, err := Split(dummySplitQuestionKoreanWordWithWhitespace, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerKoreanWordWithWhitespace, strKoreanWordWithWhitespace, "korean word with whitespace")
	assert.Empty(t, err)

	strEnglish, err := Split(dummySplitQuestionEnglish, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerEnglish, strEnglish, "english word")
	assert.Empty(t, err)

	strKoreanAndEnglish, err := Split(dummySplitQuestionKoreanAndEnglish, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerKoreanAndEnglish, strKoreanAndEnglish, "korean and english")
	assert.Empty(t, err)

	strKoreanAndEnglishWithWhitespace, err := Split(dummySplitQuestionKoreanAndEnglishWithWhitespace, SplitOptBasic)
	assert.Equal(t, dummySplitAnswerKoreanAndEnglishWithWhitespace, strKoreanAndEnglishWithWhitespace, "korean and english with whitespace")
	assert.Empty(t, err)
}

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
