package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perim() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return math.Pi * 2 * c.radius
}


func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func measure(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	measure(r)
	measure(c)
}

