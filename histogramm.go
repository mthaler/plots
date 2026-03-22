package main

import (
	"log"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func CreateHistogram(xs, ys []float64, file string) {
	p := plot.New()
	p.Title.Text = "Height Histogram"
	p.X.Label.Text = "height"

	h, err := plotter.NewHistogram(hplot.ZipXY(xs, ys), 100)
	if err != nil {
		log.Fatalf("could not create histogtam: %+v", err)
	}

	p.Add(h)

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}
}
