package main

import (
	"strconv"
)

func Convert(number int) string {
	var message string
	var a = strconv.Itoa(number)
	if number%3 == 0 {
		message += "Pling"
	}

	if number%5 == 0 {
		message += "Plang"
	} 
	
	if number%7 == 0 {
		message += "Plong"
	} 
	
	if message == "" {
		return a
	}
	return message
}

func main() {
	for number := 1; number < 106; number++ {
		Convert(number)
	}
}