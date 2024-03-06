package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) { // json -> golang
	reader, _ := os.Open("user.json")
	decoder := json.NewDecoder(reader) // parameter json.NewEncoder() bisa baca stream dari file, network, atau request body
	user := User{}

	for {
		err := decoder.Decode(&user)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("error while read file: %s", err.Error())
		}
	}

	fmt.Println(user)
	fmt.Println(user.FirstName)
	fmt.Println(user.MiddleName)
	fmt.Println(user.LastName)

}

func TestStreamEncoder(t *testing.T) { // golang -> json
	writer, err := os.OpenFile("user_encoder.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Open the file in append mode or create it if it doesn't exist
	if err != nil {
		fmt.Println("error write", err)
	}

	defer writer.Close()

	encoder := json.NewEncoder(writer)

	user := User{
		FirstName:  "Aska",
		MiddleName: "Braha",
		LastName:   "Surasa",
	}

	err = encoder.Encode(user)
	if err != nil {
		t.Fatalf("Error while encode stream data JSON: %s", err.Error())
	}

}
