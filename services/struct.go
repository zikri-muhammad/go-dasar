package main

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	println("Hello", name, "My name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Muhammad"
	customer.Address = "Indonesia"
	customer.Age = 20

	customer2 := Customer{
		Name:    "Muhammad",
		Address: "Indonesia",
		Age:     20,
	}

	customer3 := Customer{"Muhammad", "Indonesia", 20}

	println(customer.Name)
	println(customer.Address)
	println(customer.Age)

	println(customer2.Name)
	println(customer2.Address)
	println(customer2.Age)

	println(customer3.Name)
	println(customer3.Address)
	println(customer3.Age)

	customer.sayHello("Zikri")
}
