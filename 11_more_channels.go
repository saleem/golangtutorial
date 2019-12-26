package main

import (
	f "fmt"
	// "math/rand"
	// "time"
)

func MoreChannels() {
	channelSelector()
}

func channelSelector() {
	c1 := make(chan string)
	c2 := make(chan string)

	go producer(c1, "one", 1)
	go producer(c2, "two", 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			f.Println("received", msg1)
		case msg2 := <-c2:
			f.Println("received", msg2)
		}
	}
}
