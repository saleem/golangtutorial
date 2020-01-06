package main

import (
	f "fmt"
	"strconv"
	"time"
)

func MoreChannels() {
	channelSelector()
	timeouts()
	nonBlocking()
	closeChannels()
	rangeOverChannels()
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

func closeChannels() {
	f.Println("********** Closing Channel **********")

	jobs := make(chan string, 5)
	done := make(chan bool)

	go receiver(jobs, done)

	for j := 1; j <= 3; j++ {
		producer(jobs, strconv.Itoa(j), 0)
	}
	close(jobs)
	f.Println("sent all jobs")

	<-done
}

func rangeOverChannels() {
	f.Println("********** Range Over Channels **********")

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "buckle my shoe"
	close(queue)

	for elem := range queue {
		f.Printf("Received %s\n", elem)
	}

}

func receiver(from chan string, notifyOnDone chan bool) {
	for {
		j, more := <-from
		if more {
			f.Printf("Rcvd %s at %s\n", j, time.Now().Format(time.StampMicro))
		} else {
			f.Println("Rcvd all messages")
			notifyOnDone <- true
			return
		}
	}
}
