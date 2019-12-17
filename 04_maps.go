package main

import f "fmt"

func Maps() {
	f.Println("********** Maps **********")

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
