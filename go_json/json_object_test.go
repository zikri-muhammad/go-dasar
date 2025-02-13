package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Name    string
	Age     int
	Married bool
	Hobbies []string
	Address []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestObject(t *testing.T) {
	customer := Customer{
		Name:    "Muhammad zikri",
		Age:     27,
		Married: true,
		Hobbies: []string{"code", "reading"},
		Address: []Address{
			{
				Street:     "Jalan yang pernah dilalui",
				Country:    "Indonesia",
				PostalCode: "09102",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
