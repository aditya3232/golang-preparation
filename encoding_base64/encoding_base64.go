package main

import (
	"encoding/base64"
	"fmt"
)

func EncodedBase64(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded
}

func DecodedBase64(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	return string(decoded)
}

func main() {
	input := "Joko"
	result := EncodedBase64(input)
	fmt.Println(result)

	// decoded dari hasil encoding
	decoded := DecodedBase64(result)
	fmt.Println(decoded)

}
