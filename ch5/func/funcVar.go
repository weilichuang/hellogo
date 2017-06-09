package main

import (
	"fmt"
)

func square(n int) int { return n * n }

func main() {
	f := square
	fmt.Println(f(10))

	var s func(int) int
	s = square
	if s != nil {
		fmt.Println(s(100))
	}
}
