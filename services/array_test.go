package services

import (
	"fmt"
	"testing"
)

func learnArray(arr []int) int {
	arr[0] = 100
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		fmt.Println(arr[i])
	}
	return arrLength
}

func TestArrays(t *testing.T) {
	result := learnArray([]int{1, 2, 3, 4, 5}) // true

	fmt.Println("result ", result)
	expected := 5
    
    if result != expected {
        t.Errorf("Expected array length to be %d, but got %d", expected, result)
    }
}
