package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Hobbies    []string  `json:"hobbies"`   // ini slice
	Addresses  []Address `json:"addresses"` // kalau pakai array otomatis format banyak object
}

type Address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

func TestEncodeJsonComplex(t *testing.T) {
	user := User{
		FirstName:  "Eko",
		MiddleName: "Budi",
		LastName:   "Suseno",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jl Kebenaran",
				Country:    "Indonesia",
				PostalCode: "33512",
			},
			{
				Street:     "Jl Lurus",
				Country:    "Indonesia",
				PostalCode: "7751",
			},
		},
	}

	jsonByte, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Error while marshaling JSON: %s", err.Error())
	}

	fmt.Println(string(jsonByte))
}

func TestDecodeJsonComplex(t *testing.T) {
	jsonStringAddress := `[{"street":"Jl Kebenaran","country":"Indonesia","postal_code":"33512"},{"street":"Jl Lurus","country":"Indonesia","postal_code":"7751"}]`
	jsonByte := []byte(jsonStringAddress)
	address := []Address{}

	err := json.Unmarshal(jsonByte, &address)
	if err != nil {
		t.Fatalf("Error while unmarshaling JSON: %s", err.Error())
	}

	fmt.Println(address)

}

func TestEncodeJsonMap(t *testing.T) {
	product := map[string]interface{}{
		"id":           "P0001",
		"product_name": "Radeon XFX",
		"price":        8500000,
	}

	jsonByte, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("Error while marshaling JSON: %s", err.Error())
	}

	fmt.Println(string(jsonByte))
}

func TestDecodeJsonMap(t *testing.T) {
	jsonString := `{"id":"P0002","price":9000000,"product_name":"Intel i7 13000"}`
	jsonByte := []byte(jsonString)
	var product map[string]interface{}

	err := json.Unmarshal(jsonByte, &product)
	if err != nil {
		t.Fatalf("Error while unmarshaling JSON: %s", err.Error())
	}

	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["price"])
	fmt.Println(product["product_name"])
}
