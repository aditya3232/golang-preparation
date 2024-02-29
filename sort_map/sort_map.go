package main

import (
	"fmt"
	"sort"
	"testing"
)

/*
- list merupakan slice dari map
- map[string]string => kumpulan data key & value
- []string => kumpulan data slice
- [...]string => kumpulan data array
*/

func main() {
	list := []map[string]string{
		{"id": "1", "name": "John Doe", "address": "Kemang"},
		{"id": "2", "name": "Sarah Mien", "address": "Tebet"},
		{"id": "3", "name": "Michele Kyle", "address": "Kemang"},
		{"id": "4", "name": "Dior Angelie", "address": "Setiabudi"},
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i]["name"] < list[j]["name"]
	})
	fmt.Println("Sorted by name (ascending): ", list)

	sort.Slice(list, func(i, j int) bool {
		return list[i]["id"] > list[j]["id"]
	})
	fmt.Println("Sorted by name (descending): ", list)

}

type User struct {
	Id   string
	Nama string
}

func SortSliceMap(input []User) []User {
	sort.Slice(input, func(i, j int) bool {
		return input[i].Id > input[j].Id
	})

	return input
}

func TestSortSliceMap(t *testing.T) {
	user := []User{
		{"1", "Muhammad Aditya"},
		{"2", "Ichsan Ashiddiqi"},
	}

	result := SortSliceMap(user)

	fmt.Println(result)
}
