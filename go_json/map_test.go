package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonTag := `{"id":"OCU0001","name":"Apple Mac Pro M3","image_url":"http://example.com/image.jpg"}`
	jsonBytes := []byte(jsonTag)

	var product map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &product)

	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["name"])
	fmt.Println(product["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "OCU0001",
		"name":      "Apple Mac Pro M3",
		"image_url": "http://example.com/image.jpg",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
