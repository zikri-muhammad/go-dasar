package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups!")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println(number)
}

func TestSpeed(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Time taken:", time.Since(startTime))
}
