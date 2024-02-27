package main

import (
	"fmt"
)

func getCharacter(input string) []string {
	var slice []string

	for _, char := range input {
		slice = append(slice, string(char))
	}

	return slice
}

func main() {
	char := getCharacter("Budi")
	fmt.Println(char[0])
	fmt.Println(char[1])
	fmt.Println(char[2])
	fmt.Println(char[3])
}
