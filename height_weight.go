package main

type HeightWeight struct {
	height float64
	weight float64
}

func parseHeightWeight(s []string) (*HeightWeight, error) {
	h, err := parseFloat(s[0])
	if err != nil {
		return nil, err
	}
	w, err := parseFloat(s[1])
	if err != nil {
		return nil, err
	}

	return &HeightWeight{height: h, weight: w}, nil
}

func (h HeightWeight) toSlice() []float64 {
	result := make([]float64, 2)
	result[0] = h.height
	result[1] = h.weight
	return result
}
