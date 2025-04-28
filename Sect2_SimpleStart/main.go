// package main is used when creating your project as executable
// Can use package blahblah to create "reusable" code, i.e. library
package main

// All native libraries at golang.org/pkg
import (
	"fmt"
) // You don't have to have a comma after your package. example below
// import (
// 		"fmt"
// 		"flag")

// func == function (same as python def)
func main() {
	fmt.Println("Hi there!")
}


// How do you run the code?
// 		go run main.go (in the command prompt / main.go should be replaced with filename)