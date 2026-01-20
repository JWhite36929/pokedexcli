package main 

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		expected []string
	}{
		"trailing spaces":
		{
			input: "  hello  world  ",
			expected: []string{"hello","world"},
		},
		"no split":
		{
			input: "nosplitneeded",
			expected: []string{"nosplitneeded"},
		},
		"capitals":
		{
			input: "John WENT running",
			expected: []string{"john", "went", "running"},
		},
		"everything":
		{
			input: "WORDS aRe Trailing   .    ",
			expected: []string{"words", "are", "trailing", "."},
		},
	}

	for name, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Case %v: actual length does not match expected length", name)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i] 

			if word != expectedWord {
				t.Errorf("actual word, %v, does not match expected word, %v", word, expectedWord)
			}
		}
	}

}