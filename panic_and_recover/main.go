package main

import "fmt"

func main() {
	dividend, divisor := 10, 0
	fmt.Printf("%v dividendo  por %v é igual a %v", dividend, divisor, divide(dividend, divisor))
}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("erro: %v \n", msg)
		}
	}()
	return dividend / divisor
}
