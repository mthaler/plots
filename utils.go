package main

import (
	"math"
	"strconv"
)

func toInt(x float64) int {
	n := math.Round(x)
	return int(n)
}

func parseFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}
