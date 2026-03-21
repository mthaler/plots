package main

import (
	"encoding/csv"
	"log"
	"os"

	"gonum.org/v1/plot"
)

func main() {
	p := plot.New()
	p.Title.Text = "Height - Weight Plot"
	p.X.Label.Text = "heigt"
	p.Y.Label.Text = "weight"
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
