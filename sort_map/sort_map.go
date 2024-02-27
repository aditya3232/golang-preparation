package main

import (
	"fmt"
	"sort"
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
