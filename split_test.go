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
