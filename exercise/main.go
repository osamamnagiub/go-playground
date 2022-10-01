package main

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	println(s.getArea())
}

func main() {
	t := triangle{
		height: 5,
		base:   3,
	}

	s := square{
		sideLength: 5,
	}

	printArea(s)
	printArea(t)
}
