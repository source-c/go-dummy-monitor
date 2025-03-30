package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"go-dummy-monitor/constants"
	"go-dummy-monitor/ui"
)

// Application-specific constants
const (
	DATA_POINTS = 11 // Number of data points to track
)

func main() {
	a := app.New()
	w := a.NewWindow("GO System Monitor")
	darkMode := false

	// Set initial theme and colors
	a.Settings().SetTheme(theme.LightTheme())

	// Initialize monitoring system
	monitorSystem := ui.NewMonitorSystem(
		125,   // Max network speed in MB/s (125 MB/s is the 1000Mbit/s or 1Gbps)
		false, // Dark mode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle, // Empty rectangle color
		DATA_POINTS,              // Number of data points to track
	)

	// Create widget factory
	widgetFactory := ui.NewWidgetFactory(monitorSystem)

	// Variables to track layout state
	showDetailColumns := true
	var currentWidth, currentHeight float32

	// Create theme toggle button
	themeButton := widget.NewButton("Toggle Theme", func() {
		darkMode = !darkMode
		if darkMode {
			a.Settings().SetTheme(theme.DarkTheme())
			monitorSystem.UpdateTheme(true, constants.LightColors, constants.DarkColors)
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			monitorSystem.UpdateTheme(false, constants.LightColors, constants.DarkColors)
		}

		// Update the UI to use the new theme colors
		// This will be set after monitoringPanel is created
	})

	// Calculate responsive initial window size based on typical screen sizes
	primaryWindow := a.NewWindow("Temp")
	// Close immediately, just used to get screen info
	primaryWindow.Close()

	// Get window canvas size as a reasonable default
	canvasSize := primaryWindow.Canvas().Size()

	// Calculate responsive initial window size based on screen size
	initialWidth := int(float64(canvasSize.Width) * constants.INITIAL_WIDTH_FACTOR)
	initialHeight := int(float64(canvasSize.Height) * constants.INITIAL_HEIGHT_FACTOR)

	// Limit window size to reasonable bounds
	maxWidth := float32(initialWidth)
	if maxWidth > constants.MAX_WINDOW_WIDTH {
		maxWidth = constants.MAX_WINDOW_WIDTH
	}

	// Create monitoring panel
	monitoringPanel := ui.NewMonitoringPanel(monitorSystem, widgetFactory, showDetailColumns)

	// Create a container with padding and proper spacing that will be updated
	content := container.New(layout.NewVBoxLayout(),
		container.NewPadded(themeButton),
		monitoringPanel.Container,
	)

	// Update the theme button callback to use the monitoring panel
	themeButton.OnTapped = func() {
		darkMode = !darkMode
		if darkMode {
			a.Settings().SetTheme(theme.DarkTheme())
			monitorSystem.UpdateTheme(true, constants.LightColors, constants.DarkColors)
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			monitorSystem.UpdateTheme(false, constants.LightColors, constants.DarkColors)
		}

		// Recreate the widget factory to get updated colors
		widgetFactory = ui.NewWidgetFactory(monitorSystem)

		// Recreate monitoring panel with new theme colors
		monitoringPanel = ui.NewMonitoringPanel(monitorSystem, widgetFactory, showDetailColumns)
		content.Objects[1] = monitoringPanel.Container

		// Force refresh to show theme changes
		content.Refresh()
	}

	// Periodically check for size changes
	sizeCheckTicker := time.NewTicker(time.Millisecond * constants.RESIZE_CHECK_INTERVAL)
	go func() {
		for range sizeCheckTicker.C {
			size := w.Canvas().Size()

			// Check if width or height changed significantly
			if (showDetailColumns && size.Width < constants.MIN_WIDTH) ||
				(!showDetailColumns && size.Width >= constants.MIN_WIDTH) ||
				size.Width != currentWidth || size.Height != currentHeight {

				// Update size tracking
				currentWidth = size.Width
				currentHeight = size.Height

				// Determine if we should show the detail column
				compact := size.Width < constants.MIN_WIDTH

				// Update layout state
				if showDetailColumns != !compact {
					showDetailColumns = !compact
					// Update the panel with the new layout setting
					monitoringPanel.SetShowDetail(showDetailColumns)
				}
			}
		}
	}()

	// Set content and register a resize listener
	w.SetContent(content)

	// Setup minimum constraints
	w.Resize(fyne.NewSize(maxWidth, float32(initialHeight)))

	// Start the system stats update ticker
	ticker := time.NewTicker(time.Millisecond * constants.STATS_UPDATE_INTERVAL)
	go func() {
		for range ticker.C {
			// Update system stats
			monitorSystem.UpdateSystemStats()

			// Check window size
			size := w.Canvas().Size()
			compact := size.Width < constants.MIN_WIDTH

			// Update the UI if layout has changed
			if showDetailColumns != !compact {
				showDetailColumns = !compact
				monitoringPanel.SetShowDetail(showDetailColumns)
			}

			// Always update widgets with new data
			monitoringPanel.Update()
		}
	}()

	w.ShowAndRun()
}
