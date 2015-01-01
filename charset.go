package main

import "fmt"

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
	';': true, '\'': true, '"': true, ',': true, '<': true, '.': true, '>': true, '?': true, '/': true, '`': true, ' ': true,
}

// valid returns a character set containing all of the valid characters
func valid() charset {
	return union(alphabet, digit, special)
}

// union takes the union of any number of sets of runes
func union(sets ...charset) charset {
	ret := make(charset)
	for _, set := range sets {
		for elem := range set {
			ret[elem] = true
		}
	}
	return ret
}

// getList turns a set of runes into a slice of runes
func getlist(set charset) []rune {
	list := make([]rune, 0, len(set))
	for elem := range set {
		list = append(list, elem)
	}
	return list
}

// Pretty printing for debug
func (set charset) String() string {
	ret := "charset{"
	for c := range set {
		ret += "'"
		ret += fmt.Sprintf("%c", c)
		ret += "'"
		ret += ", "
	}
	ret = ret[:len(ret)-2]
	ret += "}"
	return ret
}
