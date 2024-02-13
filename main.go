package main

import (
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

const (
	buyColor  = "#00AB6B"
	sellColor = "#EE4A49"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(200000)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ChartLine}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Trade Book Chart - BBCA",
			Subtitle: "Trade Book Chart by Lot Buy/Sell Per-5 minutes",
		}),
	)

	// Put data into instance
	line.SetXAxis([]string{"09:00", "09:05", "09:10", "09:15", "09:20", "09:25", "09:30"}).
		AddSeries("Beli", generateLineItems(),
			charts.WithLineStyleOpts(opts.LineStyle{Color: buyColor}),
			charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{a} - {c} Lot"})).
		AddSeries("Jual", generateLineItems(),
			charts.WithLineStyleOpts(opts.LineStyle{Color: sellColor}),
			charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{a} - {c} Lot"})).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false, ShowSymbol: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
