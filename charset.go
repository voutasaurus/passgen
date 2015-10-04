package main

import (
	"fmt"
	"errors"
)

type charset map[rune]bool

var alphabet = charset{
	'a': true, 'A': true,
	'b': true, 'B': true,
	'c': true, 'C': true,
	'd': true, 'D': true,
	'e': true, 'E': true,
	'f': true, 'F': true,
	'g': true, 'G': true,
	'h': true, 'H': true,
	'i': true, 'I': true,
	'j': true, 'J': true,
	'k': true, 'K': true,
	'l': true, 'L': true,
	'm': true, 'M': true,
	'n': true, 'N': true,
	'o': true, 'O': true,
	'p': true, 'P': true,
	'q': true, 'Q': true,
	'r': true, 'R': true,
	's': true, 'S': true,
	't': true, 'T': true,
	'u': true, 'U': true,
	'v': true, 'V': true,
	'w': true, 'W': true,
	'x': true, 'X': true,
	'y': true, 'Y': true,
	'z': true, 'Z': true,
}

var digit = charset{
	'0': true, '1': true, '2': true, '3': true, '4': true,
	'5': true, '6': true, '7': true, '8': true, '9': true,
}

var special = charset{
	'~': true, '!': true, '@': true, '#': true, '$': true, '%': true, '^': true, '&': true, '*': true, '(': true, ')': true,
	'-': true, '_': true, '+': true, '=': true, '{': true, '}': true, '[': true, ']': true, '\\': true, '|': true, ':': true,
	';': true, '\'': true, '"': true, ',': true, '<': true, '.': true, '>': true, '?': true, '/': true, '`': true,
}

var space = charset{
	' ': true,
}

var charsFrom = map[string]charset{
	"alphabet": alphabet,
	"digit":    digit,
	"special":  special,
	"space":    space,
}

// getCharSubsets returns the charSubsets based on input
// Cannot do spaces at this stage
func getCharSubsets(sample string) (map[string]bool, error) {
	if len(sample) > 3 {
		return map[string]bool{}, errors.New("Too many characters. You only need 3 at most.")
	}
	charSubsets := map[string]bool{}
	for _, c := range []rune(sample) {
		if alphabet[c] {
			charSubsets["alphabet"] = true
		} else if digit[c] {
			charSubsets["digit"] = true
		} else if special[c] {
			charSubsets["special"] = true
		} else {
			return map[string]bool{}, errors.New("Invalid character used: '" + string(c) + "'")
		}
	}
	return charSubsets, nil
}

// valid returns a character set containing all of the valid characters
func valid(charSubsets map[string]bool) charset {
	sets := []charset{}
	for subset, include := range charSubsets {
		if include {
			sets = append(sets, charsFrom[subset])
		}
	}
	return union(sets...)
}

// union takes the union of any number of sets of runes
func union(sets ...charset) charset {
	unionSet := make(charset)
	for _, set := range sets {
		for elem, in := range set {
			// If it's in this set or already in unionSet
			unionSet[elem] = in || unionSet[elem]
		}
	}
	return unionSet
}

// getList turns a set of runes into a slice of runes
func getList(set charset) []rune {
	list := make([]rune, 0, len(set))
	for elem, in := range set {
		if in {
			list = append(list, elem)
		}
	}
	return list
}

// Pretty printing for debug
func (set charset) String() string {
	empty := true
	display := "charset{"
	for c, in := range set {
		if in {
			display += "'"
			display += fmt.Sprintf("%c", c)
			display += "'"
			display += ", "
			empty = false
		}
	}
	if !empty {
		display = display[:len(display)-2]		
	}
	display += "}"
	return display
}
