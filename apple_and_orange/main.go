package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	orange := 0
	apple := 0

	for i := range apples {
		apples[i] += a
	}

	for i := range oranges {
		oranges[i] += b
	}

	for _, eachApple := range apples {
		if eachApple <= t && eachApple >= s {
			apple++
		}
	}

	for _, eachOrange := range oranges {
		if eachOrange <= t && eachOrange >= s {
			orange++
		}	
	}

	fmt.Println(apple)
	fmt.Println(orange)
}

func main() {
	s, t, a, b, appless, oranges := 7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6}
	countApplesAndOranges(int32(s), int32(t), int32(a), int32(b), appless, oranges)
}
