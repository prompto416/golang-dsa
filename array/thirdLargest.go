package main

import (
	"sort"
)


func thirdLargest(lst []int) int {
	

	sort.Ints(lst)
	lastIndex := len(lst)

	return lst[lastIndex-2]
}

