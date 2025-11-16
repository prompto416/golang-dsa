package main

import (
	"math"

	"sort"
)




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

	if secondLargest != math.MinInt { 
		return secondLargest
	} else { 
		return -1
	}



}	
	

	
