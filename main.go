package main

import (
	"fmt"
	"os"
)

// https://tour.golang.org
func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 10; i++ {
		// to the file
		fmt.Fprintf(f, "Hello, playground %d\n", i)
		// writing to the terminal
		fmt.Fprintf(os.Stdout, "Hello, playground %d\n", i)
		fmt.Fprintf(os.Stderr, "Hello, playground %d\n", i)
	}
}
