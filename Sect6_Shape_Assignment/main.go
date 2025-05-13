package main

import "fmt"

// Interface that will generate an area value for a given shape
type shape interface {
	getArea() float64
}

// Creating square structure
type square struct{
	sideLength float64
}

// Creating triangle structure
type triangle struct{
	height float64
	base float64
}

// Main function to create sq and tr variable with given dimensions
// Will use printArea to calculate and print out area value
func main() {
	sq := square{sideLength: 10}
	tr := triangle{height: 10, base: 10}

	printArea(sq)
	printArea(tr)

}

// Will use getArea and shape's calculation to determine area
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// Area calculation for square
func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

// Area calculation for triangle
func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}