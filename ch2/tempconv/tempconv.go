// tempconv
package main

import (
	"fmt"
	"math"
	"reflect"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezlingC    Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°c", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°c", f)
}

func main() {
	c := Celsius(150)
	fmt.Println(c)

	v1 := 1e255
	fmt.Println(v1, reflect.TypeOf(v1))
	fmt.Println(math.MaxFloat32)
	f := math.MaxFloat32
	fmt.Println(f == f+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
