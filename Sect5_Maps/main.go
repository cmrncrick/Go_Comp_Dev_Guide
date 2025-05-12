package main

import "fmt"

func main() {

	// Creating map
	// var colors map[string]string

	// Creating map
	// colors := make(map[string]string)
	// Adding value to colors
	// colors["white"] = "#ffffff"

	// Deleting value
	// delete(colors, "white")

	// Declaring a map with keys as strings and values as strings
	colors := map[string]string{
		"red" : "#ff000",
		"green" : "#4bf745",
		"white" : "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	// iterating over the key value pair for the map "c" passed in
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}