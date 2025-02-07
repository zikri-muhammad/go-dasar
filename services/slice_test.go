package services

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	names := []string{"Muhammad", "Zikri", "Jane", "Doe", "naruto", "sasuke"}
	slice := names[4:6]

	days := append(slice, "sakura")

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(days)
	fmt.Println(names)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Zikri"
	// newSlice[2] = "Jane"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println("before copy", fromSlice)
	fmt.Println("before copy to slice", toSlice)

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
