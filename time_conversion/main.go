package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	time := ""

	if s[len(s)-2:] == "AM" && s[:2] == "12" {
		time = "00" + s[2:len(s)-2]
	} else if s[len(s)-2:] == "AM" {
		time = s[:len(s)-2]
	} else if s[len(s)-2:] == "PM" && s[:2] == "12" {
		time = s[:len(s)-2]
	} else {
		hour, _ := strconv.Atoi(s[:2])
		hourToString := strconv.Itoa(hour + 12)

		time = hourToString + s[2:len(s)-2]
	}

	return time
}

func main() {
	s := "12:01:00AM"
	fmt.Println(timeConversion(s))

	s = "12:01:00PM"
	fmt.Println(timeConversion(s))

	s = "01:01:00PM"
	fmt.Println(timeConversion(s))
}
