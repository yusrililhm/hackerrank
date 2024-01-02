package main

import "fmt"

func compareTriplets(a []int, b []int) []int {

	pointA := 0
	pointB := 0

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			pointB++
		} else if a[i] > b[i] {
			pointA++
		}
	}

	return []int{pointA, pointB}
}

func main() {
	fmt.Println(compareTriplets([]int{17, 28, 30}, []int{99, 16, 8}))
}
