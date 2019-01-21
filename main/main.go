package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hamdyjs/link"
)

func main() {
	htmlFile := flag.String("file", "examples/ex1.html", "The html file to parse")
	flag.Parse()

	path, _ := filepath.Abs(*htmlFile)
	fmt.Println("Using path:", path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	links, err := link.Parse(f)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Links: %+v\n", links)
}
