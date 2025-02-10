package go_contaxt

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

type DataSiswa struct {
	Name string
	Nip string

}

func TestContaxt(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b")) // nil becouse value just search to parent 

	fmt.Println(contextA.Value("b")) // nil becouse value just search to parent 
}


func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	// go func() { // terjadi leak 
	// 	defer close(destination)
	// 	counter := 1
	// 	for {
	// 		destination <- counter
	// 		counter++
	// 	}
	// }()

	go func() { 
		defer close(destination)
		counter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Totol Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Totol Goroutine", runtime.NumGoroutine())
}