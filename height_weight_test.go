package main

import (
	"log"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHeightWeightParse(t *testing.T) {
	hw, err := parseHeightWeight([]string{"1", "2"})
	if err != nil {
		log.Fatal("Error while parsing height and weight")
	}
	h := toInt(hw.height)
	if h != 1 {
		t.Errorf("Height = %d, expected 1", h)
	}
	w := toInt(hw.weight)
	if w != 2 {
		t.Errorf("Weight = %d, expected 2", w)
	}
}
