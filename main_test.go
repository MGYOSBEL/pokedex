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
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu CHARMANDER Bulbasaur ",
			expected: []string{"pikachu", "charmander", "bulbasaur"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The result length should be %d, but actually it is %d", len(c.expected), len(actual))
			t.Fail()
		}

		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("Expected: %s but found %s", c.expected[i], word)
				t.Fail()
			}
		}

	}
}
