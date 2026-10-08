package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"only whitespace":             {input: "  ", expected: []string{}},
		"whitespace on start and end": {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"without space":               {input: "hellothere", expected: []string{"hellothere"}},
		"uppercase":                   {input: "HELLO THERE", expected: []string{"hello", "there"}},
		"numbers":                     {input: "1234 1 2 3 4", expected: []string{"1234", "1", "2", "3", "4"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
