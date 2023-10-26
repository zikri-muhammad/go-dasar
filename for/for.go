package main

import "fmt"

func main() {
	person := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"} // AS maps

	for key, p := range person {
		fmt.Println(key, "adalah", p)
	}

	names := [3]string{"muhammad", "zikri", "naruto"} // as arrayclear
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}