package widgets

import (
	"testing"

	"go-dummy-monitor/constants"

	"fyne.io/fyne/v2/test"
)

func TestNewGenericGraph(t *testing.T) {
	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	if graph == nil {
		t.Error("Expected non-nil graph, got nil")
	}
}

func TestCreateGraphContainer(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	container, bg := graph.CreateGraphContainer()
	if container == nil {
		t.Fatal("Expected non-nil container, got nil")
	}
	if bg == nil {
		t.Fatal("Expected non-nil background, got nil")
	}

	// Check background size
	size := bg.Size()
	if size.Width != float32(constants.GRAPH_WIDTH) {
		t.Errorf("Expected background width to be %d, got %f", constants.GRAPH_WIDTH, size.Width)
	}
	if size.Height != float32(constants.GRAPH_HEIGHT) {
		t.Errorf("Expected background height to be %d, got %f", constants.GRAPH_HEIGHT, size.Height)
	}
}

func TestAddGraphBorder(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	container, _ := graph.CreateGraphContainer()
	border := graph.AddGraphBorder(container, float32(constants.GRAPH_WIDTH))
	if border == nil {
		t.Error("Expected non-nil border, got nil")
	}
}

func TestAddAxisLabels(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	container, _ := graph.CreateGraphContainer()
	graph.AddAxisLabels(container, "0", "100")
}

func TestDrawSingleGraph(t *testing.T) {
	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	container, _ := graph.CreateGraphContainer()
	data := []float64{10, 20, 30, 40, 50}
	graph.DrawSingleGraph(container, data, 100, constants.LightColors.Grid, float32(constants.GRAPH_WIDTH))
}

func TestDrawDualGraph(t *testing.T) {
	graph := NewGenericGraph(
		constants.GRAPH_WIDTH,
		constants.GRAPH_HEIGHT,
		constants.GRAPH_PADDING,
		constants.ELEMENT_SPACING,
		constants.LABEL_HEIGHT,
		constants.LightColors.BG,
		constants.LightColors.Grid,
		constants.LightColors.Text,
		constants.STROKE_WIDTH,
		constants.EmptyRectangle,
		constants.TRANSLUCENT_ALPHA,
	)

	container, _ := graph.CreateGraphContainer()
	primaryData := []float64{10, 20, 30, 40, 50}
	secondaryData := []float64{5, 15, 25, 35, 45}
	graph.DrawDualGraph(container, primaryData, secondaryData, 100, constants.LightColors.Grid, constants.LightColors.Grid, float32(constants.GRAPH_WIDTH))
}
