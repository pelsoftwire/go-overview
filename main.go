package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isCreator(e string) bool {
	var creators = [3]string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	for _, a := range creators {
		if a == e {
			return true
		}
	}
	return false
}

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func isPalindrome(e string) bool {
	var strlen = len(e)
	for i := 0; i < strlen/2; i++ {
		if e[i] != e[strlen-i-1] {
			return false
		}
	}
	return true
}

func main() {
	name := getName()
	fmt.Println(respondToName(name))
}

func respondToName(name string) string {
	if len(name) == 0 {
		return "Okay, no greeting for you"
	} else {
		return greeting(name)
	}
}

func greeting(name string) string {
	// length limit 20
	formattedName := name

	if strings.Index(formattedName, " ") != -1 {
		formattedName = formattedName[:strings.Index(formattedName, " ")]
	}
	if len(formattedName) > 20 {
		formattedName = formattedName[:21]
	}

	// note: name is used to decide on greeting addons
	sentence := "Hello, " + formattedName
	if isCreator(name) {
		sentence = sentence + " Thanks for creating me!"
	}
	if len(name) > 20 {
		sentence = sentence + "... Wow, that name's too long for me!"
	}
	if isPalindrome(name) {
		sentence = sentence + " Cool, a palindromic name!"
	}

	return sentence
}
