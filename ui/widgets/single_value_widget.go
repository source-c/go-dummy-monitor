package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// SingleValueWidget is a widget that displays a single value graph
type SingleValueWidget struct {
	GenericGraph
	Provider GraphDataProvider
	InfoRows []InfoRow
}

// InfoRow represents a row of information in the detailed view
type InfoRow struct {
	Label    string
	GetValue func() string
}

// NewSingleValueWidget creates a new SingleValueWidget
func NewSingleValueWidget(baseGraph GenericGraph, provider GraphDataProvider, infoRows []InfoRow) *SingleValueWidget {
	return &SingleValueWidget{
		GenericGraph: baseGraph,
		Provider:     provider,
		InfoRows:     infoRows,
	}
}

// CreateViewWithOptions creates a view with optional details
func (s *SingleValueWidget) CreateViewWithOptions(showDetails bool) *fyne.Container {
	if showDetails {
		return s.CreateDetailedView()
	}
	return s.CreateCompactView()
}

// CreateCompactView creates a compact view of the widget
func (s *SingleValueWidget) CreateCompactView() *fyne.Container {
	graphContainer, graphBg := s.CreateGraphContainer()

	// Draw the actual graph
	drawGraph := func() {
		graphContainer.Objects = []fyne.CanvasObject{graphBg}

		// Get container width for responsive layout
		containerWidth := graphBg.Size().Width
		if containerWidth < s.GraphWidth {
			containerWidth = s.GraphWidth
		}

		// Add header with combined info
		s.AddTitle(graphContainer, fmt.Sprintf("%s: %.1f",
			s.Provider.GetTitle(), s.Provider.GetCurrentValue()), s.TextColor)

		// Add axis labels
		s.AddAxisLabels(graphContainer, "0", fmt.Sprintf("%.0f", s.Provider.GetMaxValue()))

		// Add graph border
		s.AddGraphBorder(graphContainer, containerWidth)

		// Draw the graph lines
		s.DrawSingleGraph(
			graphContainer,
			s.Provider.GetData(),
			s.Provider.GetMaxValue(),
			s.Provider.GetColor(),
			containerWidth,
		)

		canvas.Refresh(graphContainer)
	}

	drawGraph()
	return graphContainer
}

// CreateDetailedView creates a detailed view of the widget
func (s *SingleValueWidget) CreateDetailedView() *fyne.Container {
	// Create left column for graph visualization
	graphContainer, graphBg := s.CreateGraphContainer()

	// Create right column for info display
	infoContainer := container.NewVBox()

	// Info content
	infoTitle := widget.NewLabel(fmt.Sprintf("%s INFO", s.Provider.GetTitle()))
	infoTitle.TextStyle = fyne.TextStyle{Bold: true}
	infoContainer.Add(infoTitle)

	// Create dynamic info rows
	infoLabels := make([]*widget.Label, len(s.InfoRows))
	for i, row := range s.InfoRows {
		infoLabels[i] = widget.NewLabel(fmt.Sprintf("%s: %s", row.Label, row.GetValue()))
		infoContainer.Add(infoLabels[i])
	}

	// Draw the actual graph
	drawGraph := func() {
		graphContainer.Objects = []fyne.CanvasObject{graphBg}

		// Get container width for responsive layout
		containerWidth := graphBg.Size().Width
		if containerWidth < s.GraphWidth {
			containerWidth = s.GraphWidth
		}

		// Add header information
		titleLabel := s.AddTitle(graphContainer, s.Provider.GetTitle(), s.TextColor)
		titleLabel.Move(fyne.NewPos(s.ElementSpacing, s.ElementSpacing))

		valueLabel := canvas.NewText(fmt.Sprintf("%.2f", s.Provider.GetCurrentValue()), s.Provider.GetColor())
		valueLabel.Move(fyne.NewPos(s.ElementSpacing*5, s.ElementSpacing))
		graphContainer.Add(valueLabel)

		// Add axis labels
		s.AddAxisLabels(graphContainer, "0", fmt.Sprintf("%.0f", s.Provider.GetMaxValue()))

		// Add graph border
		s.AddGraphBorder(graphContainer, containerWidth)

		// Draw the graph lines
		s.DrawSingleGraph(
			graphContainer,
			s.Provider.GetData(),
			s.Provider.GetMaxValue(),
			s.Provider.GetColor(),
			containerWidth,
		)

		// Update info rows
		for i, row := range s.InfoRows {
			infoLabels[i].SetText(fmt.Sprintf("%s: %s", row.Label, row.GetValue()))
		}

		canvas.Refresh(graphContainer)
	}

	drawGraph()

	// Create two-column layout
	mainContainer := container.New(layout.NewGridLayoutWithColumns(2), graphContainer, infoContainer)

	return mainContainer
}
