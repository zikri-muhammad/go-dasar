package main

type Backlist func(string) bool

func registerUser(name string, backlist Backlist) {
	nameFiltered := backlist(name)
	if nameFiltered {
		println("You are backlisted", name)
	} else {
		println("Welcome", name)
	}
}

func main() {
	backlist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", backlist)
	registerUser("Muhammad", backlist)
}
