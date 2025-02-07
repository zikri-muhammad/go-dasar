package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu ", totalCpu)

	totalTread := runtime.GOMAXPROCS(-1)
	fmt.Println("total tread ", totalTread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine ", totalGoroutine)
}