package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	sq := square{
		sideLength: 2,
	}

	tr := triangle{
		height: 2,
		base:   3,
	}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("Area :", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength

}
