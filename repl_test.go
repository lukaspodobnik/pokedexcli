package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello    world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu Charmander BULBASAUR",
			expected: []string{"pikachu", "charmander", "bulbasaur"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "   Pikachu   ",
			expected: []string{"pikachu"},
		},
		{
			input:    "hello\tworld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "foo\nbar",
			expected: []string{"foo", "bar"},
		},
		{
			input:    " \n\t Bulbasaur \t\n ",
			expected: []string{"bulbasaur"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual and length of expected are not the same: len(actual)=%d != %d=len(expected)", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("actual does not match expected: actual[%d]=%s != %s=expected[%d]", i, actual[i], c.expected[i], i)
			}
		}
	}
}
