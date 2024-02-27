package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 4, 5, 10, 3, 20, 9}
	users := []string{"adit", "diqi", "john", "rudy"}

	fmt.Println(slices.Max(users))
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Contains(numbers, 6))
	fmt.Println(slices.Contains(users, "adit"))
	fmt.Println(slices.Index(numbers, 2))

}
