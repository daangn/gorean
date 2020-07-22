package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dummyQuestionSortKoreanEnglishNumber = []string{"ㄱ", "a", "1", "ㄴ", "b", "2", "ㄷ", "c", "3"}
	dummyAnswerSortKoreanEnglishNumber   = []string{"1", "2", "3", "a", "b", "c", "ㄱ", "ㄴ", "ㄷ"}

	dummyQuestionSortKorean = []string{
		"고동",
		"동백",
		"왕초",
		"고등",
		"마장동",
		"하기스",
		"고고학자",
		"고무",
		"까치꾸치",
		"밥상머리",
		"사자왕왕",
		"고황",
		"고창",
		"창고",
		"김치볶음밥",
		"자루소바오이시",
	}
	dummyAnswerSortKoreanDesc = []string{
		"하기스",
		"창고",
		"자루소바오이시",
		"왕초",
		"사자왕왕",
		"밥상머리",
		"마장동",
		"동백",
		"까치꾸치",
		"김치볶음밥",
		"고황",
		"고창",
		"고무",
		"고등",
		"고동",
		"고고학자",
	}
	dummyAnswerSortKoreanAsc = []string{
		"고고학자",
		"고동",
		"고등",
		"고무",
		"고창",
		"고황",
		"김치볶음밥",
		"까치꾸치",
		"동백",
		"마장동",
		"밥상머리",
		"사자왕왕",
		"왕초",
		"자루소바오이시",
		"창고",
		"하기스",
	}
)

func Test_Sort(t *testing.T) {
	assert.Equal(t, dummyAnswerSortKoreanEnglishNumber, Sort(dummyQuestionSortKoreanEnglishNumber, SortOptAsc), "equal")
	assert.Equal(t, dummyAnswerSortKoreanAsc, Sort(dummyQuestionSortKorean, SortOptAsc))
	assert.Equal(t, dummyAnswerSortKoreanDesc, Sort(dummyQuestionSortKorean, SortOptDesc))

}
