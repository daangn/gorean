package gorean

func Chosung(str string) ([]string, error) {
	var result []string

	characters, err := Split(str, SplitOptBasic)
	if err != nil {
		return result, err
	}

	for _, c := range characters {
		r := []rune(c[chosungIndex])
		if isKoreanToken(r[chosungIndex], typeChosung) {
			result = append(result, c[chosungIndex])
		} else {
			result = append(result, c[chosungIndex])
		}
	}
	return result, nil
}
