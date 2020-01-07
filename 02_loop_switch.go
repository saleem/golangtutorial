package main

import (
	f "fmt"
	t "time"
)

func LoopAndSwitch() {
	looper()
	switcheroo()
}

func looper() {
	f.Println("********** For Loop **********")

	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			f.Print("fizz")
		}
		if i%5 == 0 {
			f.Print("buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			f.Print(i)
		}
		f.Println()
	}
}

func switcheroo() {
	f.Println("********** Switch Case **********")

	switch t.Now().Weekday() {
	case t.Saturday, t.Sunday:
		f.Println("Hooray for the weekend!")
	default:
		f.Println("Back to the weekly grind...")
	}
}
