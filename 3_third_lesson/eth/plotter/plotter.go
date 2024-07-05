package plotter

import (
	"eth/csv_process"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func GeneratePlotFor(pricePairs []csv_process.EthereumPrice) (*plot.Plot, error) {
	points := make(plotter.XYs, len(pricePairs))

	for i := range pricePairs {
		points[i].X = float64(pricePairs[i].Date.Unix())
		points[i].Y = pricePairs[i].Price
	}

	pricePlot := plot.New()

	pricePlot.Title.Text = "eth pr chart"
	pricePlot.X.Label.Text = "Date"
	pricePlot.Y.Label.Text = "Time"

	if err := plotutil.AddLinePoints(pricePlot, "eth", points); err != nil {
		return nil, fmt.Errorf("err adding line %w", err)
	}

	pricePlot.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	return pricePlot, nil
}
