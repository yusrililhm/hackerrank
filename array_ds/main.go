package main

import "fmt"

func reverseArray(a []int32) []int32 {

	low := 0
	high := len(a) - 1

	for low < high {

		a[low], a[high] = a[high], a[low]

		low++
		high--
	}

	return a
}

func main() {
	a := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(a))
}
