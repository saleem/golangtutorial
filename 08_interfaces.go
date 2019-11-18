package main

import (
    f "fmt"
    "math"
)

type shape interface {
	area() float64
	perimeter() float64
}

func main() {
	printShapes()
}

func printShapes() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangle{sideA: 3, sideB: 4, sideC: 5}
	printShape(r)
	printShape(c)
	printShape(t)
}

func printShape(s shape) {
	f.Println(s)
	f.Printf("Area: %.2f, Perimeter: %.2f\n", s.area(), s.perimeter())
}


type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	sideA, sideB, sideC float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	p := t.perimeter() / 2.0
	return math.Sqrt(p * (p - t.sideA) * (p - t.sideB) * (p - t.sideC))
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}