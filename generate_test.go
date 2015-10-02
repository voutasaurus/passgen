package main

import (
	"testing"
	"strings"
	"errors"
)

// testing generate function
// len(generate(n, X)) = n
func TestGenerateLength(t *testing.T) {
	for i := 0; i < 100; i++ {
		s, err := generate(i, valid(defaultCharSubsets))
		if err != nil {
			t.Error("Error generating at length", i, err)
		} else if len(s) != i {
			t.Error("len(generate("+string(i)+"), X)) =", len(s))
		}
	}
}

// generate(n, X) ⊆ X
func TestGenerateContents(t *testing.T) {
	sets := combinations(alphabet, digit, special, space)

	for _, set := range sets {
		for i := 0; i < 100; i++ {
			s, err := generate(20, set)
			if err != nil {
				t.Error("Error generating.", err)
			} else {
				for _, c := range s {
					if !set[c] {
						t.Error(c, "generated as part of", s, "only expected elements from", set)
					}
				}
			}
		}
	}

}

// TO FIX: This is a temporary hack. 
// There is a smarter more general way to do this
func combinations(sets ...charset) []charset {
	combos := []charset{}
	combos = []charset{alphabet, digit, special, space, union(alphabet,digit), union(alphabet,special), union(alphabet,space), union(digit,special), union(digit,space), union(special,space), union(alphabet,digit, special), union(alphabet,digit, space), union(alphabet,special, space), union(digit,special,space), union(alphabet,digit,special,space)}
	return combos
}

// generate(n, X) has no discernable pattern between subsequent calls
func TestGenerateRandomness(t *testing.T) {
	t.Error("TestGenerateRandomness is not implemented")
}

// testing randElem function
// randElem(∅) = ' ', CharsetEmptyError
func TestRandElemEmptySet(t *testing.T) {
	emptyCharset := charset{}
	_, err := randElem(emptyCharset)
	if err == nil {
		t.Error("Expected emptyCharset error, but error was nil")
	}
}

// randElem(X) ∈ X
func TestRandElemMembership(t *testing.T) {
	var char rune
	var err error

	// alphabet
	for i := 0; i < 1000; i++ {
		char, err = randElem(alphabet)
		if err != nil {
			t.Error("Generating random alphabetic character failed", err.Error())
		} else if !alphabet[char] {
			t.Error("Generated '"+string(char)+"' but this is not in alphabet:", alphabet)
		}
	}

	// digit
	for i := 0; i < 100; i++ {
		char, err = randElem(digit)
		if err != nil {
			t.Error("Generating random digit failed", err.Error())
		} else if !digit[char] {
			t.Error("Generated '"+string(char)+"' but this is not in digit:", digit)
		}
	}

	// special
	for i := 0; i < 1000; i++ {
		char, err = randElem(special)
		if err != nil {
			t.Error("Generating random special character failed", err.Error())
		} else if !special[char] {
			t.Error("Generated '"+string(char)+"' but this is not in special:", special)
		}
	}

	// spaces
	for i := 0; i < 10; i++ {
		char, err = randElem(space)
		if err != nil {
			t.Error("Generating random space character failed", err.Error())
		} else if !space[char] {
			t.Error("Generated '"+string(char)+"' but this is not in space:", space)
		}
	}
}

// randElem(X) has no discernable pattern between subsequent calls
func TestRandElemRandomness(t *testing.T) {
	t.Error("TestRandElemRandomness is not implemented")
}

// testing max function
// max() = 0
func TestMaxOfNone(t *testing.T) {
	m := max()
	if m != 0 {
		t.Error("max() should be 0. Instead it is", m)
	}
}

// max(a) = a
func TestMaxOfOne(t *testing.T) {
	for i := 0; i < 10; i++ {
		m := max(i)
		if m != i {
			t.Error("max("+string(i)+") should be", i, ". Instead it is", m)			
		}
	}
}

// max(a,b) = a ⟺ a > b
func TestMaxOfTwo(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m := max(i,j)
			if i >= j && m != i {
				t.Error("max("+string(i)+", "+string(j)+") should be", i, ". Instead it is", m)			
			}
			if j >= i && m != j {
				t.Error("max("+string(i)+", "+string(j)+") should be", j, ". Instead it is", m)			
			}
		}
	}
}

// max(xs...) = x_i ⟺ ∀j . x_i > x_j
func TestMaxOfN(t *testing.T) {
	if 2 != max(0,1,2) { t.Error("max(0,1,2) should be 2. Instead is", max(0,1,2)) }
	if 2 != max(0,2,1) { t.Error("max(0,2,1) should be 2. Instead is", max(0,2,1)) }
	if 2 != max(1,0,2) { t.Error("max(1,0,2) should be 2. Instead is", max(1,0,2)) }
	if 2 != max(1,2,0) { t.Error("max(1,2,0) should be 2. Instead is", max(1,2,0)) }
	if 2 != max(2,0,1) { t.Error("max(2,0,1) should be 2. Instead is", max(2,0,1)) }
	if 2 != max(2,1,0) { t.Error("max(2,1,0) should be 2. Instead is", max(2,1,0)) }
}

// testing tooMany function
// ∀i . errs[i].String() substring of tooMany(errs...) 
func TestTooManySubstring(t *testing.T) {
	errs := []error{}
	for i := 0; i < 10; i++ {
		errs = append(errs, errors.New(strings.Repeat(string(i), 10)))
		for j := 0; j <= i; j++ {
			found := strings.Contains(tooMany(errs).Error(), errs[j].Error())
			if !found {
				t.Error(tooMany(errs).Error(), "does not contain", errs[j])
			}
		}
	}
}

// #lines in tooMany(errs...) = len(errs) + 1
func TestTooManyLineCount(t *testing.T) {
	errs := []error{}
	for i := 0; i < 10; i++ {
		errs = append(errs, errors.New(string(i)))
		lineCount := strings.Count(tooMany(errs).Error(), "\n")
		lineMismatch := len(errs) + 1 != lineCount
		if lineMismatch { t.Error("len(errs):", len(errs), "lines:", lineCount, "") }
	}
}
