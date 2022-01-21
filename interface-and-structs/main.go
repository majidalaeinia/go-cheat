package main

import "fmt"

func main() {
	var cat, parrot Pet
	cat = Cat{
		name:     "Jack",
		tailSize: 10,
	}
	parrot = Parrot{
		name:  "Tom",
		color: "Green",
	}
	DescribeAll(cat, parrot)
}

type Pet interface {
	Describe()
}

type Cat struct {
	name     string
	tailSize int
}

type Parrot struct {
	name  string
	color string
}

func (c Cat) Describe() {
	fmt.Printf("Cat -> name: %s, tail size: %d.\n", c.name, c.tailSize)
}

func (p Parrot) Describe() {
	fmt.Printf("Parrot -> name: %s, color: %s.\n", p.name, p.color)
}

func DescribeAll(first Pet, second Pet) {
	first.Describe()
	second.Describe()
}
