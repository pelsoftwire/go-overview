package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(greeting())
}

func greeting() string {
	return quote.Hello()
}
