package main

import "fmt"

func sort(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i],arr[j]
			}
		}
	}
}

func miniMaxSum(arr []int32) {
	min := 0
	max := 0

	sort(arr)

	for i := 0; i < len(arr); i++ {

		if i >= 1 {
			max += int(arr[i])
		}

		if i < len(arr)-1 {
			min += int(arr[i])
		}
	}

	fmt.Println(int64(min), int64(max))
}

func main() {
	arr := []int32{7, 69, 2, 221, 8974}
	miniMaxSum(arr)
}
