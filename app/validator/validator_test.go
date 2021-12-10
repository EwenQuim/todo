package validator

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"truc ceci ", "truc ceci"},
		{"truc: ceci", "truc: ceci"},
		{"truc : ceci", "truc: ceci"},
		{"  truc  : ceci", "truc: ceci"},
		{"  tRUc  : ceci", "truc: ceci"},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", i), func(t *testing.T) {
			t.Parallel()

			obtained := CleanItem(tc.input)
			if obtained != tc.expected {
				t.Errorf("Error in test %v: got %q, expected %q", i, obtained, tc.expected)
			}
		})
	}
}
