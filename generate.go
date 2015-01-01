package main

import (
	"fmt"
	"errors"
	"crypto/rand"
	"math/big"
)

// generate creates a random sequence of valid characters of a certain length
func generate(length int, valid charset) (string, error) {
	pass := ""
	errs := []error{}
	for i := 0; i < length; i++ {
		// Attempt to get a random element
		char, err := randElem(valid)
		if err != nil {
			// Record error
			errs := append(errs, err)
			// Tolerate up to 50% error rate, and at least 5
			if len(errs) > max(length / 2, 5) {
				return "", tooMany(errs)
			}
		} else {
			// If there is no error, add the rune to the return string
			pass += fmt.Sprintf("%c", char)
		}
	}
	return pass, nil
}

// randElem gets a random rune from a charset
func randElem(set charset) (rune, error) {
	// Create a list to choose a random index from
	list := getlist(set)

	// Set the maximum index to choose - casting to big int for crypto/rand
	max := big.NewInt(int64(len(list)))

	// Generate a random index (See godoc for crypto/rand for info)
	i, err := rand.Int(rand.Reader, max)

	// Error reading from os random source
	if err != nil {
		return ' ', err
	}

	// No error, return rune at random index - casting back to int
	return list[int(i.Int64())], nil
}

// max gets the maximum value of a list of values
func max(is ...int) int {
	m := 0
	for i := range is {
		if i > m {
			m = i
		}
	}
	return m
}

// tooMany returns an amalgamated error message listing all of the
// errors that caused the calling function to give up
func tooMany(errs []error) error {
	message := "Too many errors: \n"
	for i, e := range errs {
		message += fmt.Sprintln("\t", i, "-", e)
	}
	return errors.New(message)
}
