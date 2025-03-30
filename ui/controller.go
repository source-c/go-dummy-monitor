package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"go-dummy-monitor/ui/widgets"
)

// WidgetController manages a widget and handles its updates
type WidgetController struct {
	System     MonitoringSystem
	Widget     widgets.MonitorWidget
	Container  *fyne.Container
	ShowDetail bool
}

// NewWidgetController creates a new widget controller
func NewWidgetController(system MonitoringSystem, widget widgets.MonitorWidget, showDetail bool) *WidgetController {
	controller := &WidgetController{
		System:     system,
		Widget:     widget,
		ShowDetail: showDetail,
	}

	// Create initial container
	controller.Container = container.NewPadded(widget.CreateViewWithOptions(showDetail))

	return controller
}

// Update updates the widget with the latest data
func (c *WidgetController) Update() {
	// Get a fresh widget view with latest data
	newView := c.Widget.CreateViewWithOptions(c.ShowDetail)

	// Replace the old view with the new one
	if len(c.Container.Objects) > 0 {
		c.Container.Objects[0] = newView
	} else {
		c.Container.Objects = []fyne.CanvasObject{newView}
	}

	// Ensure the container is refreshed to show new data
	c.Container.Refresh()

	// Widget updated successfully
}

// SetShowDetail sets whether to show detailed information
func (c *WidgetController) SetShowDetail(showDetail bool) {
	if c.ShowDetail != showDetail {
		c.ShowDetail = showDetail
		c.Update()
	}
}

// MonitoringPanel manages all widget controllers
type MonitoringPanel struct {
	System        MonitoringSystem
	WidgetFactory *WidgetFactory
	Container     *fyne.Container
	Controllers   map[ComponentType]*WidgetController
	ShowDetail    bool
}

// NewMonitoringPanel creates a new monitoring panel
func NewMonitoringPanel(system MonitoringSystem, factory *WidgetFactory, showDetail bool) *MonitoringPanel {
	panel := &MonitoringPanel{
		System:        system,
		WidgetFactory: factory,
		ShowDetail:    showDetail,
		Controllers:   make(map[ComponentType]*WidgetController),
	}

	// Create widget controllers
	allWidgets := factory.CreateAllWidgets()

	panel.Controllers[CPUComponent] = NewWidgetController(system, allWidgets[CPUComponent], showDetail)
	panel.Controllers[RAMComponent] = NewWidgetController(system, allWidgets[RAMComponent], showDetail)
	panel.Controllers[DiskComponent] = NewWidgetController(system, allWidgets[DiskComponent], showDetail)
	panel.Controllers[NetworkComponent] = NewWidgetController(system, allWidgets[NetworkComponent], showDetail)

	// Create panel container
	panel.Container = container.New(layout.NewVBoxLayout(),
		panel.Controllers[CPUComponent].Container,
		panel.Controllers[RAMComponent].Container,
		panel.Controllers[DiskComponent].Container,
		panel.Controllers[NetworkComponent].Container,
	)

	return panel
}

// Update updates all widget controllers
func (p *MonitoringPanel) Update() {
	for _, controller := range p.Controllers {
		controller.Update()
	}
}

// SetShowDetail sets whether to show detailed information
func (p *MonitoringPanel) SetShowDetail(showDetail bool) {
	if p.ShowDetail != showDetail {
		p.ShowDetail = showDetail
		for _, controller := range p.Controllers {
			controller.SetShowDetail(showDetail)
		}
	}
}
