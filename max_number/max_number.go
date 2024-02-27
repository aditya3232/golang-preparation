package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 4, 2, 1, 5}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	fmt.Println("max number: ", numbers[0])
}
