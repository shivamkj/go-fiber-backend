package helper

import (
	"net/url"

	"net/mail"
	"regexp"
)

func Verify(condition bool, errMsg string, errs *[]string) {
	if !condition {
		*errs = append(*errs, errMsg)
	}
}

func IsValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsValidEmail(email string) bool {
	// First, check if the email string is empty
	if email == "" {
		return false
	}

	// Use the net/mail package to validate the email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		// If the email format is invalid, try validating with a regular expression
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
		return re.MatchString(email)
	}

	return true
}
