package main

import "fmt"

func main() {

	// Two ways of declaring variables
	// var card string = "Ace of Spades"
	
	// := only used when declaring NEW variables, not when variable is reassigned
	// card := "Ace of Spades"

	card := newCard()

	fmt.Println(card)

}


// by placing string after the () Go knows that this function will return a str
func newCard() string {
	return "Five of Diamonds"
}

// You can initialize a variable outside of a function, you just can't assign
// a value to it using := syntax.

// Valid:
// package main
 
// import "fmt"
 
// var deckSize int
 
// func main() {
//   deckSize = 50
//   fmt.Println(deckSize)
// }

// NOT Valid:
// package main
 
// import "fmt"
 
// var deckSize int = 52
 
// func main() {

//   fmt.Println(deckSize)
// }