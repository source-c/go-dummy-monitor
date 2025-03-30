package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
)

// GraphDataProvider is an interface for components that provide data to be graphed
type GraphDataProvider interface {
	GetData() []float64
	GetMaxValue() float64
	GetCurrentValue() float64
	GetTitle() string
	GetColor() color.Color
}

// DualGraphDataProvider is an interface for components that provide two sets of data to be graphed
type DualGraphDataProvider interface {
	GetReadData() []float64
	GetWriteData() []float64
	GetMaxValue() float64
	GetCurrentReadValue() float64
	GetCurrentWriteValue() float64
	GetTitle() string
	GetColor() color.Color
}

// MonitorWidget is an interface for all monitoring widgets
type MonitorWidget interface {
	CreateDetailedView() *fyne.Container
	CreateCompactView() *fyne.Container
	CreateViewWithOptions(showDetails bool) *fyne.Container
}
