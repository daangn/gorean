package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
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
	dummyQuestionGenerateEdgeNGramTokensKoreanWithWhitespace = "마따 말기"
	dummyAnswerGenerateEdgeNGramTokensKoreanWithWhitespace   = []string{
		"ㅁ",
		"마",
		"마ㄷ",
		"맏",
		"마ㄸ",
		"마따",
		"마따 ",
		"마따 ㅁ",
		"마따 마",
		"마따 말",
		"마따 맑",
		"마따 말ㄱ",
		"마따 말기",
	}
	dummyQuestionGenerateEdgeNGramTokensKoreanWithEnglish = "님부스brush JCM 2000우효"
	dummyAnswerGenerateEdgeNGramTokensKoreanWithEnglish   = []string{
		"ㄴ",
		"니",
		"님",
		"님ㅂ",
		"님부",
		"님붓",
		"님부ㅅ",
		"님부스",
		"님부스b",
		"님부스br",
		"님부스bru",
		"님부스brus",
		"님부스brush",
		"님부스brush ",
		"님부스brush J",
		"님부스brush JC",
		"님부스brush JCM",
		"님부스brush JCM ",
		"님부스brush JCM 2",
		"님부스brush JCM 20",
		"님부스brush JCM 200",
		"님부스brush JCM 2000",
		"님부스brush JCM 2000ㅇ",
		"님부스brush JCM 2000우",
		"님부스brush JCM 2000웋",
		"님부스brush JCM 2000우ㅎ",
		"님부스brush JCM 2000우효",
	}

	dummyQuestionGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase1 = "립밤"
	dummyAnswerGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase1   = []string{
		"ㄹ",
		"리",
		"립",
		"립ㅂ",
		"립바",
		"립밤",
	}

	dummyQuestionGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase2 = "낮잠"
	dummyAnswerGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase2   = []string{
		"ㄴ",
		"나",
		"낮",
		"낮ㅈ",
		"낮자",
		"낮잠",
	}
)

func Test_GenerateForwardNGramTokens(t *testing.T) {
	tokensKorean, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensKorean)
	assert.Empty(t, err)
	assert.Equal(t, dummyAnswerGenerateEdgeNGramTokensKorean, tokensKorean, "")

	tokensKoreanWithWhitespace, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensKoreanWithWhitespace)
	assert.Empty(t, err)
	assert.Equal(t, dummyAnswerGenerateEdgeNGramTokensKoreanWithWhitespace, tokensKoreanWithWhitespace, "equal")

	tokensKoreanWithEnglish, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensKoreanWithEnglish)
	assert.Empty(t, err)
	assert.Equal(t, dummyAnswerGenerateEdgeNGramTokensKoreanWithEnglish, tokensKoreanWithEnglish, "equal")

	tokensSameJaumAtJongsungChosungCase1, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase1)
	assert.Empty(t, err)
	assert.Equal(t, dummyAnswerGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase1, tokensSameJaumAtJongsungChosungCase1, "equal")

	tokensSameJaumAtJongsungChosungCase2, err := GenerateEdgeNGramTokens(dummyQuestionGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase2)
	assert.Empty(t, err)
	assert.Equal(t, dummyAnswerGenerateEdgeNGramTokensSameJaumAtJongsungChosungCase2, tokensSameJaumAtJongsungChosungCase2, "equal")
}
