package main

import (
	"log"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHeightWeightParse(t *testing.T) {
	hw1 := HeightWeight{}
	hw2, err := hw1.parse([]string{"1", "2"})
	if err != nil {
		log.Fatal("Error while parsing height and weight")
	}
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
