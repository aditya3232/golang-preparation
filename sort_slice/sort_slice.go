package main

import (
	"fmt"
	"sort"
)

/*
- sort.Ints(var) mengurutkan slice secara ascending.
- sort.Slice(var, func(i, j int) bool{}) mengurutkan slice secara descending,
- jika elemen pertama lebih besar dari elemen kedua, maka elemen pertama ditempatkan sebelum elemen kedua.
*/

func main() {
	numbers := []int{5, 1, 3, 2, 4}

	sort.Ints(numbers)
	fmt.Println("ascending :", numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("ascending :", numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println("descending: ", numbers)

}
