package equilateral_triangle_test

import (
	"fmt"
	"testing"
)

func TestEquilateralTriangleTest(t *testing.T) {
	input := 5

	for i := 1; i <= input; i++ {
		for j := 1; j <= input-i; j++ { // spasi segitiga kiri
			fmt.Print("  ")
		}
		for k := 1; k <= i; k++ { // segitiga siku kiri
			fmt.Print("* ")
		}
		for l := 1; l < i; l++ { // segitiga siku kanan
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
