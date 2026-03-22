package main

import (
	"errors"
	"log"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// NewScatter returns a Scatter that uses the
// default glyph style.
func NewScatterPlot() error {
	return errors.New("Cannot create a scatter plot")
}

func CreateScatterPlot(xs, ys []float64) {
	p := plot.New()
	p.Title.Text = "Height - Weight Plot"
	p.X.Label.Text = "height"
	p.Y.Label.Text = "weight"

	err := plotutil.AddScatters(p, "", hplot.ZipXY(xs, ys))
	if err != nil {
		log.Fatalf("could not create scatters: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, "scatter.png")
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}
}
