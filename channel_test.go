package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Arif"
		fmt.Println("Selesai Mengirim Data Ke Channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Arif"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// ini channel untuk mengambil data ditandai dengan syntax chan<-
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Arif"
}

// ini channel untuk menerima data ditandai dengan syntax <-chan
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Arif"
		channel <- "Bukittinggi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	//anonymous function channel range menfirim data
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	// channel range menerima data
	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		default:
			fmt.Println("Sedang Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}

}
