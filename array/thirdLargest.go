package main

import (
	"math"
)

//Time Complexity: BigO(n)
func thirdLargest(lst []int) int {
	

	largest := math.MinInt
	secondLargest := math.MinInt
	thirdLargest := math.MinInt



	for _, num := range lst{
		if num > largest { 
			if secondLargest > thirdLargest{ 
				thirdLargest = secondLargest
				secondLargest = largest 
			} else if thirdLargest > secondLargest { 
				secondLargest = thirdLargest
				thirdLargest = largest
			}
			
			largest = num 
		} else if num > secondLargest && num < largest { 

			thirdLargest = secondLargest
			secondLargest = num 
		} else if num > thirdLargest && num < largest { 
			thirdLargest = num
		}
	}

	if thirdLargest != math.MinInt  { 
		return thirdLargest
	} else { 
		return -1
	}

}

