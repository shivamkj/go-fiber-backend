package helper

import (
	"fmt"
)

// Below codes are limited to printable charcaters only
// For ref. total printable character in ASCII - 95
const (
	asciiStartCode = 32
	asciiEndCode   = 126
	idLenLimit2    = 2
	idStartForLen2 = "  "
)

// GenerateNextId creates the next id of 2 character length based on current ID within ASCII printable characters.
func GenerateNextIdLen2(currentId string) string {
	runes := []rune(currentId)
	nextId := incrementRune(runes, idLenLimit2)
	if nextId[0] > 120 {
		Logger.Warn("ids reaching unsafe limit", "left range (approx.)", (126-nextId[0])*95)
	}
	if nextId[0] == asciiStartCode && nextId[1] == asciiStartCode {
		fmt.Println("No IDs available within the range")
		return currentId
	}
	return string(runes)
}

// incrementRune increments the given rune to the next combination within the
// limited set of printable ASCII characters code point
func incrementRune(runes []rune, length int) []rune {
	for i := length - 1; i >= 0; i-- {
		if runes[i] < asciiEndCode {
			runes[i]++
			break
		} else {
			runes[i] = asciiStartCode
		}
	}
	return runes
}
