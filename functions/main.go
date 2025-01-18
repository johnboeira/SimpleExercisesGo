package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}

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
			printMenu()
		case "2":
			addItem()
		case "q":
			break loop
		default:
			fmt.Println("Opção não conhecida")
		}
	}
}

func addItem() {
	fmt.Println("Digite o item novo:")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}