package main

import "fmt"

func plusMinus(arr []int32) {

	zero := 0
	positive := 0
	negative := 0

	for _, eachArr := range arr {
		if eachArr == 0 {
			zero++
		} else if eachArr > 0 {
			positive++
		} else {
			negative++
		}
	}

	fmt.Printf("%.6f \n", float64(float64(positive)/float64(len(arr))))
	fmt.Printf("%.6f \n", float64(float64(negative)/float64(len(arr))))
	fmt.Printf("%.6f \n", float64(float64(zero)/float64(len(arr))))
}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
