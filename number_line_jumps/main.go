package main

import "fmt"

func kangaroo(x1, v1, x2, v2 int32) string {
	if ((x1-x2 >= 0 && v2-v1 >= 0) || (x1-x2 < 0 && v2-v1 < 0)) && (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	x1, v1, x2, v2 := 0, 3, 4, 2
	fmt.Println(kangaroo(int32(x1), int32(v1), int32(x2), int32(v2)))

	x1, v1, x2, v2 = 2, 1, 1, 2
	fmt.Println(kangaroo(int32(x1), int32(v1), int32(x2), int32(v2)))

	x1, v1, x2, v2 = 14, 4, 98, 2
	fmt.Println(kangaroo(int32(x1), int32(v1), int32(x2), int32(v2)))

	x1, v1, x2, v2 = 0, 2, 5, 3
	fmt.Println(kangaroo(int32(x1), int32(v1), int32(x2), int32(v2)))
}
