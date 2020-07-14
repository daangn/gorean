package gorean

import (
	"math/rand"
)

type SortOption int16

const (
	SortOptAsc SortOption = iota
	SortOptDesc
	SortOptAscPivotFirst
	SortOptDescPivotFirst
)

// Sort can sort any strings
func Sort(ary []string, opt SortOption) []string {
	quickSort(ary, opt)
	return ary
}

// quickSort is sort using string array
func quickSort(ary []string, opt SortOption) {
	if len(ary) < 2 {
		return
	}

	lenAry := len(ary)
	left := 0
	right := lenAry - 1

	pivot := rand.Int() % lenAry

	ary[pivot], ary[right] = ary[right], ary[pivot]

	for i, _ := range ary {
		switch opt {
		case SortOptAsc:
			if sum([]rune(ary[i])) < sum([]rune(ary[right])) {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		case SortOptDesc:
			if sum([]rune(ary[i])) > sum([]rune(ary[right])) {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		case SortOptAscPivotFirst:
			if []rune(ary[i])[0] < []rune(ary[right])[0] {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		case SortOptDescPivotFirst:
			if []rune(ary[i])[0] > []rune(ary[right])[0] {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		}
	}

	ary[left], ary[right] = ary[right], ary[left]

	quickSort(ary[:left], opt)
	quickSort(ary[left+1:], opt)
}

func sum(runes []rune) int32 {
	var i int32
	for _, r := range runes {
		i = r + i
	}
	return i
}
