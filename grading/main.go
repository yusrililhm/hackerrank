package main

import (
	"fmt"
	"math"
)

func round(n float64) float64 {
	if int(n)%5 == 0 {
		return n
	}

	return math.Ceil(n/5) * 5
}

func gradingStudents(grades []int32) []int32 {

	for i := 0; i < len(grades); i++ {
		if int32(round(float64(grades[i])))-grades[i] < 3 && grades[i] >= 38 {
			grades[i] = int32(round(float64(grades[i])))
		}
	}

	return grades
}

func main() {
	grades := []int32{
		86,
		30,
		0,
		16,
		51,
		53,
		42,
		48,
		22,
		69,
		12,
		27,
		34,
		24,
		95,
		16,
		32,
		22,
		52,
		56,
		71,
		95,
	}
	fmt.Println(gradingStudents(grades))
}
