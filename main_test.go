package main

import "testing"

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Bob", "Hello, Bob"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}
