package main

import (
	"fmt"
	"io"
	"os"
)

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("myfile.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	content, err := io.ReadAll(file)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(content))
// }

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}