package helper

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
	testCases := map[string]bool{
		"https://www.example.com":        true,
		"http://localhost:8080":          true,
		"ftp://ftp.example.com/file.txt": true,
		"invalidurl":                     false,
		"notavalidurl.com":               false,
	}

	for input, expected := range testCases {
		Check(IsValidURL(input) == expected, t)
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"Empty email", "", false},
		{"Simple valid email", "john.doe@example.com", true},
		{"Email with digits", "jane.doe99@example.org", true},
		{"Email with underscore", "jane_doe@example.net", true},
		{"Email with plus sign", "jane+doe@example.co.uk", true},
		{"Email with top-level domain with more than 2 letters", "john.doe@example.museum", true},
		{"Invalid email (missing @)", "johndoeexample.com", false},
		{"Invalid email (missing domain)", "john.doe@", false},
	}

	for _, test := range tests {
		Check(IsValidEmail(test.email) == test.expected, t)
	}
}
