package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffe", "Espresso", "Cappucino"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "Chai Tea"

	fmt.Println(arr)
}
