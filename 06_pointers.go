package main

import f "fmt"

func main() {
  recursive()
  pointer()
}

func recursive() {
	f.Println("Recursive function: ")
	f.Println("  5!=", factR(5))
	f.Println("  15!=", factR(15))
}

func factR(n int) int {
	if n == 1 {
		return 1
	} 
	return n * factR(n-1)
}

func pointer() {
	f.Println("Recursive with pointers: ")
  	r := 1
	factP(5, &r)
	f.Println("  5!=", r)
	r = 10
	factP(15, &r)
	f.Println("  15!=", r)
}

func factP(i int, rptr *int) {
  if i == 1 {
  	return
  } else {
  	*rptr *= i
  	factP(i-1, rptr)
  }
}
