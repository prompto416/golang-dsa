package main

import "fmt"

//Time Complexity: BigO(n)
func ReverseArray(lst []int) []int {

	for i := 0; i < len(lst)/2; i++ {
		fmt.Println(i,len(lst)-1-i)
		lst[i], lst[len(lst)-1-i] = lst[len(lst)-1-i], lst[i]
	}

	return lst
}