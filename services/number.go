package main

func AmIWilson(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	fact := 1
	mod := n * n
	for i := 2; i < n; i++ {
		fact = (fact * i) % mod
	}
	return (fact+1)%mod == 0
}

// func main() {
// 	fmt.Println(len(LearnString())) // true
// 	// fmt.Println(AmIWilson(563))     // false
// 	// fmt.Println(AmIWilson(7))       // false
// 	// fmt.Println(AmIWilson(8))       // false
// }
