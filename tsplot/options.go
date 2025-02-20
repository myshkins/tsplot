package tsplot

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

// PlotOption defines the type used to configure the underlying *plot.Plot.
// A function that returns PlotOption can be used to set options on the *plot.Plot.
type PlotOption func(p *plot.Plot)

// WithForeground sets the foreground colors of the *plot.Plot to the provided color.
func WithForeground(color color.Color) PlotOption {
	return func(p *plot.Plot) {
		p.Title.TextStyle.Color = color
		p.X.Color = color
		p.X.Label.TextStyle.Color = color
		p.X.Tick.Color = color
		p.X.Tick.Label.Color = color
		p.Y.Tick.Label.Color = color
		p.Y.Color = color
		p.Y.Label.TextStyle.Color = color
		p.Y.Tick.Color = color
		p.Y.Tick.Label.Color = color
	}
}

// WithBackground sets the background color on the *plot.Plot.
func WithBackground(color color.Color) PlotOption {
	return func(p *plot.Plot) {
		p.BackgroundColor = color
	}
}

// WithTitle configures the title text of the *plot.Plot.
func WithTitle(title string) PlotOption {
	return func(p *plot.Plot) {
		p.Title.Text = title
	}
}

// WithGrid adds a grid to the *plot.Plot and sets the color.
func WithGrid(color color.Color) PlotOption {
	return func(p *plot.Plot) {
		grid := plotter.NewGrid()
		grid.Horizontal.Color = color
		grid.Vertical.Color = color
		p.Add(grid)
	}
}

// WithLegend sets the text color of the legend.
func WithLegend(color color.Color) PlotOption {
	return func(p *plot.Plot) {
		p.Legend.TextStyle.Color = color
	}
}

// WithXAxisName sets the name of the X Axis.
func WithXAxisName(name string) PlotOption {
	return func(p *plot.Plot) {
		p.X.Label.Text = name
	}
}

// WithYAxisName sets the name of the Y Axis.
func WithYAxisName(name string) PlotOption {
	return func(p *plot.Plot) {
		p.Y.Label.Text = name
	}
}

// WithXTimeTicks configures the format for the tick marks on the X Axis.
func WithXTimeTicks(format string) PlotOption {
	return func(p *plot.Plot) {
		p.X.Tick.Marker = plot.TimeTicks{
			Format: format,
		}
	}
}

// WithFontSize configures the font sizes for the plot's Title, Axis', and Legend.
func WithFontSize(size float64) PlotOption {
	return func(p *plot.Plot) {
		p.Title.TextStyle.Font.Size = font.Length(size)
		p.X.Label.TextStyle.Font.Size = font.Length(size)
		p.Y.Label.TextStyle.Font.Size = font.Length(size)
		p.Legend.TextStyle.Font.Size = font.Length(size)
	}
}

// ApplyDefaultHighContrast applies the default high contrast color scheme to the *plot.Plot.
func ApplyDefaultHighContrast(p *plot.Plot) {
	opts := []PlotOption{
		WithBackground(DefaultColors_HighContrast.Background),
		WithForeground(DefaultColors_HighContrast.Foreground),
		WithLegend(DefaultColors_HighContrast.Foreground),
	}
	for _, opt := range opts {
		opt(p)
	}
}

// WithLineFromPoints creates a *plotter.Line from the passed in data points and adds
// it to the plot.
func WithLineFromPoints(pts []*monitoringpb.Point) PlotOption {
	return func(p *plot.Plot) {
		line, _ := createLine(pts)
		line.Width = vg.Points(1)
		p.Add(line)
	}
}

func WithColoredLineFromPoints(pts []*monitoringpb.Point, color color.Color) PlotOption {
	return func(p *plot.Plot) {
		line, _ := createLine(pts)
		line.Width = vg.Points(1)
		line.Color = color
		p.Add(line)
	}
}
