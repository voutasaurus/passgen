package main

import (
	"fmt"
	"os"
	"strconv"
	"io"
)

var defaultCharSubsets = map[string]bool{
		"alphabet": true,
		"digit":    true,
		"special":  true,
		"space":    false,
	}

func main() {
	randomPassword(os.Args, os.Stdout)
}

// randomPassword handles the length argument and calls the generate function
// printing the results
func randomPassword(args []string, out io.Writer) {
	// This utility only takes 2 arguments (3 including the name of the program)
	var charSubsets map[string]bool
	switch {
	case len(args) == 0:
		fmt.Fprintln(out, "Called with no executable name or other arguments. This function needs an executable name.")
		return
	case len(args) == 1:
		fmt.Fprintln(out, args[0], "must be called with at least one argument.")
		fmt.Fprintln(out, "That is, a length for the password to be generated.")
		return
	case len(args) == 2:
		charSubsets = defaultCharSubsets
	case len(args) == 3:
		var subsetErr error
		charSubsets, subsetErr = getCharSubsets(args[2])
		if subsetErr != nil {
			fmt.Fprintln(out, args[2], "is not a valid specification for a character set.")
			fmt.Fprintln(out, "A specification contains a character from each set you wish to include.")
			fmt.Fprintln(out, "For example: j!3")
			fmt.Fprintln(out, subsetErr)
			return
		}
	case len(args) > 3:
		fmt.Fprintln(out, args[0], "must be called with at most two arguments.")
		fmt.Fprintln(out, "That is, a length for the password to be generated, and a combination of characters representing the desired character sets")
		return
	}

	// Convert the given argument so it can be used as the length of the password to generate
	length, lengthErr := strconv.Atoi(args[1])
	if lengthErr != nil || length < 0 {
		fmt.Fprintln(out, args[1], "is not a valid password length.")
		return
	}

	// generate the password
	password, err := generate(length, valid(charSubsets))
	if err != nil {
		fmt.Fprintln(out, "Could not generate a random password successfully.")
		fmt.Fprintln(out, err)
		return
	}

	// Print the generated password
	fmt.Fprintln(out, "Random password:", password, "-", len(password), "characters long")
	return
}
