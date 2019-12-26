package main

import (
	f "fmt"
	"time"
)

func Routines() {
	f.Println("********** Routines **********")

	routine("direct")

	go routine("goroutine")

	go func(m string) {
		f.Println(m)
	}("anon")

	time.Sleep(time.Second)
	f.Println("done")
}

func routine(from string) {
	for i := 0; i < 3; i++ {
		f.Println(from, ":", i)
	}
}
