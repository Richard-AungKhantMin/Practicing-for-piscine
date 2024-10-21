package main

import "testing"

func TestCountalpha(t *testing.T) {
	testCases := []struct {
		name     string
		a        string
		expected int
	}{
		{"Just a normal test", "Hello world", 10},
		{"With spaces", "H e l l o", 5},
		{"Numbers between them", "H1e2l3l4o", 5},
	}

	for _, each := range testCases {
		t.Run(each.name, func(t *testing.T) {
			result := CountAlpha(each.a)
			if result != each.expected {
				t.Errorf("The result: %d; expected %d", result, each.expected)
			}
		})
	}
}
