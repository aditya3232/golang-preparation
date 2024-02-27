package main

import "fmt"

// defer, panic, recover

func endApp() {
	fmt.Println("Aplikasi berhenti")
	message := recover()
	fmt.Println("Terjadi panic error: ", message)
}

func runApp(input bool) {
	defer endApp()
	if input {
		panic("ERROR")
	}

}

func main() {
	runApp(false)
	fmt.Println("+++Welcome+++")
}
