package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Azzin" // input data to channel

	// data := <- channel // input channel to data

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Azzin Maharil"
		fmt.Println("Mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Azzin Maharil"
}

func TestChannelAsParameter(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	// channel <- "Azzin" // input data to channel

	// data := <- channel // input channel to data

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

// initial <- after chan, the mean is in send only
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Azzin Maharil"
}

//initial <- before chan, example like <-chan the mean is recive only
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Azzin" // input data to channel

	// data := <- channel // input channel to data

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Azzin"
	channel <- "Maharil"
	channel <- "Rath"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	//Anonymous func
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke" + strconv.Itoa(i)
		}
	}()
	close(channel) // must close or deadlock

	for data := range channel {
		fmt.Println("Menerima data ke", data)
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data channel ke 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data channel ke 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
