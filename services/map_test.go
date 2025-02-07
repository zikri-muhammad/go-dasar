package services

import "testing"

func TestMap(t *testing.T) {

	person := make(map[string]string)
	person["name"] = "Muhammad Zikri"
	person["age"] = "23"
	person["address"] = "Indonesia"
	person["title"] = "Software Engineer"

	println(person["name"])
}
