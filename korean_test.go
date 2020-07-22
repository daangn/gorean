package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	questionAnyKoreanStatements = `
안녕하세요 저는 조용진 입니다.

^[^a-zA-Z0-9]+$|^[a-zA-Z0-9]+[^a-zA-Z0-9]+$

출처: https://solskjaer.tistory.com/155 [work smart !!]
`
	inputAnyKoreanStatements = []string{"안녕하세요", "저는", "조용진", "입니다", "출처"}
)

func Test_Korean(t *testing.T) {
	koreans := Korean(questionAnyKoreanStatements, 100000)
	assert.Equal(t, inputAnyKoreanStatements, koreans)
}
