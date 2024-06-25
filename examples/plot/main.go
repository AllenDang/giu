package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	linedata     []float64
	linedata2    []float64
	lineTicks    []g.PlotTicker
	bardata      []float64
	bardata2     []float64
	bardata3     []float64
	timeDataMin  float64
	timeDataMax  float64
	timeDataX    []float64
	timeDataY    []float64
	timeScatterY []float64
	scatterdata  []float64
)

func loop() {
	g.SingleWindow().Layout(
		g.Plot("Plot 基本图表").AxisLimits(0, 100, -1.2, 1.2, g.ConditionOnce).XTicks(lineTicks, false).Plots(
			g.Line("Plot Line 线图", linedata),
			g.Line("Plot Line2", linedata2),
			g.SwitchPlotAxes(g.AxisX1, g.AxisY2),
			g.Scatter("Scatter 散点图", scatterdata),
		).SetYAxisLabel(g.AxisY2, "secondary axis"),
		g.Plot("Plot Time Axe 时间线").AxisLimits(timeDataMin, timeDataMax, 0, 1, g.ConditionOnce).Plots(
			g.LineXY("Time Line 时间线", timeDataX, timeDataY),
			g.ScatterXY("Time Scatter 时间散点图", timeDataX, timeScatterY),
		),
		g.Row(
			g.Plot("Plot Bars").
				Size(500, 250).
				AxisLimits(0, 10, -1.2, 1.2, g.ConditionOnce).
				Plots(
					g.Bar("Plot Bar 柱状图", bardata),
					g.Bar("Plot Bar2", bardata2).Shift(0.2),
					g.BarH("Plot Bar H 水平柱状图", bardata3),
				),
			g.Plot("Pie Chart").
				Flags(g.PlotFlagsEqual).
				Size(250, 250).
				XAxeFlags(g.PlotAxisFlagsNoDecorations).
				YAxeFlags(g.PlotAxisFlagsNoDecorations, 0, 0).
				AxisLimits(0, 1, 0, 1, g.ConditionAlways).
				Plots(
					g.PieChart([]string{"Part 1 图例1", "Part 2", "Part 3"}, []float64{0.22, 0.38, 0.4}, 0.5, 0.5, 0.45),
				),
		),
	)
}

func main() {
	delta := 0.1
	for x := 0.0; x < 10; x += delta {
		linedata = append(linedata, math.Sin(x))
		linedata2 = append(linedata2, math.Cos(x))
		scatterdata = append(scatterdata, math.Sin(x)+0.1)
	}

	for i := 0; i < 100; i += 5 {
		lineTicks = append(lineTicks, g.PlotTicker{Position: float64(i), Label: fmt.Sprintf("P%d", i)})
	}

	delta = 1
	for x := 0.0; x < 10; x += delta {
		bardata = append(bardata, math.Sin(x))
		bardata2 = append(bardata2, math.Sin(x)-0.2)
		bardata3 = append(bardata3, rand.Float64())
	}

	for i := 0; i < 100; i++ {
		timeDataX = append(timeDataX, float64(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour*time.Duration(24*i)).Unix()))
		timeDataY = append(timeDataY, rand.Float64())
		timeScatterY = append(timeScatterY, rand.Float64())
	}

	timeDataMin = timeDataX[0]
	timeDataMax = timeDataX[len(timeDataX)-1]

	wnd := g.NewMasterWindow("Plot Demo", 1000, 900, 0)
	wnd.Run(loop)
}
