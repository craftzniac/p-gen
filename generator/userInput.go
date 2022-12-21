package generator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func GetUserInput() (int, bool, bool, bool) {
	var length int
	var includeCaps bool
	var includeDigits bool
	var includeSymbols bool

	// returns a valid length
	length = getLength()
	includeCaps = getIncludeCaps()
	includeDigits = getIncludeDigits()
	includeSymbols = getIncludeSymbols()

	return length, includeCaps, includeDigits, includeSymbols
}

func getLength() int {
	var minLength = 8
	var maxLength = 20
	// prompt user for a length
	for {
		fmt.Println("provide password length (8 - 20)")

		if scanner.Scan() {

			text := scanner.Text()
			// tries to cast to an integer
			length, err := strconv.Atoi(text)
			// if there's no error, then check to see if it's a valid password length
			if err == nil {
				if length >= minLength && length <= maxLength {
					return length
				} else {
					fmt.Println("the length provided is invalid, try again")
				}
			} else {
				fmt.Println("password length must be a number, try again")
			}

		} else {
			panic("Could not read from stdin")
		}
	}
}

func getIncludeCaps() bool {
	return getBoolValue("Can password include capital letters? y/n")
}

func getIncludeDigits() bool {
	return getBoolValue("Can password include numbers? y/n ")
}

func getIncludeSymbols() bool {
	return getBoolValue("Can password include special characters like '?', '<', '%', etc ? y/n ")
}

// a helper function to get user input for a y/n question
func getBoolValue(prompt string) bool {
	for {
		fmt.Println(prompt)
		if scanner.Scan() {
			text := scanner.Text()
			text = strings.ToLower(strings.Trim(text, " "))
			if text == "y" {
				return true
			} else if text == "n" {
				return false
			} else {
				fmt.Println("Input is not valid, try again")
			}
		} else {
			panic("Couldn't read from stdin")
		}
	}
}
