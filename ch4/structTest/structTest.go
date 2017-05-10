package main

import (
	"fmt"
)

func main() {
	user := User{ID: 1, Name: "andy", Age: 18}
	fmt.Printf("%#v\n", user)
	fmt.Printf("%v\n", user)
	age := &user.Age
	*age += 5
	fmt.Println(user.Age)
}

type User struct {
	ID   int
	Name string
	Age  int8
}
