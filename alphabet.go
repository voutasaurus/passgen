package main

import "errors"

type alphabet map[string]bool

var alphabetic = alphabet{
	"a": true, "A": true,
	"b": true, "B": true,
	"c": true, "C": true,
	"d": true, "D": true,
	"e": true, "E": true,
	"f": true, "F": true,
	"g": true, "G": true,
	"h": true, "H": true,
	"i": true, "I": true,
	"j": true, "J": true,
	"k": true, "K": true,
	"l": true, "L": true,
	"m": true, "M": true,
	"n": true, "N": true,
	"o": true, "O": true,
	"p": true, "P": true,
	"q": true, "Q": true,
	"r": true, "R": true,
	"s": true, "S": true,
	"t": true, "T": true,
	"u": true, "U": true,
	"v": true, "V": true,
	"w": true, "W": true,
	"x": true, "X": true,
	"y": true, "Y": true,
	"z": true, "Z": true,
}

var digit = alphabet{
	"0": true, "1": true, "2": true, "3": true, "4": true,
	"5": true, "6": true, "7": true, "8": true, "9": true,
}

var special = alphabet{
	"~": true, "!": true, "@": true, "#": true, "$": true, "%": true, "^": true, "&": true, "*": true, "(": true, ")": true,
	"-": true, "_": true, "+": true, "=": true, "{": true, "}": true, "[": true, "]": true, "\\": true, "|": true, ":": true,
	";": true, "\"": true, "'": true, ",": true, "<": true, ".": true, ">": true, "?": true, "/": true, "`": true,
}

var space = alphabet{
	" ": true,
}

var symbolsFrom = map[string]alphabet{
	"alphabetic":   alphabetic,
	"digit":        digit,
	"special":      special,
	"space":        space,
	"englishWords": englishWords,
}

// getSymbolSubsets returns the symbolSubsets based on input
// NB: Calling function cannot do spaces at this stage
func getSymbolSubsets(sample string) (map[string]bool, error) {
	if len(sample) > 4 {
		return map[string]bool{}, errors.New("Too many symbols. You only need 4 at most.")
	}
	symbolSubsets := map[string]bool{}
	for _, c := range []rune(sample) {
		if alphabetic[string(c)] {
			symbolSubsets["alphabetic"] = true
		} else if digit[string(c)] {
			symbolSubsets["digit"] = true
		} else if special[string(c)] {
			symbolSubsets["special"] = true
		} else if space[string(c)] {
			symbolSubsets["space"] = true
		} else {
			return map[string]bool{}, errors.New("Invalid symbol used: \"" + string(c) + "\"")
		}
	}
	return symbolSubsets, nil
}

// valid returns a symbol set containing all of the valid symbols
func valid(symbolSubsets map[string]bool) alphabet {
	sets := []alphabet{}
	for subset, include := range symbolSubsets {
		if include {
			sets = append(sets, symbolsFrom[subset])
		}
	}
	return union(sets...)
}

// union takes the union of any number of sets of runes
func union(sets ...alphabet) alphabet {
	unionSet := make(alphabet)
	for _, set := range sets {
		for elem, in := range set {
			// If it"s in this set or already in unionSet
			unionSet[elem] = in || unionSet[elem]
		}
	}
	return unionSet
}

// getList turns a set of runes into a slice of runes
func getList(set alphabet) []string {
	list := make([]string, 0, len(set))
	for elem, in := range set {
		if in {
			list = append(list, elem)
		}
	}
	return list
}

// Pretty printing for debug
func (set alphabet) String() string {
	empty := true
	display := "alphabet{"
	for c, in := range set {
		if in {
			display += "\""
			display += c
			display += "\""
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
