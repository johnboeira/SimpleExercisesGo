package main

import "fmt"

func main() {
	const a = 42

	const (
		d = 2 * 5
		e // = 2 * 5
		f = iota
		g
		h = 10 * iota
	)

	fmt.Println(d, e, f, g, h)
}
