package main

import (
	"errors"
	"fmt"
)

/*
- variable errors.New harus diawali dengan Err
- jika error tidak nil, maka akan mengecek error apa
*/

var (
	ErrValidationError = errors.New("validation error")
	ErrNotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ErrValidationError
	}

	if id != "eko" {
		return ErrNotFoundError
	}

	return nil
}

func main() {
	err := GetById("budi")
	if err != nil {
		if errors.Is(err, ErrValidationError) {
			fmt.Println(err)
		} else if errors.Is(err, ErrNotFoundError) {
			fmt.Println(err)
		} else {
			fmt.Println("unknown error")
		}
	}
}
