package main

import (
	"fmt"
	"os"
	"strconv"
)

var defaultCharSubsets = map[string]bool{
		"alphabet": true,
		"digit":    true,
		"special":  true,
		"space":    false,
	}

// main handles the length argument and calls the generate function
// printing the results
func main() {
	// This utility only takes 2 arguments (3 including the name of the program)
	var charSubsets map[string]bool
	switch {
	case len(os.Args) == 0:
		fmt.Println("How'd you do that? Heck, I'm not even mad; that's amazing.")
		return
	case len(os.Args) == 1:
		fmt.Println(os.Args[0], "must be called with at least one argument.")
		fmt.Println("That is, a length for the password to be generated.")
		return
	case len(os.Args) == 2:
		charSubsets = defaultCharSubsets
	case len(os.Args) == 3:
		var subsetErr error
		charSubsets, subsetErr = getCharSubsets(os.Args[2])
		if subsetErr != nil {
			fmt.Println(os.Args[2], "is not a valid specification for a character set.")
			fmt.Println("A specification contains a character from each set you wish to include.")
			fmt.Println("For example: j!3")
			fmt.Println(subsetErr)
			return
		}
	case len(os.Args) > 3:
		fmt.Println(os.Args[0], "must be called with at most two arguments.")
		fmt.Println("That is, a length for the password to be generated, and a combination of characters representing the desired character sets")
		return
	}

	// Convert the given argument so it can be used as the length of the password to generate
	length, lengthErr := strconv.Atoi(os.Args[1])
	if lengthErr != nil || length < 0 {
		fmt.Println(os.Args[1], "is not a valid password length.")
		return
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
