package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{"a": 32, "x": 30, "cat": 4}
	fmt.Println(ages)
	fmt.Println(ages["a"])

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	sexs := make(map[string]string)
	sexs["a"] = "boy"
	sexs["b"] = "girl"
	fmt.Println(sexs)
	delete(sexs, "a")
	fmt.Println(sexs)

	sex, ok := sexs["cat"]
	if !ok {
		fmt.Println("can not find sex")
	}
	fmt.Println("sex is" + sex)
}
