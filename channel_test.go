package main

import (
	"fmt"
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
