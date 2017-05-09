package main

import (
	"fmt"
)

func main() {
	a, b := 100, 200
	fmt.Printf("a:%v,b:%v\n", a, b)
	a, b = b, a
	fmt.Printf("a:%v,b:%v\n", a, b)

	goldTypes := []string{"gold", "diamond", "secondGold"}

	for _, goldType := range goldTypes {
		fmt.Println(goldType)
	}

	type Color float64
	type Color2 float64

	c := Color(100)
	c2 := Color2(100)
	fmt.Println("color", c)
	fmt.Println(c == Color(c2))
}
