package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	delete(colors, "red")
	//colors := make(map[string]string)

	printMap(colors)
}

func printMap(c map[string]string) {
	for _, hex := range c {
		fmt.Println(hex)
	}
}
