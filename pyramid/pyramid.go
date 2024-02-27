package main

import "fmt"

func main() {
	rows := 5
	space := rows - 1
	for i := 1; i <= rows; i++ {
		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}

		space--

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
