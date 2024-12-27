package main

func main() {
	person := make(map[string]string)
	person["name"] = "Muhammad Zikri"
	person["age"] = "23"
	person["address"] = "Indonesia"
	person["title"] = "Software Engineer"

	println(person["name"])
}
