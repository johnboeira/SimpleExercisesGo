package main

import "fmt"

func main() {
	a := 99

	for {
		fmt.Println(a)
		a += 1
		break
	}

	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Outra forma:")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
