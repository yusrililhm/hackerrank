package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	a := 2
	b := 5
	fmt.Println(solveMeFirst(uint32(a), uint32(b)))
}
