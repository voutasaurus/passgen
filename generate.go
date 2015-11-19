package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// generate creates a random sequence of valid symbols of a certain length
func generate(length int, valid alphabet) (string, error) {
	pass := ""
	errs := []error{}
	for i := 0; i < length; i++ {
		// Attempt to get a random element
		symbol, err := randElem(valid)
		if err != nil {
			// Record error
			errs = append(errs, err)
			// Tolerate up to 50% error rate, and at least 5
			if len(errs) > max(length/2, 5) {
				return "", tooMany(errs)
			}
		} else {
			// If there is no error, add the string to the return string
			pass += symbol
		}
	}
	return pass, nil
}

// randElem gets a random string from a alphabet
func randElem(set alphabet) (string, error) {
	// Create a list to choose a random index from
	list := getList(set)

	if len(list) == 0 {
		return "", errors.New("randElem: alphabet empty. Check alphabet generation")
	}

	// Set the maximum index to choose - casting to big int for crypto/rand
	max := big.NewInt(int64(len(list)))

	// Generate a random index
	// crypto/rand.Int generates a number in [0,max) - See godoc for crypto/rand for more info
	i, err := rand.Int(rand.Reader, max)

	// Error reading from os random source
	if err != nil {
		return "", err
	}

	// No error, return string at random index - casting back to int
	return list[int(i.Int64())], nil
}

// max gets the maximum value of a list of values
func max(is ...int) int {
	m := 0
	for _, i := range is {
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
