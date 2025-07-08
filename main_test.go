package main

import "testing"

// TODO: refactor struct with typedef
var greetingTests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Bob", "Hello, Bob"},
	{"Robert Griesemer", "Hello, Robert Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken Thanks for creating me!"},
	{"Pel 123456789123456789", "Hello, Pel... Wow, that name's too long for me!"},
	{"Pel B", "Hello, Pel"},
	{"aba", "Hello, aba Cool, a palindromic name!"},
	{"aba p", "Hello, aba"},
	{"pppppppppppppppppppp pppppppppppppppppppp", "Hello, pppppppppppppppppppp... Wow, that name's too long for me! Cool, a palindromic name!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

var respondToNameTests = []struct {
	name     string
	expected string
}{
	{"", "Ok, no greeting for you"},
	{"Pel B", "Hello, Pel B"},
}

func testRespondToName(t *testing.T) {
	for _, test := range respondToNameTests {
		result := respondToName(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}
