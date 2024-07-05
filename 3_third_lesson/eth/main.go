package main

import (
	"eth/csv_process"
	"eth/plotter"
	"fmt"
	"gonum.org/v1/plot/vg"
	"os"
)

func main() {

	pricePairs, err := csv_process.LoadDataFrom("prices.csv")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading from csv %v%n", err)
		os.Exit(1)
	}

	pricesPlot, err := plotter.GeneratePlotFor(pricePairs)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating %v%n", err)
		os.Exit(1)
	}

	if err := pricesPlot.Save(15*vg.Inch, 4*vg.Inch, "eth_pr.png"); err != nil {
		fmt.Fprintf(os.Stderr, "error generating %v%n", err)
		os.Exit(1)
	}
}
