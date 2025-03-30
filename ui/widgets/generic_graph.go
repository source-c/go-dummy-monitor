package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// GenericGraph contains common functionality for all graph widgets
type GenericGraph struct {
	GraphWidth       float32
	GraphHeight      float32
	GraphPadding     float32
	ElementSpacing   float32
	LabelHeight      float32
	BackgroundColor  color.Color
	GridColor        color.Color
	TextColor        color.Color
	StrokeWidth      float32
	EmptyRectangle   color.Color
	TranslucentAlpha uint8
}

// NewGenericGraph creates a new GenericGraph with default settings
func NewGenericGraph(graphWidth, graphHeight, graphPadding, elementSpacing, labelHeight float32,
	backgroundColor, gridColor, textColor color.Color, strokeWidth float32,
	emptyRectangle color.Color, translucentAlpha uint8) *GenericGraph {
	return &GenericGraph{
		GraphWidth:       graphWidth,
		GraphHeight:      graphHeight,
		GraphPadding:     graphPadding,
		ElementSpacing:   elementSpacing,
		LabelHeight:      labelHeight,
		BackgroundColor:  backgroundColor,
		GridColor:        gridColor,
		TextColor:        textColor,
		StrokeWidth:      strokeWidth,
		EmptyRectangle:   emptyRectangle,
		TranslucentAlpha: translucentAlpha,
	}
}

// CreateGraphContainer creates a basic graph container with background
func (b *GenericGraph) CreateGraphContainer() (*fyne.Container, *canvas.Rectangle) {
	graphContainer := container.NewWithoutLayout()

	// Create responsive background for graph area
	graphBg := canvas.NewRectangle(b.BackgroundColor)
	graphBg.SetMinSize(fyne.NewSize(b.GraphWidth, b.GraphHeight))
	graphContainer.Add(graphBg)

	return graphContainer, graphBg
}

// AddGraphBorder adds a border to the graph area
func (b *GenericGraph) AddGraphBorder(graphContainer *fyne.Container, containerWidth float32) *canvas.Rectangle {
	// Graph border - resize to fit container width
	border := canvas.NewRectangle(b.EmptyRectangle)
	border.StrokeWidth = b.StrokeWidth
	border.StrokeColor = b.GridColor
	border.Resize(fyne.NewSize(containerWidth-b.ElementSpacing*2, b.GraphPadding*2.5))
	border.Move(fyne.NewPos(b.ElementSpacing, b.GraphPadding))

	graphContainer.Add(border)
	return border
}

// AddAxisLabels adds Y axis labels to the graph
func (b *GenericGraph) AddAxisLabels(graphContainer *fyne.Container, min string, max string) {
	// When the theme changes, GenericGraph.TextColor gets updated via factory.createBaseGraph()
	// So each time we draw we're using the current theme's text color
	axisY0 := canvas.NewText(min, b.TextColor)
	axisY0.Move(fyne.NewPos(b.ElementSpacing/2, b.GraphPadding+b.GraphPadding*2.5))

	axisYMax := canvas.NewText(max, b.TextColor)
	axisYMax.Move(fyne.NewPos(b.ElementSpacing/2, b.GraphPadding))

	graphContainer.Add(axisY0)
	graphContainer.Add(axisYMax)
}

// AddTitle adds a title to the graph
func (b *GenericGraph) AddTitle(graphContainer *fyne.Container, title string, color color.Color) *canvas.Text {
	titleLabel := canvas.NewText(title, color)
	titleLabel.Move(fyne.NewPos(b.ElementSpacing, b.ElementSpacing))
	graphContainer.Add(titleLabel)
	return titleLabel
}

// AddSubtitle adds a subtitle to the graph
func (b *GenericGraph) AddSubtitle(graphContainer *fyne.Container, subtitle string, color color.Color, textSize float32, yOffset float32) *canvas.Text {
	subtitleLabel := canvas.NewText(subtitle, color)
	subtitleLabel.Move(fyne.NewPos(b.ElementSpacing, yOffset))
	subtitleLabel.TextSize = textSize
	graphContainer.Add(subtitleLabel)
	return subtitleLabel
}

// DrawLine draws a line on the graph
func (b *GenericGraph) DrawLine(graphContainer *fyne.Container, x1, y1, x2, y2 float32, lineColor color.Color, strokeWidth float32) {
	line := canvas.NewLine(lineColor)
	line.Position1 = fyne.NewPos(x1, y1)
	line.Position2 = fyne.NewPos(x2, y2)
	line.StrokeWidth = strokeWidth
	graphContainer.Add(line)
}

// DrawSingleGraph draws a single dataset graph
func (b *GenericGraph) DrawSingleGraph(graphContainer *fyne.Container, data []float64, maxValue float64, color color.Color, containerWidth float32) {
	// Calculate points spacing based on container width
	pointSpacing := (containerWidth - b.LabelHeight) / float32(len(data))

	// Draw graph lines with adaptive spacing
	for i := 1; i < len(data); i++ {
		x1 := b.ElementSpacing + pointSpacing*float32(i-1)
		y1 := float32(b.GraphPadding*2.5-(float32(data[i-1])*b.GraphPadding*2.5/float32(maxValue))) + b.GraphPadding
		x2 := b.ElementSpacing + pointSpacing*float32(i)
		y2 := float32(b.GraphPadding*2.5-(float32(data[i])*b.GraphPadding*2.5/float32(maxValue))) + b.GraphPadding

		b.DrawLine(graphContainer, x1, y1, x2, y2, color, b.StrokeWidth)
	}
}

// DrawDualGraph draws a graph with two datasets
func (b *GenericGraph) DrawDualGraph(graphContainer *fyne.Container, primaryData []float64, secondaryData []float64,
	maxValue float64, primaryColor color.Color, secondaryColor color.Color, containerWidth float32) {

	// Calculate points spacing based on container width
	pointSpacing := (containerWidth - b.LabelHeight) / float32(len(primaryData))

	// Draw primary data
	b.DrawSingleGraph(graphContainer, primaryData, maxValue, primaryColor, containerWidth)

	// Use slightly different color for secondary data
	if rgba, ok := secondaryColor.(color.RGBA); ok {
		rgba.A = b.TranslucentAlpha // Make it slightly translucent to differentiate
		secondaryColor = rgba
	}

	// Draw secondary data
	for i := 1; i < len(secondaryData); i++ {
		x1 := b.ElementSpacing + pointSpacing*float32(i-1)
		y1 := float32(b.GraphPadding*2.5-(float32(secondaryData[i-1])*b.GraphPadding*2.5/float32(maxValue))) + b.GraphPadding
		x2 := b.ElementSpacing + pointSpacing*float32(i)
		y2 := float32(b.GraphPadding*2.5-(float32(secondaryData[i])*b.GraphPadding*2.5/float32(maxValue))) + b.GraphPadding

		b.DrawLine(graphContainer, x1, y1, x2, y2, secondaryColor, b.StrokeWidth)
	}
}
