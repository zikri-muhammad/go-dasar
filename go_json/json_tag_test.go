package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "OCU0001",
		Name:     "Apple Mac Pro M3",
		ImageUrl: "http://example.com/image.jpg",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(r *testing.T) {
	jsonTag := `{"id":"OCU0001","name":"Apple Mac Pro M3","image_url":"http://example.com/image.jpg"}`
	jsonBytes := []byte(jsonTag)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
