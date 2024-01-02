package main

import "fmt"

func aVeryBigSum(arr []int64) int64 {
	total := 0

	for _, eachArr := range arr {
		total += int(eachArr)
	}

	return int64(total)
}

func main() {
	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(arr))
}
