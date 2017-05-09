// intptr
package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	fmt.Println(x)
	p1 := &x
	fmt.Println(p == p1)

	v := 5
	fmt.Println(v)
	incr(&v)
	fmt.Println(v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++ //只是增加p指向的变量的值，并未改变p指针
	return *p
}
