package main

import (
	"errors"
)

type HeightWeight struct {
	height float64
	weight float64
}

func (h HeightWeight) parse(text []string) (*HeightWeight, error) {
	return nil, errors.New("could not parse height and weight")
}

func (h HeightWeight) toSlice() []float64 {
	result := make([]float64, 2)
	result[0] = h.height
	result[1] = h.weight
	return result
}
