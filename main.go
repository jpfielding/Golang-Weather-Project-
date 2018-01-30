package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// https://tour.golang.org
/*func main() {
	filename := flag.String("file", "web.html", "The file to writ to")

	flag.Parse()

	resp, err := http.Get("http://samples.openweathermap.org")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read from example.com, write to command line
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
} */

func main() {
	city := flag.String("city", "Pittsburgh", "Default city")

	flag.Parse()

	appID := "b6907d289e10d714a6e88b30761fae22"
	url := fmt.Sprintf("http://samples.openweathermap.org/data/2.5/forecast?q=%s&mode=xml&appid=%s", *city, appID)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	if strings.Contains(b.String(), "snow") {
		fmt.Println("Snow bitches!")
	}
}
