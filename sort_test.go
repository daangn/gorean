package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dummyQuestionSortKoreanEnglishNumber = []string{"ㄱ", "a", "1", "ㄴ", "b", "2", "ㄷ", "c", "3"}
	dummyAnswerSortKoreanEnglishNumber   = []string{"1", "2", "3", "a", "b", "c", "ㄱ", "ㄴ", "ㄷ"}

	//dummyQuestionSortKoreanDifferentRange = []string{
	//	"하기스",
	//	"김치볶음밥",
	//	"사자왕왕",
	//	"자루소바오이시",
	//	"왕초",
	//	"밥상머리",
	//	"까치꾸치",
	//	"마장동",
	//	"동백",
	//}
	//dummyAnswerSortKoreanDifferentRange1 = []string{
	//	"동백",
	//	"왕초",
	//	"마장동",
	//	"하기스",
	//	"까치꾸치",
	//	"밥상머리",
	//	"사자왕왕",
	//	"김치볶음밥",
	//	"자루소바오이시",
	//}
	//dummyAnswerSortKoreanDifferentRange2 = []string{
	//	"자루소바오이시",
	//	"김치볶음밥",
	//	"사자왕왕",
	//	"밥상머리",
	//	"까치꾸치",
	//	"하기스",
	//	"마장동",
	//	"왕초",
	//	"동백",
	//}
	//dummyAnswerSortKoreanDifferentRange3 = []string{
	//	"김치볶음밥",
	//	"까치꾸치",
	//	"동백",
	//	"마장동",
	//	"밥상머리",
	//	"사자왕왕",
	//	"왕초",
	//	"자루소바오이시",
	//	"하기스",
	//}
	//dummyAnswerSortKoreanDifferentRange4 = []string{
	//	"하기스",
	//	"자루소바오이시",
	//	"왕초",
	//	"사자왕왕",
	//	"밥상머리",
	//	"마장동",
	//	"동백",
	//	"까치꾸치",
	//	"김치볶음밥",
	//}
)

func Test_Sort(t *testing.T) {
	assert.Equal(t, dummyAnswerSortKoreanEnglishNumber, Sort(dummyQuestionSortKoreanEnglishNumber, SortOptAsc), "equal")

	//assert.Equal(t, dummyAnswerSortKoreanDifferentRange1, Sort(dummyQuestionSortKoreanDifferentRange, SortOptAsc), "equal")
	//assert.Equal(t, dummyAnswerSortKoreanDifferentRange2, Sort(dummyQuestionSortKoreanDifferentRange, SortOptDesc), "equal")
	//assert.Equal(t, dummyAnswerSortKoreanDifferentRange3, Sort(dummyQuestionSortKoreanDifferentRange, SortOptAscPivotFirst), "equal")
	//assert.Equal(t, dummyAnswerSortKoreanDifferentRange4, Sort(dummyQuestionSortKoreanDifferentRange, SortOptDescPivotFirst), "equal")
}
