package main

func Fibonaci(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonaci(n-1) + Fibonaci(n-2)
}