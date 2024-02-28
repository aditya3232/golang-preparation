package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Name    string
	Age     int
	Married bool
}

func logJson(data any) string {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func TestEncode(t *testing.T) {
	result_1 := logJson(1)
	expected_1 := "1"
	if result_1 != expected_1 {
		t.Errorf("result_1 returned %s, expected %s\n", result_1, expected_1)
	}
	fmt.Println(logJson(1))

	result_2 := logJson("Adit")
	expected_2 := "\"Adit\"" // Adding escaped quotes to the expected string
	if result_2 != expected_2 {
		t.Errorf("result_2 returned %s, expected %s\n", result_2, expected_2)
	}
	fmt.Println(logJson("Adit"))

	result_3 := logJson(true)
	expected_3 := "true"
	if result_3 != expected_3 {
		t.Errorf("result_3 returned %s, expected %s\n", result_3, expected_3)
	}
	fmt.Println(logJson(true))

	result_4 := logJson([]string{"adit", "diqi", "budi", "arya"})
	expected_4 := `["adit","diqi","budi","arya"]`
	if result_4 != expected_4 {
		t.Errorf("result_4 returned %s, expected %s\n", result_4, expected_4)
	}
	fmt.Println(logJson([]string{"adit", "diqi", "budi", "arya"}))

}

func TestEncodeJSON(t *testing.T) {
	customer := Customer{
		Name:    "Muhammad Aditya",
		Age:     27,
		Married: false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Error while marshaling JSON: %s", err.Error())
	}

	result_5 := string(bytes)
	expected_5 := `{"Name":"Muhammad Aditya","Age":27,"Married":false}`
	if result_5 != expected_5 {
		t.Errorf("result_5 returned %s, expected %s", result_5, expected_5)
	}

	fmt.Println(result_5)
}

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Name":"Ichsan Ashiddiqi","Age":24,"Married":false}`
	jsonByte := []byte(jsonString)
	customer := Customer{}

	err := json.Unmarshal(jsonByte, &customer)
	if err != nil {
		t.Fatalf("Error while unmarshaling JSON: %s", err.Error())
	}

	fmt.Println(customer)
}
