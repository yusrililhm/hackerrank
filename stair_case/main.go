package main

import "fmt"

func stairCase(n int32) {
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			if i < int(n-1-int32(j)) {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	stairCase(4)
}
