package main

func ReverseArrayInGroup(lst []int, k int) []int {
	i := 0
	n := len(lst)

	for i < k {
		left := i
		right := min(i+k-1, n-1)

		for left < right {
			lst[left], lst[right] = lst[right], lst[left]
			left += 1
			right -= 1
		}

		i += k
	}

	return lst
}