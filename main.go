package main

import (
	"fmt"
	"os"
	"p-gen/generator"
	utils "p-gen/helper"

	flag "github.com/ogier/pflag"
)

var pLength int
var iCaps bool
var iDigits bool
var iSpecialChars bool

func init() {
	flag.IntVar(&pLength, "length", 0, "Password length (between 8 and 20)")
	flag.BoolVarP(&iCaps, "caps", "c", false, "Can include capital letters?")
	flag.BoolVarP(&iSpecialChars, "spec-chars", "s", false, "Can include special characters?")
	flag.BoolVarP(&iDigits, "digits", "d", false, "Can include digits 0 - 9 ?")

}

func main() {
	// parse flags so that the rest of the app can use them
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !utils.IsLengthValid(pLength, 8, 20) {
		fmt.Println("Password length must be between 8 and 20")
		os.Exit(1)
	}

	generator := generator.PasswordGenerator{PLength: pLength, IsCaps: iCaps, IsDigits: iDigits, IsSpecialChars: iSpecialChars}
	password := generator.Generate()
	fmt.Println(password)
}
