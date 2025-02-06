package main

func LearnArray(arr []int) int {
	arr[0] = 100
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		println(arr[i])
	}
	return arrLength
}

func main() {
	LearnArray([]int{1, 2, 3, 4, 5}) // true
}
