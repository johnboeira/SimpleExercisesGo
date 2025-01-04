package main

import "fmt"

func main() {
	s := "Hello"
	p := &s

	fmt.Println(p)  //Address
	fmt.Println(*p) //value

	*p = "Hello, gopher"

	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p)
}
