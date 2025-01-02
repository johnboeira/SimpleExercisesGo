package main

import "fmt"

func main() {
	a, b := 5, 2

	fmt.Println(a + b)

	fmt.Println(a % b) // a%2 ==0

	fmt.Println(float32(a) / float32(b))

	//Comparation operators
	fmt.Println(a == b)
	fmt.Println(a < b)

}
