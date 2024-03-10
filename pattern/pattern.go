// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Circle struct {
	JariJari float64
}

type Square struct {
	Panjang float64
	Lebar   float64
}
type Save interface {
	printArea() float64
}

func (c Circle) printArea() float64 {
	return c.JariJari * c.JariJari
}

func (s Square) printArea() float64 {
	return s.Lebar * s.Panjang
}

func main() {
	cirlce := Circle{JariJari: 4}
	fmt.Println(cirlce.printArea())

	square := Square{Panjang: 4, Lebar: 5}
	fmt.Println(square.printArea())
}
