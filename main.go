package main

import (
	"encoding/csv"
	"log"
	"os"

	"gonum.org/v1/plot/plotter"
)

func main() {
	// The data is from NHANES (https://www.cdc.gov/nchs/nhanes/index.html)
	whs := readCsvFile("NHANES-2017-2018-height-weight.csv")

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

	var hs plotter.Values
	for i, wh := range whs {
		if i != 0 {
			h, err := parseFloat(wh[4])
			if err != nil {
				log.Fatalf("could not parse %s", wh[4])
			}
			hs = append(hs, h)
		}
	}

	var ws plotter.Values
	for i, wh := range whs {
		if i != 0 {
			w, err := parseFloat(wh[3])
			if err != nil {
				log.Fatalf("could not parse %s", wh[4])
			}
			hs = append(hs, w)
		}
	}

	CreateScatterPlot(xs, ys, "scatter.png")
	CreateHistogram(hs, "histogram_heights.png")
	CreateHistogram(ws, "histogram_heights.png")
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
