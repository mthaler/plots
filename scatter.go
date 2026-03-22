package main

import (
	"log"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreateScatterPlot(xs, ys []float64, file string) {
	p := plot.New()
	p.Title.Text = "Height - Weight Plot"
	p.X.Label.Text = "height"
	p.Y.Label.Text = "weight"

	err := plotutil.AddScatters(p, "", hplot.ZipXY(xs, ys))
	if err != nil {
		log.Fatalf("could not create scatters: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}
}
