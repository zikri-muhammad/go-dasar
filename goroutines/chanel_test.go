package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)


	

func TestChanel(t *testing.T) {
	chanel := make(chan string) // deklarasi chanel
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "Hello Muhammad zikri"
		fmt.Println("Done")
	}()

	message := <- chanel
	fmt.Println(message)

	time.Sleep(5 * time.Second)
}


func SendMessage(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Hello Muhammad zikri"
}

func TestChanelWithParameter(t *testing.T) {
	chanel := make(chan string)
	go SendMessage(chanel)
	message := <- chanel
	fmt.Println(message)
}

func OnlyOut(chanel <-chan string) {
	message := <- chanel
	fmt.Println(message)
}

func OnlyIn(chanel chan<- string) {
	time.Sleep(2 * time.Second)
	chanel <- "Hello Muhammad zikri"
}

func TestChanelInOut(t *testing.T) {
	chanel := make(chan string)
	go OnlyIn(chanel)
	go OnlyOut(chanel)

	time.Sleep(5 * time.Second)
}

// chanel buffered
func TestChanelBuffered(t *testing.T) {
	chanel := make(chan string, 3) // 3 adalah jumlah buffer
	defer close(chanel)

	chanel <- "Hello"
	chanel <- "World"
	chanel <- "Muhammad zikri"

	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}


// chanel range
func TestChanelRange(t *testing.T) {
	chanel := make(chan string)
	// defer close(chanel)

	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "Hello Muhammad zikri " + strconv.Itoa(i)
		}
		close(chanel)
	}()

	for message := range chanel {
		fmt.Println(message)
	}

	fmt.Println("Done")
}


func TestSelectChanel(t *testing.T) {
	channel := make(chan string)
	channel1 := make(chan string)
	defer close(channel)
	defer close(channel1)


	go SendMessage(channel)
	go SendMessage(channel1)

	counter := 0
	for {
		select {
		case data := <- channel:
			fmt.Println("Data dari chanel 1", data)
			counter++
		case data := <- channel1:
			fmt.Println("Data dari chanel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChanel(t *testing.T) {
	channel := make(chan string)
	channel1 := make(chan string)
	defer close(channel)
	defer close(channel1)


	go SendMessage(channel)
	go SendMessage(channel1)

	counter := 0
	for {
		select {
		case data := <- channel:
			fmt.Println("Data dari chanel 1", data)
			counter++
		case data := <- channel1:
			fmt.Println("Data dari chanel 2", data)
			counter++
		default:
			fmt.Println("Menunngu Data lek ku")
		}

		if counter == 2 {
			break
		}
	}
}
