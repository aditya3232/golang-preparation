package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Adi")
	data.PushBack("Budi")
	data.PushBack("Cintya")
	data.PushBack("Dodit")

	i := 1
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Printf("data antrian %d: %s\n\n", i, e.Value)
		i++
	}

	head := data.Front()
	fmt.Println("data pertama: ", head.Value)

	next := head.Next()
	fmt.Println("data kedua: ", next.Value)

	next = next.Next()
	fmt.Println("data ketiga: ", next.Value)
}
