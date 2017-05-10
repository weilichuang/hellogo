package main

import (
	"fmt"
)

func main() {
	p := Point{X: 100, Y: 100}
	fmt.Printf("%#v\n", p)

	p3 := Point3{Point: Point{X: 100., Y: 100.}, Z: 200.}
	fmt.Printf("%#v\n", p3)

	p4 := Point4{Point3: Point3{Point: Point{X: 100., Y: 100.}, Z: 200.}, W: 300.}
	fmt.Printf("%#v\n", p4)

	fmt.Println(p.X)
	fmt.Println(p3.Y)
	fmt.Println(p4.Z)
}

type Point struct {
	X, Y float64
}

type Point3 struct {
	Point
	Z float64
}

type Point4 struct {
	Point3
	W float64
}
