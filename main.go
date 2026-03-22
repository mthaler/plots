package main

import (
	"encoding/csv"
	"log"
	"os"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// The data is from NHANES (https://www.cdc.gov/nchs/nhanes/index.html)
	whs := readCsvFile("NHANES-2017-2018-height-weight.csv")

	p := plot.New()
	p.Title.Text = "Height - Weight Plot"
	p.X.Label.Text = "heigt"
	p.Y.Label.Text = "weight"

	xs := make([]float64, 0)
	ys := make([]float64, 0)

	for i, ws := range whs {
		if i != 0 {
			h, err := parseFloat(ws[4])
			if err != nil {
				log.Fatalf("could not parse %s", ws[4])
			}
			xs = append(xs, h)
		}
	}

	for i, ws := range whs {
		if i != 0 {
			w, err := parseFloat(ws[3])
			if err != nil {
				log.Fatalf("could not parse %s", ws[3])
			}
			ys = append(ys, w)
		}
	}

	err := plotutil.AddScatters(p, "f(x) = 2*x", hplot.ZipXY(xs, ys))
	if err != nil {
		log.Fatalf("could not create scatters: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, "scatter.png")
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
