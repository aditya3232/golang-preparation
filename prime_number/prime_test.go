package prime_number

import (
	"fmt"
	"testing"
)

func IsPrime(input int) bool {
	if input <= 1 {
		return false
	}

	for i := 2; i*i <= input; i++ {
		if input%i == 0 {
			return false
		}
	}

	return true
}

func TestIsPrime(t *testing.T) {
	input := 20

	for i := 1; i <= input; i++ {
		if IsPrime(i) {
			fmt.Printf("Is Prime: %d\n", i)
		}
	}
}
