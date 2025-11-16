package main

import (
	"fmt"
	"math"

	"sort"
)

func main() {
	res1 := secondLargestOnlogn([]int{12, 35, 1, 10, 34, 1})
	fmt.Println(res1,"hi")


	res2 := secondLargestOn([]int{10, 5, 10})
	fmt.Println(res2)
}



func secondLargestOnlogn(lst []int) int {
	

	sort.Ints(lst)
	lastIndex := len(lst)

	return lst[lastIndex-2]
}


func secondLargestOn(lst []int) int { 

	largest := math.MinInt
	secondLargest := math.MinInt


	for _, num := range lst{
		if num > largest { 
			secondLargest = largest 
			largest = num 
		} else if num > secondLargest && num < largest { 
			secondLargest = num 
		} 
	}
		
	

	return secondLargest
}