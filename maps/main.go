package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {
			"Espresso",
			"Cappccino",
		},
		"tea": {
			"Hot tea",
			"Chai Tea",
		},
	}

	fmt.Println(m["coffee"])

	fmt.Println(m)

	m["others"] = []string{"Hot chocolate"}

	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	v, ok := m["tea"]

	fmt.Println(v, ok)
}
