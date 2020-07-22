package gorean

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	questionChosungKorean = "초성퀴즈"
	answerChosungKorean   = []string{"ㅊ", "ㅅ", "ㅋ", "ㅈ"}

	questionChosungKoreanEnglish = "초성퀴즈 with English"
	answerChosungKoreanEnglish   = []string{"ㅊ", "ㅅ", "ㅋ", "ㅈ", " ", "w", "i", "t", "h", " ", "E", "n", "g", "l", "i", "s", "h"}

	questionChosungKoreanWithWhitespace = "초 대박 개봉박두 아이언맨"
	answerChosungKoreanWithWhitespace   = []string{"ㅊ", " ", "ㄷ", "ㅂ", " ", "ㄱ", "ㅂ", "ㅂ", "ㄷ", " ", "ㅇ", "ㅇ", "ㅇ", "ㅁ"}
)

func Test_Chosung(t *testing.T) {
	chosung, err := Chosung(questionChosungKorean)
	assert.Empty(t, err)
	assert.Equal(t, answerChosungKorean, chosung)

	fmt.Println("---")
	chosungKoreanEnglis, err := Chosung(questionChosungKoreanEnglish)
	assert.Empty(t, err)
	assert.Equal(t, answerChosungKoreanEnglish, chosungKoreanEnglis)

	fmt.Println("---")
	chosungKoreanWithWhitespace, err := Chosung(questionChosungKoreanWithWhitespace)
	assert.Empty(t, err)
	assert.Equal(t, answerChosungKoreanWithWhitespace, chosungKoreanWithWhitespace)
}
