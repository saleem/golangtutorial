package main

import (
    f "fmt"
    "time"
    "math/rand"
)

func producer(someChannel chan string) {
	for r := 0; r != 6; {
		f.Print("Rolling the dice ... ")
		rand.Seed(time.Now().UnixNano())
		r = rand.Int() % 6 + 1
		if r == 6 {
		  f.Printf("finally got %d!\n", r)
		} else {
			f.Printf("got %d\n", r)
		}
	}

	f.Println("Message sent at", time.Now().Format(time.StampMicro))
	someChannel <- "poof!"
}

func Channels() {
	f.Println("********** Channels **********")

	messages := make(chan string)
	f.Printf("Messages is of type '%T'\n", messages)

	f.Println("Seeking mesg at", time.Now().Format(time.StampMicro))
	go producer(messages)

	msg := <-messages
	f.Println(msg)
    f.Println("Message rcvd at", time.Now().Format(time.StampMicro))
}