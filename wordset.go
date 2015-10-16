package main

import (
	"fmt"
)

type wordset map[string]float64

var englishFrequency = wordset{
	"this": 1.0,
}

func trivial() {
	fmt.Println(englishFrequency["that"])
}