package gorean

import "regexp"

// Korean is find korean text at statement.
// Warning: terms of str is over 5000 terms. You'd better use flash text algorithm.
func Korean(str string, findMax int) []string {
	compile, _ := regexp.Compile("[ㄱ-힣]+")
	return compile.FindAllString(str, findMax)
}
