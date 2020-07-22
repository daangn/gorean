package gorean

import (
	"math/rand"
)

type SortOption int16

const (
	SortOptAsc SortOption = iota
	SortOptDesc
)

// Sort can sort any strings
func Sort(ary []string, opt SortOption) []string {
	quickSort(ary, opt)
	return ary
}

// quickSort is sort using string array
func quickSort(ary []string, opt SortOption) {
	lenAry := len(ary)

	if lenAry < 2 {
		return
	}

	left := 0
	right := lenAry - 1

	pivot := rand.Int() % lenAry

	ary[pivot], ary[right] = ary[right], ary[pivot]

	for i, _ := range ary {
		switch opt {
		case SortOptAsc:
			if ary[i] < ary[right] {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		case SortOptDesc:
			if ary[i] > ary[right] {
				ary[left], ary[i] = ary[i], ary[left]
				left++
			}
		}
	}

	ary[left], ary[right] = ary[right], ary[left]

	quickSort(ary[:left], opt)
	quickSort(ary[left+1:], opt)
}
