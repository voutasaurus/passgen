package main

import (
	"testing"
)

// testing generate function
// len(generate(n, X)) = n
// generate(n, X) ⊆ X
// generate(n, X) has no discernable pattern between subsequent calls
func TestGenerate(t *testing.T) {
	t.Error("TestGenerate is not implemented")
	return
}

// testing randElem function
// randElem(∅) = ' ', CharsetEmptyError
// randElem(X) ∈ X
// randElem(X) has no discernable pattern between subsequent calls
func TestRandElem(t *testing.T) {
	t.Error("TestRandElem is not implemented")
	return
}

// testing max function
// max() = 0
// max(a) = a
// max(a,b) = a ⟺ a > b
// max(xs...) = x_i ⟺ ∀j . x_i > x_j
func TestMax(t *testing.T) {
	t.Error("TestMax is not implemented")
	return
}

// testing tooMany function
// ∀i . errs[i].String() substring of tooMany(errs...) 
// #lines in tooMany(errs...) = len(errs) + 1
func TestTooMany(t *testing.T) {
	t.Error("TestTooMany is not implemented")
	return
}
