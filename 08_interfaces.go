package main

import (
    f "fmt"
    "math"
    "errors"
)

func Interfaces() {
	f.Println("********** Interfaces **********")

	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangle{sideA: 3, sideB: 4, sideC: 5}

	rImpossible := rectangle{width: 3, height: -4}
	cImpossible := circle{radius: -5}
	tImpossible := triangle{sideA: 1, sideB: 1, sideC: 100}

	printShape(r)
	printShape(c)
	printShape(t)

	printShape(rImpossible)
	printShape(cImpossible)
	printShape(tImpossible)
}

func printShape(s shape) {
	f.Println(s)
	f.Print("Area: ")
	f.Println(s.area())
	f.Print("Perimeter: ")
	f.Println(s.perimeter())
}

type shape interface {
	area() (float64, error)
	perimeter() (float64, error)
	valid() error
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

func (r rectangle) valid() error {
	if r.width > 0 && r.height > 0 {
		return nil
	} else {
		return errors.New("Width and height of rectangle must be greater than 0")
	}
}

func (c circle) valid() error {
	if c.radius > 0 {
		return nil
	} else {
		return errors.New("Radius of circle must be greater than 0")
	}
}

func (t triangle) valid() error {
	p := (t.sideA + t.sideB + t.sideC) / 2.0
	if p > t.sideA && p > t.sideB && p > t.sideC {
		return nil
	} else {
		return errors.New("Sum of any two sides of a triangle must be greater than the third side")
	}
}

func (r rectangle) area() (float64, error) {
	if r.valid() != nil {
		return -1, r.valid()
	}
	return r.width * r.height, nil
}

func (c circle) area() (float64, error) {
	if c.valid() != nil {
		return -1, c.valid()
	}
	return math.Pi * c.radius * c.radius, nil
}

func (t triangle) area() (float64, error) {
	if t.valid() != nil {
		return -1, t.valid()
	}
	p := (t.sideA + t.sideB + t.sideC) / 2.0
	return math.Sqrt(p * (p - t.sideA) * (p - t.sideB) * (p - t.sideC)), nil
}

func (r rectangle) perimeter() (float64, error) {
	if r.valid() != nil {
		return -1, r.valid()
	}
	return 2 * (r.width + r.height), nil
}

func (c circle) perimeter() (float64, error) {
	if c.valid() != nil {
		return -1, c.valid()
	}
	return 2 * math.Pi * c.radius, nil
}

func (t triangle) perimeter() (float64, error) {
	if t.valid() != nil {
		return -1, t.valid()
	}
	return t.sideA + t.sideB + t.sideC, nil
}
