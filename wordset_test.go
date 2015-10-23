package main

import (
	"testing"
)

func TestWordsetCommonNonmember(t *testing.T) {
	nonsense := "xkxcd"
	if englishFrequency[nonsense].Seen != 0 {
		t.Error("English frequency information contains nonzero instance of '" + nonsense + "'")
	}
}
