package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"go-dummy-monitor/ui/widgets"
)

// WidgetFactory creates monitoring widgets
type WidgetFactory struct {
	System MonitoringSystem
}

// NewWidgetFactory creates a new widget factory
func NewWidgetFactory(system MonitoringSystem) *WidgetFactory {
	return &WidgetFactory{
		System: system,
	}
}

// createBaseGraph creates a base graph with system theme settings
func (f *WidgetFactory) createBaseGraph() *widgets.GenericGraph {
	colorScheme := f.System.GetColorScheme()

	return widgets.NewGenericGraph(
		GraphWidth,
		GraphHeight,
		GraphPadding,
		ElementSpacing,
		LabelHeight,
		colorScheme.BG,
		colorScheme.Grid,
		colorScheme.Text,
		StrokeWidth,
		f.System.GetEmptyRectangle(),
		TranslucentAlpha,
	)
}

// CreateCPUWidget creates a CPU monitoring widget
func (f *WidgetFactory) CreateCPUWidget() widgets.MonitorWidget {
	baseGraph := f.createBaseGraph()
	provider := &CPUDataProvider{System: f.System}
	infoRows := GetCPUInfoProvider(f.System)

	return widgets.NewSingleValueWidget(*baseGraph, provider, infoRows)
}

// CreateRAMWidget creates a RAM monitoring widget
func (f *WidgetFactory) CreateRAMWidget() widgets.MonitorWidget {
	baseGraph := f.createBaseGraph()
	provider := &RAMDataProvider{System: f.System}
	infoRows := GetRAMInfoProvider(f.System)

	return widgets.NewSingleValueWidget(*baseGraph, provider, infoRows)
}

// CreateDiskWidget creates a Disk monitoring widget
func (f *WidgetFactory) CreateDiskWidget() widgets.MonitorWidget {
	baseGraph := f.createBaseGraph()
	provider := &DiskDataProvider{System: f.System}
	infoRows := GetDiskInfoProvider(f.System)

	return widgets.NewDualValueWidget(
		*baseGraph,
		provider,
		infoRows,
		// infoRows already contains the properly formatted IO values
		"",
		"",
		"R:%.1f W:%.1f MB/s",
	)
}

// CreateNetworkWidget creates a Network monitoring widget
func (f *WidgetFactory) CreateNetworkWidget() widgets.MonitorWidget {
	baseGraph := f.createBaseGraph()
	provider := &NetworkDataProvider{System: f.System}
	infoRows := GetNetworkInfoProvider(f.System)

	return widgets.NewDualValueWidget(
		*baseGraph,
		provider,
		infoRows,
		"",
		"",
		"D:%.1f U:%.1f MB/s",
	)
}

// CreateAllWidgets creates all system monitoring widgets
func (f *WidgetFactory) CreateAllWidgets() map[ComponentType]widgets.MonitorWidget {
	return map[ComponentType]widgets.MonitorWidget{
		CPUComponent:     f.CreateCPUWidget(),
		RAMComponent:     f.CreateRAMWidget(),
		DiskComponent:    f.CreateDiskWidget(),
		NetworkComponent: f.CreateNetworkWidget(),
	}
}

// CreateMonitoringPanel creates a monitoring panel with all widgets
func (f *WidgetFactory) CreateMonitoringPanel(showDetails bool) *fyne.Container {
	allWidgets := f.CreateAllWidgets()

	cpuWidget := container.NewPadded(allWidgets[CPUComponent].CreateViewWithOptions(showDetails))
	ramWidget := container.NewPadded(allWidgets[RAMComponent].CreateViewWithOptions(showDetails))
	diskWidget := container.NewPadded(allWidgets[DiskComponent].CreateViewWithOptions(showDetails))
	netWidget := container.NewPadded(allWidgets[NetworkComponent].CreateViewWithOptions(showDetails))

	panel := container.New(layout.NewVBoxLayout(),
		cpuWidget,
		ramWidget,
		diskWidget,
		netWidget,
	)

	return panel
}
