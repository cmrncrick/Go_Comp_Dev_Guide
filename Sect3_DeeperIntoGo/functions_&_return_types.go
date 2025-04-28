package main

import "fmt"

// variable types when declaring variable
// bool string int float64
var test_bool bool
var test_string string = ""
var test_int int = 1
var test_float float64 = 3.14

func main() {

	// Two ways of declaring variables
	// Directly declaring it as a string
	// var card string = "Ace of Spades"
	
	// Or Go will see that it will be a string and assign it as such itself
	// := only used when declaring NEW variables, not when variable is reassigned
	// card := "Ace of Spades"

	card := newCard()

	fmt.Println(card)

}


// by placing string after the () Go knows that this function will return a str
func newCard() string {
	return "Five of Diamonds"
}

