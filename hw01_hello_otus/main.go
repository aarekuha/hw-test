package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func reversePhrase(s string) string {
	return reverse.String(s)
}

func main() {
	fmt.Println(reversePhrase("Hello, OTUS!"))
}
