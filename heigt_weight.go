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
