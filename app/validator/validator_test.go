package validator

import (
	"fmt"
	"testing"
)

func TestCleanItem(t *testing.T) {

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

func TestGetGroupAndContent(t *testing.T) {

	testCases := []struct {
		input           string
		expectedGroup   string
		expectedContent string
	}{
		{"truc ceci ", "", "truc ceci"},
		{"truc: ceci", "truc", "ceci"},
		{"truc : ceci", "truc", "ceci"},
		{"  truc  : ceci", "truc", "ceci"},
		{"  tRUc  : ceci", "truc", "ceci"},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", i), func(t *testing.T) {
			t.Parallel()

			obtainedGroup, obtainedContent := GetGroupAndContent(tc.input)
			if obtainedGroup != tc.expectedGroup {
				t.Errorf("Error in test %v: group error: got %q, expected %q", i, obtainedGroup, tc.expectedGroup)
			}
			if obtainedContent != tc.expectedContent {
				t.Errorf("Error in test %v: content error: got %q, expected %q", i, obtainedContent, tc.expectedContent)
			}
		})
	}
}
