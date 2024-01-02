package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {

	end := len(arr) - 1
	left := 0
	right := 0

	for i := 0; i < len(arr); i++ {
		left += int(arr[i][i])
		right += int(arr[i][end])
		end--
	}

	return int32(math.Abs(float64(left) - float64(right)))
}

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	fmt.Println(diagonalDifference(arr))
}
