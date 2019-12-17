package main

import (
  f "fmt"
  "math"
)

func DataTypes() {
	f.Println("********** DataTypes **********")
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
