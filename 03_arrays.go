package main

import f "fmt"

func ArraysAndRanges() {
	a := variadicArrays("a", "quick", "brown", "fox")
	ranges(a)
}

func variadicArrays(a ...string) []string {
	f.Println("********** Arrays **********")

	b := make([]string, len(a))
	for i, s := range a {
		b[i] = s
	}
	f.Println("len(a)=", len(a), "; a=", a)
	f.Println("len(b)=", len(b), "; b=", b)

	b = append(b, "jumps", "over", "the", "lazy", "dog")
	f.Println("len(a)=", len(a), "; a=", a)
	f.Println("len(b)=", len(b), "; b=", b)
	f.Println("slice of last 3 of a=", a[len(a)-3:])
	f.Println("slice of last 3 of b=", b[len(b)-3:])
	return b
}

func ranges(a []string) {
	f.Println("********** Ranges **********")
	
	length := 0
	m := make(map[int]string)
	for i, s := range a {
		length += len(s)
		m[i] = s
	}
	f.Println("Map of strings:", m)
	f.Println("Total length of all strings =", length)

	f.Println("Characters in \"Hello Go!\"")
	for i,c := range "Hello Go!" {
		f.Println(" ", i, c)
	}
}
