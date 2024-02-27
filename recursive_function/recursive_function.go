package main

import "fmt"

/*
- result = result * i
- result = 4 * 1 * 3 * 1 * 2 * 1 * 1 * 1
*/
func factorialLoop(input int) int {
	result := 1

	for i := input; i > 0; i-- {
		result *= i
	}

	return result
}

/*
- misal input 4
- 4 bukan 1 maka masuk else, jadi 4 * factorialRecursiceFunction(4-1)
- input sebelumnya dikurangi 1, jadi 3 * factorialRecursiceFunction(3-1)
- next, 2 * factorialRecursiceFunction(2-1)
- next, inputnya 1 maka return 1
*/
func factorialRecursiceFunction(input int) int {
	if input == 1 {
		return 1
	} else {
		return input * factorialRecursiceFunction(input-1)
	}
}

func main() {
	fmt.Println(factorialLoop(4))
	fmt.Println(factorialRecursiceFunction(10))
}
