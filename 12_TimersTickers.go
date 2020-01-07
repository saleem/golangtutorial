package main

import (
	f "fmt"
	"time"
)

func TimersAndTickers() {
	timers()
	tickers()
}

func timers() {
	f.Println("********** Timers **********")

	t1 := time.NewTimer(2 * time.Second)

	<-t1.C
	f.Printf("T1 fired at %s\n", time.Now().Format(time.StampMicro))

	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		f.Printf("T2 fired at %s\n", time.Now().Format(time.StampMicro))
	}()

	s2 := t2.Stop()

	if s2 {
		f.Printf("T2 stopped at %s\n", time.Now().Format(time.StampMicro))
	}

	time.Sleep(2 * time.Second)
}

func tickers() {
	f.Println("********** Tickers **********")

	tick := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tick.C:
				f.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	tick.Stop()
	done <- true
	f.Printf("Ticker stopped at %s\n", time.Now().Format(time.StampMicro))
}
