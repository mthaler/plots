package main

import "math"

func toInt(x float64) int {
	n := math.Round(x)
	return int(n)
}
