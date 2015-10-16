package main

import (
	"testing"
)

func TestWordsetCommonNonmember(t *testing.T) {
	nonsense := "xkxcd"
	if englishFrequency[nonsense] != 0.0 {
		t.Error("English frequency information contains nonzero instance of '"+nonsense+"'")
	}
}