package services

import (
	"fmt"
	"testing"
)

func learnString(str string) int {
	strLength := len(str)
	for i := 0; i < strLength; i++ {
		fmt.Println(str[i])
	}
	return strLength
}

func TestStrings(t *testing.T) {

	fmt.Println(learnString("Hello World"))    // true
	fmt.Println(learnString("Hello Muhammad")) // true
}
