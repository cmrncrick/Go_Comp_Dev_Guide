package main

import Sect3_NewFunction "cards/Sect3_NewPackage"

// variable types when declaring variable
// bool string int float64
// var test_bool bool
// var test_string string = ""
// var test_int int = 1
// var test_float float64 = 3.14

// func main() {

// 	// Two ways of declaring variables
// 	// Directly declaring it as a string
// 	// var card string = "Ace of Spades"

// 	// Or Go will see that it will be a string and assign it as such itself
// 	// := only used when declaring NEW variables, not when variable is reassigned
// 	// card := "Ace of Spades"

// 	// Creating card and calling newCard function
// 	card := newCard()

// 	// Printing the value of card
// 	fmt.Println(card)

// }

// // by placing string after the () Go knows that this function will return a str
// func newCard() string {
// 	return "Five of Diamonds"
// }

// *************************************************************************************************
// Video 17

// Slices have to have one datatype. No mismatch.

func main () {
	// *********************************************************************************************
	// cards := deck{"Ace of Diamonds", newCard()}

	// Does not modify existing. Returns new item named cards.
	// cards = append(cards, "Six of Spades")
	// *********************************************************************************************
	// // Calling function from deck.go
	// cards := newDeck()

	// // 
	// // cards.print()

	// // Calling the deal function
	// hand, remainingDeck := deal(cards, 5)

	// // Can call the print function because hand and remainingDeck are both deck types
	// hand.print()
	// fmt.Println("*************")
	// remainingDeck.print()
	// *********************************************************************************************
	// // Type conversion from string to Byte
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// *********************************************************************************************
	// Converting deck into string and printing it
	// fmt.Println(cards.toString())
	// *********************************************************************************************
	// // Saving the cards to file
	// cards.saveToFile("my_cards")
	// *********************************************************************************************
	// // Loading the file back into program
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// *********************************************************************************************
	cards := newDeck()

	cards.shuffle()

	cards.print()

	Sect3_NewFunction.EvenOdd()


}




// *************************************************************************************************
// Video 18 -- OO Approach vs GO Approach




