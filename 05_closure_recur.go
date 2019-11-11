package main

import (
  f "fmt"
  t "time"
  "math/rand"
)

func main() {
	timeAndTemp()
	closures()
	recursive()
}

func timeAndTemp() {
	late, cold := stringAndInt()
	f.Printf("It's %d degrees at %s\n", cold, late)
}

func stringAndInt() (t.Time, int) {
	return t.Now().Local(), rand.Int() % 111
}

func closures() {
	init := 42
	r := ratchet(init)
	f.Printf("Ratcheting up from %d yields: %d, %d, %d\n",
		init, r(), r(), r())

	init = 1e5
	r = ratchet(init)	
	f.Printf("Ratcheting up from %d yields: %d, %d, %d\n",
		init, r(), r(), r())
}

func ratchet(n int) func() int {
	return func() int {
		n++ 
		return n
	}
}

func recursive() {
		f.Println("5!=", fact(5))
		f.Println("15!=", fact(15))
}

func fact(n int) int {
	if n == 1 {
		return 1
	} 
	return n * fact(n-1)
}
