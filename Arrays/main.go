package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffe", "Espresso", "Cappucino"}
	fmt.Println(s)

	fmt.Println(s[1])

	s[1] = "Chai Tea"

	fmt.Println(s)

	s = append(s, "Hot Chocolate", "Hot Tead") //Add

	fmt.Println(s)

	slices.Delete(s, 1, 2) //retira apenas o 1

	fmt.Println(s)
}
