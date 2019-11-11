	package main

import (
  f "fmt"
  t "time"
  "math"
)

func main() {
	dataTypes()
	looper()
	switcheroo()
	arrays()
	maps()
}

func dataTypes() {
	const h = "Hello, " + "World!"
	f.Println(h)

	const y, z = 7/3.0, 7.0/4
    f.Println("7/3.0 =", y)
    f.Println("7.0/4 =", z)

    const i, j, k, l, m = 1+10, 7/3, 7%3, 7^3, 7*3
	f.Println("1+10 =", i)
	f.Println("7/3 =", j)
	f.Println("7%3 =", k)
	f.Println("7^3 =", l)
	f.Println("7*3 =", m)

    const b, c, d =  true && false, true || false, !true
	f.Println("true AND false is", b)
	f.Println("true OR false is", c)
	f.Println("not true is", d)

	// Cannot be defined as a constant
	xkcd217 := math.Exp(math.Pi) - math.Pi
	f.Println("www.xkcd.com/217 punchline:", xkcd217)
}

func looper() {
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
	switch t.Now().Weekday() {
	case t.Saturday, t.Sunday:
		f.Println("Hooray for the weekend!")
	default:
		f.Println("Back to the weekly grind...")
	}
}

func arrays() {
	a := []int{10, 20, 30, 40, 50} // string{"a", "quick", "brown"}
	b := make([]int, len(a))
	copy(b, a)
	f.Println("len(a)=", len(a), "; a=", a)
	f.Println("len(b)=", len(b), "; b=", b)

	a = append(a, 60, 70, 80)
	f.Println("len(a)=", len(a), "; a=", a)
	f.Println("len(b)=", len(b), "; b=", b)
	f.Println("slice of last 3 of a=", a[len(a)-3:])
	f.Println("slice of last 3 of b=", b[len(b)-3:])
}

func maps() {
	m := make(map[string]int)
	m["a"] = 9
	m["b"] = 3
	m["x"] = 1
	m["y"] = 4
	m["z"] = 1
	m["0"] = 0
	f.Println("len(m)=", len(m), "; m=", m)
	delete(m, "ä") // non-existent key
	delete(m, "z")
	f.Println("len(m)=", len(m), "; m=", m)
	examine(m, "0")
    examine(m, "ë")
}

func examine(m map[string]int, s string) {
	v, ok := m[s]
    if ok {
		f.Println(s, "is present in map with a value of", v)
    } else {
		f.Println(s, "is not present in map")
    }
}
