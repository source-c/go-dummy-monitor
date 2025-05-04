package ui

import (
	"testing"

	"go-dummy-monitor/constants"
)

func TestCreateBaseGraph(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create a base graph
	graph := factory.createBaseGraph()

	// Test basic properties
	if graph.GraphWidth != constants.GRAPH_WIDTH {
		t.Errorf("Expected GraphWidth to be %f, got %f", float32(constants.GRAPH_WIDTH), graph.GraphWidth)
	}
	if graph.GraphHeight != constants.GRAPH_HEIGHT {
		t.Errorf("Expected GraphHeight to be %f, got %f", float32(constants.GRAPH_HEIGHT), graph.GraphHeight)
	}
	if graph.GraphPadding != constants.GRAPH_PADDING {
		t.Errorf("Expected GraphPadding to be %f, got %f", float32(constants.GRAPH_PADDING), graph.GraphPadding)
	}
	if graph.ElementSpacing != constants.ELEMENT_SPACING {
		t.Errorf("Expected ElementSpacing to be %f, got %f", float32(constants.ELEMENT_SPACING), graph.ElementSpacing)
	}
	if graph.LabelHeight != constants.LABEL_HEIGHT {
		t.Errorf("Expected LabelHeight to be %f, got %f", float32(constants.LABEL_HEIGHT), graph.LabelHeight)
	}
	if graph.StrokeWidth != constants.STROKE_WIDTH {
		t.Errorf("Expected StrokeWidth to be %f, got %f", float32(constants.STROKE_WIDTH), graph.StrokeWidth)
	}
}

func TestCreateCPUWidget(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create a CPU widget
	widget := factory.CreateCPUWidget()

	// Test basic properties
	if widget == nil {
		t.Error("Expected widget to not be nil")
	}
}

func TestCreateRAMWidget(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create a RAM widget
	widget := factory.CreateRAMWidget()

	// Test basic properties
	if widget == nil {
		t.Error("Expected widget to not be nil")
	}
}

func TestCreateDiskWidget(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create a Disk widget
	widget := factory.CreateDiskWidget()

	// Test basic properties
	if widget == nil {
		t.Error("Expected widget to not be nil")
	}
}

func TestCreateNetworkWidget(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create a Network widget
	widget := factory.CreateNetworkWidget()

	// Test basic properties
	if widget == nil {
		t.Error("Expected widget to not be nil")
	}
}

func TestCreateAllWidgets(t *testing.T) {
	// Create a test system
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Create a widget factory
	factory := NewWidgetFactory(system)

	// Create all widgets
	widgets := factory.CreateAllWidgets()

	// Test that all expected widgets are created
	if widgets[CPUComponent] == nil {
		t.Error("Expected CPU widget to not be nil")
	}
	if widgets[RAMComponent] == nil {
		t.Error("Expected RAM widget to not be nil")
	}
	if widgets[DiskComponent] == nil {
		t.Error("Expected Disk widget to not be nil")
	}
	if widgets[NetworkComponent] == nil {
		t.Error("Expected Network widget to not be nil")
	}
}
