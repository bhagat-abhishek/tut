package main

import (
	"fmt"
	"strings"
)

func main() {

	// Using the strings library

	greetings := "Hey! Abby Good Morning."

	// checking if the given string contains the passed sting, the reurm will be a type of bool
	fmt.Println(strings.Contains(greetings, "Hey!!"))

	// replcaing the strings in given strng -- will return a string outpt
	fmt.Println(strings.ReplaceAll(greetings, "Hey!", "Hello,"))

	// convert te given string to uppercase
	fmt.Println(strings.ToUpper(greetings))

	// gives the inder of the passed parameter
	fmt.Println(strings.Index(greetings, "Goo"))

	// returns a slice of array of the strng via a given paramenter
	fmt.Println(strings.Split(greetings, " ")) // here (" ") and empty character will split them into slices and return
}
