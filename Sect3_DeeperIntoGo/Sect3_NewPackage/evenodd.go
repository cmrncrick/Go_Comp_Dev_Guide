package Sect3_NewFunction

import "fmt"

// evenOdd indicates "private" function. EvenOdd indicates "public" function
// that can be used inside of main package
func EvenOdd() {
	// Creating empty slice of integer type
	// var numbers []int

	// Creating slice of integers 0 -10
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Printf("%v is even", value)
			fmt.Println("")
		} else {
			fmt.Printf("%v is odd", value)
			fmt.Println("")
		}
	}
}