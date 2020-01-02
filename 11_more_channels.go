package main

import (
	f "fmt"
	"time"
)

func MoreChannels() {
	channelSelector()
	timeouts()
	nonBlocking()
}

func channelSelector() {
	f.Println("********** Channel Selectors **********")

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

func timeouts() {
	f.Println("********** Timeouts **********")
	willTimeOut()
	willNotTimeOut()
}

func willTimeOut() {
	createAndWait("will timeout", 2, 1)
}

func willNotTimeOut() {
	createAndWait("will not timeout", 2, 3)
}

func createAndWait(msg string, creationWait time.Duration, timeoutWait time.Duration) {
	c2 := make(chan string, 1)
	go producer(c2, msg, creationWait)
	select {
	case res := <-c2:
		f.Println("received message: " + res)
	case <-time.After(timeoutWait * time.Second):
		f.Println("timed out for: " + msg)
	}
}

func nonBlocking() {
	f.Println("********** Non-blocking Channels **********")

	messages := make(chan string) // no buffer
	signals := make(chan string)  // no buffer
	select {
	case msg := <-messages:
		f.Println("received message", msg)
	default:
		f.Println("no message received")
	}

	msg := "aloha"
	select {
	case messages <- msg:
		f.Println("sent message", msg)
	default:
		f.Println("no message sent")
	}

	select {
	case msg := <-messages:
		f.Println("received message", msg)
	case sig := <-signals:
		f.Println("received signal", sig)
	default:
		f.Println("no activity")
	}
}
