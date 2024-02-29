// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"testing"
)

type Userx struct {
	Id   string
	Nama string
}

func SortSliceStruct(input []Userx) []Userx {
	sort.Slice(input, func(i, j int) bool {
		return input[i].Id > input[j].Id
	})

	return input
}

func TestSortSliceStruct(t *testing.T) {
	userx := []Userx{
		{"1", "Muhammad Aditya"},
		{"2", "Ichsan Ashiddiqi"},
	}

	result := SortSliceStruct(userx)

	fmt.Println(result)
}
