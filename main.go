package main

import (
	"fmt"
	"password_generator/generator"
)

var message1 = "Password generator starting..."

func main() {
	fmt.Println(message1)

	// get input from user
	passwordLength, includeCaps, includeDigits, includeSpecialChars := generator.GetUserInput()

	fmt.Println(passwordLength, includeCaps, includeDigits, includeSpecialChars)

	generator := generator.PasswordGenerator{PLength: passwordLength, IsCaps: includeCaps, IsDigits: includeDigits, IsSpecialChars: includeSpecialChars}
	password := generator.Generate()

	fmt.Println(password)
}
