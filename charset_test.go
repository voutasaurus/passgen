package main

// charset_test.go

import (
	"testing"
	"strings"
	"errors"
)

// Extra example sets to test edge cases
var emptyset = charset{}
var negset = charset{'a': false}

// helper functions

// returns the something
func combinations(sets ...charset) []charset {
	combos := []charset{emptyset}

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

func allSets() []charset {
	return combinations(alphabet, digit, special, space, negset)
}

func allSetsNonEmpty() []charset {
	return combinations(alphabet, digit, special, space)[1:]
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
func remove(list []rune, c rune) ([]rune, error) {
	before := len(list)
	list = []rune(strings.Replace(string(list), string(c), "", 1))
	after := len(list)
	if before != after + 1 {
		return list, errors.New("Removing '"+string(c)+"'did not shrink list")
	}
	return list, nil
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
					t.Error("'"+string(element)+"' in union but not either of the arguments")
					t.Log("X:", set1, "Y:", set2, "union(X,Y):", uset)			
				}
			}
			for element, in := range set1 {
				if in && !uset[element] {
					t.Error("'"+string(element)+"' in first argument but not in union")
					t.Log("X:", set1, "Y:", set2, "union(X,Y):", uset)			
				}
			}
			for element, in := range set2 {
				if in && !uset[element] {
					t.Error("'"+string(element)+"' in second argument but not in union")
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
						t.Error("'"+string(element)+"' in union but not the arguments")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)			
					}
				}
				for element, in := range set1 {
					if in && !uset[element] {
						t.Error("'"+string(element)+"' in first argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)			
					}
				}
				for element, in := range set2 {
					if in && !uset[element] {
						t.Error("'"+string(element)+"' in second argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)			
					}
				}
				for element, in := range set3 {
					if in && !uset[element] {
						t.Error("'"+string(element)+"' in third argument but not in union")
						t.Log("X:", set1, "Y:", set2, "Z:", set3, "union(X,Y,Z):", uset)			
					}
				}
			}
		}
	}
}

// testing valid function
// c ∈ valid(F) ⟺ charClass(c) ∈ F
func TestValid(t *testing.T) {
	t.Error("TestValid is not implemented")
}

// testing getCharSubsets function
// s ∈ getCharSubsets(str) ⟺ s ∩ str ≠ ∅
func TestGetCharSubsets(t *testing.T) {
	t.Error("TestGetCharSubsets is not implemented")
}
