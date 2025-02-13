package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonRequest := `{"Name":"Muhammad zikri","Age":27,"Married":true}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestDecodeJsonArray(t *testing.T) {
	jsonArray := `[{"Street":"Jalan yang pernah dilalui","Country":"Indonesia","PostalCode":"09102"}]`
	jsonBytes := []byte(jsonArray)

	address := &[]Address{}
	json.Unmarshal(jsonBytes, address)

	fmt.Println(address)
}
