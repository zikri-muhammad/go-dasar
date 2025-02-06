package main

import "fmt"

func LearnString(str string) int {
	strLength := len(str)
	for i := 0; i < strLength; i++ {
		fmt.Println(str[i])
	}
	return strLength
}

func main() {
	fmt.Println(LearnString("Hello World"))    // true
	fmt.Println(LearnString("Hello Muhammad")) // true
}
