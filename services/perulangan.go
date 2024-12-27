package main

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for index, day := range days {
		println(index, day)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println("perulangan ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println("perulangan ke", i)
	}
}
