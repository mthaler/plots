package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func CreateHistogram(xs plotter.Values, file string, title string) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "height"

	h, err := plotter.NewHist(xs, 100)
	if err != nil {
		log.Fatalf("could not create histogtam: %+v", err)
	}

	p.Add(h)

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}
}
