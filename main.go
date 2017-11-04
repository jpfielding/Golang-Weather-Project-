package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://tour.golang.org
func main() {
	filename := flag.String("file", "web.html", "The file to writ to")

	flag.Parse()

	resp, err := http.Get("http://example.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read from exmaple.com, write to command line
	f, _ := os.Create(*filename)
	io.Copy(f, resp.Body)
	fmt.Printf("saved file %s\n", *filename)
}

func printStuff() {
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
