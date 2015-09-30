package main

import (
	"fmt"
	"os"
	"strconv"
)

// main handles the length argument and calls the generate function
// printing the results
func main() {
	// This utility only takes one argument (two including the name of the program)
	if len(os.Args) != 2 {
		fmt.Println(os.Args[0], "must be called with one argument.")
		fmt.Println("That is, a length for the password to be generated.")
		return
	}

	// Convert the given argument so it can be used as the length of the password to generate
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(os.Args[1], "is not a valid length.")
		return
	}

	// Specify desired character sets
	charSubsets := map[string]bool{
		"alphabet": true,
		"digit":    true,
		"special":  true,
		"space":    false,
	}

	// generate the password
	password, err := generate(length, valid(charSubsets))
	if err != nil {
		fmt.Println("Could not generate a random password successfully.")
		fmt.Println(err)
		return
	}

	// Print the generated password
	fmt.Println("Random password:", password, "-", len(password), "characters long")
}
