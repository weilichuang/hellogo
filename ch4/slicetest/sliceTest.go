package main

import (
	"fmt"
)

func main() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	s := a[:]
	fmt.Printf("before s:%v\n", s)
	fmt.Printf("before a:%v\n", a)
	reverseSlice(s)
	fmt.Printf("after s:%v\n", s)
	fmt.Printf("after a:%v\n", a)
	reverseSlice(s[3:])
	fmt.Printf("after a:%v\n", a)
	fmt.Printf("Slice cap:%v\n", cap(s))
	s = append(s, 8)
	fmt.Printf("after a:%v\n", a)
	fmt.Printf("after s:%v\n", s)
	fmt.Printf("Slice length:%v\n", len(s))
	fmt.Printf("Slice cap:%v\n", cap(s))

	var runes []rune
	for _, r := range "Hello 堵塞饿" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//反转Slice
func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
