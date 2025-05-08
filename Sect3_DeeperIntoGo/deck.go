package main


import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// New function that will return a type of deck
func newDeck() deck {
	
	cards := deck{}

	// Creating two lists
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Looping over two strings and joining values together
	// The '_' is saying that we know there is a variable there 
	// but we are not going to use it.
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	// Return the slice
	return cards
}

// Receiver -- it's type is deck and it's name is print
// d represents a variable name of type 'deck'
func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)

	}
}



// New function that will deal a hand of cards
// (deck, deck) tells Go we are going to return 2 values,
// with both being of the type "deck"
func deal(d deck, handSize int) (deck, deck) {
	
	// Return everything in the deck up to the handsize as first value
	// Return everything from handsize to end of deck as second value
	return d[:handSize], d[handSize:]
}

// takes deck type then converts to string
func (d deck) toString() string {
	// Converting d to list of strings
	// s := []string(d)

	// Join list of strings together in one string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	// If there is an error then do this
	if err != nil {
		// Print the error
		fmt.Println("Error:", err)
		
		// Log the error

		
		// Quit the program
		os.Exit(1)
	}

	// Splitting apart the bit slice that's returned from Read File 
	s := strings.Split(string(bs), ",")

	// Creating a new deck type object with the previous bit slice
	return deck(s)
	
}

func (d deck) shuffle() {
	// Take input deck and shuffle cards

	// Creating a source object (seed value)
	// Using Now and UnixNano to have new int64 on every run of code
	source := rand.NewSource(time.Now().UnixNano())

	// Create New rand object
	r := rand.New((source))

	// Counting how many index items are in d
	for i := range d{
		newPosition := r.Intn(len(d) - 1)

		// Swapping postiions
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}





