package main

import "fmt"

func main() {
	var a string
	a = "foo"

	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	//Inferred
	var c = true
	fmt.Println(c)

	//short
	d := 3.14
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)

}
