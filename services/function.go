package main

// if we want to pass a function as a parameter, we need to define the function signature
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	println("Hello", nameFiltered)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Muhammad", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
