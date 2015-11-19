package main

// alphabet_test.go

import (
	"errors"
	"testing"
)

// Extra example sets to test edge cases
var emptyset = alphabet{}
var negset = alphabet{"a": false}

// helper functions

// Combinations returns a list of the alphabets that are combinations of the alphabets passed as arguments
func combinations(sets ...alphabet) []alphabet {
	combos := []alphabet{emptyset}

	// Builds up by doubling {∅} → {∅, X} -> {∅, X, Y, X∪Y} → ...
	for _, set := range sets {
		// Note that the length of combos is evaluated before the loop
		// So the new values are not iterated over
		for _, combo := range combos {
			combos = append(combos, union(set, combo))
		}
	}

	return combos
}

func allSets() []alphabet {
	return combinations(alphabetic, digit, special, space, negset)
}

func allSetsNonEmpty() []alphabet {
	return combinations(alphabetic, digit, special, space)[1:]
}

// testing getList function
// c ∈ getList(X) ⟺ c ∈ X
func TestGetList(t *testing.T) {

	sets := allSets()

	for _, set := range sets {
		list := getList(set)
		// c ∈ getList(X) ⟹ c ∈ X
		for _, x := range list {
			if !set[x] {
				t.Error(x, "found in", list, "but not in", set)
			}
		}
		// c ∈ X ⟹ c ∈ getList(X)
		for x, in := range set {
			if in {
				var err error
				list, err = remove(list, x)
				if err != nil {
					t.Error(x, "found in", set, "but not in", list, err)
				}
			}
		}
		// ∀c ∈ X . ∃! c ∈ getList(x)
		if len(list) != 0 {
			t.Error("Elements of", set, "repeated by getList. Expected empty list. Got", list)
		}
	}

}

// remove removes a sinlge element matching c from the list
func remove(list []string, s string) ([]string, error) {
	for i, x := range list {
		if x == s {
			if i+1 == len(list) {
				return list[:len(list)-1], nil
			} else {
				return append(list[:i], list[i+1:]...), nil
			}
		}
	}
	return list, errors.New("Removing '" + s + "'did not shrink list")
}

// testing union function
// union() = ∅
func TestUnionNone(t *testing.T) {
	testEmpty := union()
	if len(testEmpty) != 0 {
		t.Error("union() should be empty, instead it is", testEmpty)
	}
}

// union(X) = X
func TestUnionOne(t *testing.T) {

	sets := allSets()

	for _, set := range sets {
		testSingle := union(set)
		if len(testSingle) != len(set) {
			t.Error("union(X) should be X, instead union(", set, ") is", testSingle)
		}
	}
}

// c ∈ union(X,Y) ⟺ c ∈ X ∪ Y
func TestUnionTwo(t *testing.T) {

	sets := allSets()

	for _, set1 := range sets {
		for _, set2 := range sets {
			uset := union(set1, set2)

			for element, in := range uset {
				if in && !set1[element] && !set2[element] {
					t.Error("'" + string(element) + "' in union but not either of the arguments")
					t.Log("X:", set1, "Y:", set2, "union(X,Y):", uset)
				}
			}
			for element, in := range set1 {
				if in && !uset[element] {
					t.Error("'" + string(element) + "' in first argument but not in union")
					t.Log("X:", set1, "Y:", set2, "union(X,Y):", uset)
				}
			}
			for element, in := range set2 {
				if in && !uset[element] {
					t.Error("'" + string(element) + "' in second argument but not in union")
					t.Log("X:", set1, "Y:", set2, "union(X,Y):", uset)
				}
			}
		}
	}

}

