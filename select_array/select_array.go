package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	selectQueue := numbers[:2]

	fmt.Println(selectQueue)
}
