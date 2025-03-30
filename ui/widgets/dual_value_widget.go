package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// DualValueWidget is a widget that displays a graph with two sets of data
type DualValueWidget struct {
	GenericGraph
	Provider        DualGraphDataProvider
	InfoRows        []InfoRow
	ReadLabel       string
	WriteLabel      string
	CompactValueFmt string // Format string for compact view values (e.g. "D:%.1f U:%.1f MB/s")
}

// NewDualValueWidget creates a new DualValueWidget
func NewDualValueWidget(
	baseGraph GenericGraph,
	provider DualGraphDataProvider,
	infoRows []InfoRow,
	readLabel string,
	writeLabel string,
	compactValueFmt string,
) *DualValueWidget {
	return &DualValueWidget{
		GenericGraph:    baseGraph,
		Provider:        provider,
		InfoRows:        infoRows,
		ReadLabel:       readLabel,
		WriteLabel:      writeLabel,
		CompactValueFmt: compactValueFmt,
	}
}

// CreateViewWithOptions creates a view with optional details
func (d *DualValueWidget) CreateViewWithOptions(showDetails bool) *fyne.Container {
	if showDetails {
		return d.CreateDetailedView()
	}
	return d.CreateCompactView()
}

// CreateCompactView creates a compact view of the widget
func (d *DualValueWidget) CreateCompactView() *fyne.Container {
	graphContainer, graphBg := d.CreateGraphContainer()

	// Draw the actual graph
	drawGraph := func() {
		graphContainer.Objects = []fyne.CanvasObject{graphBg}

		// Get container width for responsive layout
		containerWidth := graphBg.Size().Width
		if containerWidth < d.GraphWidth {
			containerWidth = d.GraphWidth
		}

		// Add header with title
		d.AddTitle(graphContainer, fmt.Sprintf("%s", d.Provider.GetTitle()), d.TextColor)

		// Add IO info as subtitle
		d.AddSubtitle(
			graphContainer,
			fmt.Sprintf(d.CompactValueFmt, d.Provider.GetCurrentReadValue(), d.Provider.GetCurrentWriteValue()),
			d.TextColor,
			12,
			d.LabelHeight+5,
		)

		// Add axis labels
		d.AddAxisLabels(graphContainer, "0", fmt.Sprintf("%.0f", d.Provider.GetMaxValue()))

		// Add graph border
		d.AddGraphBorder(graphContainer, containerWidth)

		// Draw the graph lines for both datasets
		d.DrawDualGraph(
			graphContainer,
			d.Provider.GetReadData(),
			d.Provider.GetWriteData(),
			d.Provider.GetMaxValue(),
			d.Provider.GetColor(),
			d.Provider.GetColor(),
			containerWidth,
		)

		canvas.Refresh(graphContainer)
	}

	drawGraph()
	return graphContainer
}

// CreateDetailedView creates a detailed view of the widget
func (d *DualValueWidget) CreateDetailedView() *fyne.Container {
	// Create left column for graph visualization
	graphContainer, graphBg := d.CreateGraphContainer()

	// Create right column for info display
	infoContainer := container.NewVBox()

	// Info content
	infoTitle := widget.NewLabel(fmt.Sprintf("%s INFO", d.Provider.GetTitle()))
	infoTitle.TextStyle = fyne.TextStyle{Bold: true}
	infoContainer.Add(infoTitle)

	// Create info for read/write values
	var readInfo *widget.Label
	if len(d.ReadLabel) > 0 {
		readInfo = widget.NewLabel(fmt.Sprintf("%s: %.2f", d.ReadLabel, d.Provider.GetCurrentReadValue()))
		infoContainer.Add(readInfo)
	}
	var writeInfo *widget.Label
	if len(d.WriteLabel) > 0 {
		writeInfo = widget.NewLabel(fmt.Sprintf("%s: %.2f", d.WriteLabel, d.Provider.GetCurrentWriteValue()))
		infoContainer.Add(writeInfo)
	}

	// Create dynamic info rows
	infoLabels := make([]*widget.Label, len(d.InfoRows))
	for i, row := range d.InfoRows {
		infoLabels[i] = widget.NewLabel(fmt.Sprintf("%s: %s", row.Label, row.GetValue()))
		infoContainer.Add(infoLabels[i])
	}

	// Draw the actual graph
	drawGraph := func() {
		graphContainer.Objects = []fyne.CanvasObject{graphBg}

		// Get container width for responsive layout
		containerWidth := graphBg.Size().Width
		if containerWidth < d.GraphWidth {
			containerWidth = d.GraphWidth
		}

		// Add header information
		titleLabel := d.AddTitle(graphContainer, d.Provider.GetTitle(), d.TextColor)
		titleLabel.Move(fyne.NewPos(d.ElementSpacing, d.ElementSpacing))

		valueLabel := canvas.NewText(
			fmt.Sprintf(d.CompactValueFmt, d.Provider.GetCurrentReadValue(), d.Provider.GetCurrentWriteValue()),
			d.TextColor,
		)
		valueLabel.Move(fyne.NewPos(d.GraphPadding*2, d.ElementSpacing))
		graphContainer.Add(valueLabel)

		// Add axis labels
		d.AddAxisLabels(graphContainer, "0", fmt.Sprintf("%.0f", d.Provider.GetMaxValue()))

		// Add graph border
		d.AddGraphBorder(graphContainer, containerWidth)

		// Draw the graph lines for both datasets
		d.DrawDualGraph(
			graphContainer,
			d.Provider.GetReadData(),
			d.Provider.GetWriteData(),
			d.Provider.GetMaxValue(),
			d.Provider.GetColor(),
			d.Provider.GetColor(),
			containerWidth,
		)

		// Update read/write info if any
		if readInfo != nil {
			readInfo.SetText(fmt.Sprintf("%s: %.2f", d.ReadLabel, d.Provider.GetCurrentReadValue()))
		}
		if writeInfo != nil {
			writeInfo.SetText(fmt.Sprintf("%s: %.2f", d.WriteLabel, d.Provider.GetCurrentWriteValue()))
		}

		// Update info rows
		for i, row := range d.InfoRows {
			infoLabels[i].SetText(fmt.Sprintf("%s: %s", row.Label, row.GetValue()))
		}

		canvas.Refresh(graphContainer)
	}

	drawGraph()

	// Create two-column layout
	mainContainer := container.New(layout.NewGridLayoutWithColumns(2), graphContainer, infoContainer)

	return mainContainer
}
