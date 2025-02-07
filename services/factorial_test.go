package services

import "testing"

func factorialLoop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func TestFactory(t *testing.T) {
	println(factorialLoop(5))
	println(factorialRecursive(5))
}
