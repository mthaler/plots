package main

import "gonum.org/v1/plot"

func CreateHistogram(xs []float64, file string) {
	p := plot.New()
	p.Title.Text = "Height Histogram"
	p.X.Label.Text = "height"
}
