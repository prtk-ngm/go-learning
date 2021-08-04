package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{}
	tr := triangle{}

	sq.sideLength = 2
	tr.base = 2
	tr.height = 5

	printArea(sq)
	printArea(tr)

}
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
