package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input: " HelloW orlD",
			expected: []string{"hellow", "orld"},
		}, {
			input: " H e L l 0 w O rL d",
			expected: []string{"h", "e", "l", "l", "0", "w", "o", "rl", "d"},
		}, {
			input: "",
			expected: []string{},
		}, {
			input: " ",
			expected: []string{},
		}, {
			input: " A ",
			expected: []string{"a"},
		},
	}

	for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("len(actual) = %v != len(expected) = %v - Output sizes don't match", len(actual), len(c.expected))
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]

		if word != expectedWord {
			t.Errorf("word = %v != expectedWord = %v - Output contents don't match", word, expectedWord)
		}
	}
}
}
