package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func Channels() {
	unbufferedChannel()
	bufferedChannel()
	synchronizedChannel()
	directionalChannels()
}

func unbufferedChannel() {
	f.Println("********** Unbuffered Channels **********")

	messages := make(chan string)
	f.Printf("Messages is of type '%T'\n", messages)

	f.Println("Seeking mesg at", time.Now().Format(time.StampMicro))
	go producer(messages, "poof!")

	msg := <-messages
	f.Println(msg)
	f.Println("Message rcvd at", time.Now().Format(time.StampMicro))
}

func producer(someChannel chan string, msg string, optionalWait ...time.Duration) {
	if len(optionalWait) == 0 {
		for r := 0; r != 6; {
			f.Print("Waiting for random length by rolling a dice until 6 appears ... ")
			time.Sleep(200 * time.Millisecond)
			rand.Seed(time.Now().UnixNano())
			r = rand.Int()%6 + 1
			if r == 6 {
				f.Printf("finally got %d!\n", r)
			} else {
				f.Printf("got %d\n", r)
			}
		}
	} else {
		time.Sleep(optionalWait[0] * time.Second)
	}

	f.Printf("Sent %s at %s\n", msg, time.Now().Format(time.StampMicro))
	someChannel <- msg
}

func bufferedChannel() {
	f.Println("********** Buffered Channels **********")

	messages := make(chan string, 2)
	messages <- "Buffered"
	messages <- "Channel"

	f.Println(<-messages)
	f.Println(<-messages)
}

func synchronizedChannel() {
	done := make(chan bool, 1)
	go worker(done)
	<-done // Commenting this out would likely not read all the stuff in the channel
}

func worker(done chan bool) {
	f.Print("Working ... ")
	time.Sleep(time.Second)
	f.Println("done!")

	done <- false // the actual value doesn't matter in this case: we could stuff 'true' into done and achieve the same result
}

func directionalChannels() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "ping pong message")
	pong(pings, pongs)
	f.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
