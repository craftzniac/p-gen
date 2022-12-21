package generator

import (
	"math/rand"
	"time"
)

var capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var smallLetters = "abcdefghijklmnopqrstuvwxyz"
var digits = "0123456789"
var specialChars = "!@#$%^&*()_-+={}[]<>.,|"

type PasswordGenerator struct {
	PLength        int
	IsCaps         bool
	IsDigits       bool
	IsSpecialChars bool
}

func (gen PasswordGenerator) Generate() string {
	passwordCharSet := smallLetters
	var password string

	if gen.IsCaps {
		passwordCharSet += capitalLetters
	}

	if gen.IsDigits {
		passwordCharSet += digits
	}

	if gen.IsSpecialChars {
		passwordCharSet += specialChars
	}

	for i := 0; i < gen.PLength; i++ {
		rand.Seed(time.Now().UnixNano())
		charIndex := rand.Intn(len(passwordCharSet))
		password += string(passwordCharSet[charIndex])
	}
	return password
}