// c ∈ union([X0,..,Xn]...) ⟺ c ∈ ⋃{X_i}
func TestUnionMulti(t *testing.T) {
	sets := allSets()

	for _, set1 := range sets {
		for _, set2 := range sets {
			for _, set3 := range sets {
				uset := union(set1, set2, set3)

				for element, in := range uset {
					if in && !set1[element] && !set2[element] && !set3[element] {
						t.Error("'" + string(element) + "' in union but not the arguments")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)
					}
				}
				for element, in := range set1 {
					if in && !uset[element] {
						t.Error("'" + string(element) + "' in first argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)
					}
				}
				for element, in := range set2 {
					if in && !uset[element] {
						t.Error("'" + string(element) + "' in second argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)
					}
				}
				for element, in := range set3 {
					if in && !uset[element] {
						t.Error("'" + string(element) + "' in third argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)
					}
				}
			}
		}
	}
}

// testing valid function
// c ∈ valid(symbolCombos) ⟺ symbolClass(c) ∈ symbolCombos
func TestValid(t *testing.T) {
	allSymbolCombos := allSymbolClasses()

	for _, symbolCombos := range allSymbolCombos {
		symbols := valid(symbolCombos)
		// c ∈ valid(symbolCombos) ⟹ symbolClass(c) ∈ symbolCombos
		for c, in := range symbols {
			if in {
				if !symbolCombos[symbolClass(c)] {
					t.Error("'"+string(c)+"' in valid(", symbolCombos, ") =", symbols, "but '"+string(c)+"' not in", symbolCombos)
				}
			}
		}
		// symbolClass(c) ∈ symbolCombos ⟹ c ∈ valid(symbolCombos)
		for class, in := range symbolCombos {
			if in {
				for c, in := range symbolsFrom[class] {
					if in {
						if !symbols[c] {
							t.Error("'"+string(c)+"' not in valid(", symbolCombos, ") =", symbols, "but '"+string(c)+"' in", symbolCombos)
						}
					}
				}
			}
		}
	}
}

func allSymbolClasses() []map[string]bool {

	symbolclasses := []map[string]bool{}

	for i := 0; i < 16; i++ {
		a := map[string]bool{
			"alphabetic": i >= 8,
			"digit":      i%8 >= 4,
			"special":    i%4 >= 2,
			"space":      i%2 == 1,
		}
		symbolclasses = append(symbolclasses, a)
	}

	return symbolclasses
}

func symbolClass(c string) string {
	switch {
	case alphabetic[c]:
		return "alphabetic"
	case digit[c]:
		return "digit"
	case special[c]:
		return "special"
	case space[c]:
		return "space"
	}
	return ""
}

// testing getSymbolSubsets function
// s ∈ getSymbolSubsets(str) ⟺ s ∩ str ≠ ∅
func TestGetSymbolSubsets(t *testing.T) {
	testStrings := allTestStrings()

	for _, testString := range testStrings {
		subsets, err := getSymbolSubsets(testString)
		if err != nil {
			t.Error("getSymbolSubsets("+testString+") produced error:", err)
		}
		// ∀subset ∈ subsets ∃symbol ∈ subset ∩ testString
		for subset, in := range subsets {
			if in {
				found := false
				for _, symbol := range []rune(testString) {
					if symbolsFrom[subset][string(symbol)] {
						found = true
					}
				}
				if !found {
					t.Error("Couldn't find any symbols from", testString, "in", subset, "\ngetSymbolSubsets("+testString+") =", subsets)
				}
			}
		}
		// ∀symbol ∈ testString ∃subset ∈ subsets . symbol ∈ subset
		for _, symbol := range []rune(testString) {
			found := false
			for subset, in := range subsets {
				if in && symbolsFrom[subset][string(symbol)] {
					found = true
				}
			}
			if !found {
				t.Error("Couldn't find any subsets from", subsets, "containing", symbol, "\ngetSymbolSubsets("+testString+") =", subsets)
			}
		}
	}

}

func allTestStrings() []string {
	testStrings := []string{}
	allSymbols := getList(union(alphabetic, digit, special, space))
	for _, symbol := range allSymbols {
		testStrings = append(testStrings, string(symbol))
	}
	for _, symbol1 := range allSymbols {
		for _, symbol2 := range allSymbols {
			testStrings = append(testStrings, string(symbol1)+string(symbol2))
		}
	}
	for _, symbol1 := range allSymbols {
		for _, symbol2 := range allSymbols {
			for _, symbol3 := range allSymbols {
				testStrings = append(testStrings, string(symbol1)+string(symbol2)+string(symbol3))
			}
		}
	}

	testStrings = append(testStrings, "aaaa")
	testStrings = append(testStrings, "1111")
	testStrings = append(testStrings, "####")
	testStrings = append(testStrings, "    ")
	testStrings = append(testStrings, "a$1 ")

	return testStrings
}
