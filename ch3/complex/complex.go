// complex
package main

import (
	"fmt"
)

func main() {
	var x complex128 = complex(10, 15)
	y := 1 + 25i
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
