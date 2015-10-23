package main

import (
	"fmt"
)

type frequency struct {
	Seen  int
	OutOf int
}

type wordset map[string]frequency

var englishFrequency = wordset{
	"this": frequency{1, 1},
}

func trivial() {
	fmt.Println(englishFrequency["that"])
}
