package main

import "fmt"

func main() {

	// Declaring maps:
	// var colors map[string]string
	// colors := make(map[int]string)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//Adding stuff to map:
	// colors["blue"] = "#someheximtoolazytogoole" 


	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c{
		fmt.Printf("Color %v has hex value of %v \n", color, hex)
	}
}