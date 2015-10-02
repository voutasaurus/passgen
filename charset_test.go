package main

// charset_test.go

import (
	"testing"
	"strings"
	"errors"
)

// testing getList function
// c ∈ getList(X) ⟺ c ∈ X
func TestGetList(t *testing.T) {

	sets := combinations(alphabet, digit, special, space)

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
// c ∈ union(X,Y) ⟺ c ∈ X ∪ Y
// c ∈ union([X0,..,Xn]...) ⟺ c ∈ ⋃{X_i}
func TestUnion(t *testing.T) {
	t.Error("TestUnion is not implemented")
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
