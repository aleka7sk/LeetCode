package main

import "fmt"

func main() {
	text := "55"
	fmt.Println(convertToInt((text)))
}

func convertToInt(text string) int {
	var number int
	for i := 0; i < len(text); i++ {
		number = number*10 + int(text[i]-48)
	}
	return number
}
