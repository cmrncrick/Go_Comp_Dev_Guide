// Video 62

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}


func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println((resp))

	// // Empty byte slice with empty space for 99999 elements
	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
	
	lw := logWriter{}

	io.Copy(lw, resp.Body)
	// Does the same thing as the bs code
	// io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}