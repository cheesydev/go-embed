package main

import (
	_ "embed"
	"fmt"

	"github.com/cheesydev/go-embed/storage"
)

//go:embed hello.txt
var s string

func main() {
	// print content of hello.txt
	fmt.Println("hello.txt: ", s)

	_, err := storage.Find("unknown.txt")
	if err != nil {
		fmt.Println(err)
	}

	content, err := storage.Find("alice.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("alice found: ", content)
}
