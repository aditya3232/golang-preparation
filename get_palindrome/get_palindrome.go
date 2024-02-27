package main

import (
	"fmt"
	"strings"
)

/*
- strings.ToLower(strings.ReplaceAll(input, " ", "")) => menghilangkan spasi, dan ganti ke lower
- for i := 0; i <= len(inputString)/2; i++ => ambil setengah perulangan dari character, misal malam, panjang 5 kurang dari 5/2 = 2, maka ambil ma
- inputString[len(inputString)-1-i] => ambil 2 string dari belakang, karena perulangan pertama adalah index 0, 0 - 1, maka akan ambil dari belakang
- lalu bandingkan 2 string pertama dan 2 string dibelakang
*/

func getPalindrome(input string) bool {
	inputString := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	for i := 0; i <= len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(getPalindrome("Budi"))
	fmt.Println(getPalindrome("malam"))
}
