package main

import "fmt"

// method of Type T will copy a instance??
// methhod of Type *t will not copy a instance when call?

type Point struct {
	X, Y int
}

//copy of object
func (p Point) ScaleBy(factor int) Point {
	p.X *= factor
	p.Y *= factor

	return p
}

//reference to object
func (p *Point) ScaleBy1(factor int) {
	p.X *= factor
	p.Y *= factor
}

//two copy of this object refer to the same underlying Point object?
type TestCopy struct {
	//reference
	refPoint *Point
}

func (T TestCopy) ScaleBy(factor int) {
	T.refPoint.X *= factor
	T.refPoint.Y *= factor
}

func main() {

	p1 := Point{1, 2}
	fmt.Printf("%v \n", p1)

	p1.ScaleBy(2)
	fmt.Printf("%v \n", p1)

	p1.ScaleBy1(3)
	fmt.Printf("%v \n", p1) //reference

	p3 := TestCopy{nil}
	p3.refPoint = &p1
	fmt.Printf("%v", p3)

	p3.ScaleBy(4)
	fmt.Printf("%v \n", p1)
}
