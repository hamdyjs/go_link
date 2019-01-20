package main

import (
	"fmt"
	"os"
	"path/filepath"

	link "github.com/hamdyjs/go_link"
)

func main() {
	path, _ := filepath.Abs("examples/ex1.html")
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
