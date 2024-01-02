package main

import "fmt"

func findMax(candles []int32) int32 {
	max := 0

	for _, eachCandle := range candles {
		if eachCandle > int32(max) {
			max = int(eachCandle)
		}
	}

	return int32(max)
}

func birthdayCakeCandles(candles []int32) int32 {
	max := findMax(candles)
	candle := 0

	for _, eachCandle := range candles {
		if eachCandle == max {
			candle++
		}
	}

	return int32(candle)
}

func main() {
	candles := []int32{4, 4, 1, 3}
	fmt.Println(birthdayCakeCandles(candles))
}
