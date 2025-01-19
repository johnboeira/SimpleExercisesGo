package main

import (
	"bufio"
	"demo/package/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Selecione uma opção")
		fmt.Println("1) Menu")
		fmt.Println("2) Add")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Add()
		case "q":
			break loop
		default:
			fmt.Println("Opção não conhecida")
		}
	}
}
