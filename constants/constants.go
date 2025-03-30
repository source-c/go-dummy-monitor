package constants

import (
	"image/color"
)

// Stubs
const (
	UNKNOWN_CPU = "Unknown CPU"
)

// UI sizing constants
const (
	// Timer intervals
	RESIZE_CHECK_INTERVAL = 500  // Milliseconds between resize checks
	STATS_UPDATE_INTERVAL = 1000 // Milliseconds between stats updates

	// Window sizing constants
	MIN_WIDTH         = 350 // Minimum width before hiding right column
	MIN_HEIGHT        = 400 // Minimum height before collapsing
	MIN_GRAPH_HEIGHT  = 100 // Minimum height of graph component
	GRAPH_WIDTH       = 300 // Width of graph component
	GRAPH_HEIGHT      = 180 // Height of graph component
	GRAPH_PADDING     = 40  // Padding around graph
	LABEL_HEIGHT      = 20  // Standard height for labels
	DEFAULT_PADDING   = 10  // Default padding between elements
	TEXT_PADDING      = 5   // Padding for text elements
	ELEMENT_SPACING   = 10  // Spacing between UI elements
	HEADING_TEXT_SIZE = 16  // Size for heading text
	NORMAL_TEXT_SIZE  = 14  // Size for normal text
	SMALL_TEXT_SIZE   = 12  // Size for small text
	STROKE_WIDTH      = 2   // Default stroke width

	// Window sizing factors
	INITIAL_WIDTH_FACTOR  = 0.6 // Percentage of screen width for initial window
	INITIAL_HEIGHT_FACTOR = 0.7 // Percentage of screen height for initial window
	MAX_WINDOW_WIDTH      = 800 // Maximum window width
	MIN_WINDOW_WIDTH      = 400 // Minimum window width
	MIN_WINDOW_HEIGHT     = 600 // Minimum window height
)

// Alpha values for colors
const (
	FULL_ALPHA             = 255
	TRANSLUCENT_ALPHA      = 180
	SEMI_TRANSLUCENT_ALPHA = 100
	TRANSPARENT_ALPHA      = 0
)

// Color definitions for light theme
const (
	// Light theme base colors
	LIGHT_GREEN_R, LIGHT_GREEN_G, LIGHT_GREEN_B    = 40, 180, 40
	LIGHT_PURPLE_R, LIGHT_PURPLE_G, LIGHT_PURPLE_B = 180, 40, 180
	LIGHT_CYAN_R, LIGHT_CYAN_G, LIGHT_CYAN_B       = 40, 180, 180
	LIGHT_YELLOW_R, LIGHT_YELLOW_G, LIGHT_YELLOW_B = 180, 180, 40
	LIGHT_BG_R, LIGHT_BG_G, LIGHT_BG_B             = 240, 240, 240
	LIGHT_TEXT_R, LIGHT_TEXT_G, LIGHT_TEXT_B       = 0, 0, 0
	LIGHT_GRID_R, LIGHT_GRID_G, LIGHT_GRID_B       = 200, 200, 200
)

// Color definitions for dark theme
const (
	// Dark theme base colors
	DARK_GREEN_R, DARK_GREEN_G, DARK_GREEN_B    = 40, 255, 40
	DARK_PURPLE_R, DARK_PURPLE_G, DARK_PURPLE_B = 255, 40, 255
	DARK_CYAN_R, DARK_CYAN_G, DARK_CYAN_B       = 40, 255, 255
	DARK_YELLOW_R, DARK_YELLOW_G, DARK_YELLOW_B = 255, 255, 40
	DARK_BG_R, DARK_BG_G, DARK_BG_B             = 0, 0, 0
	DARK_TEXT_R, DARK_TEXT_G, DARK_TEXT_B       = 255, 255, 255
	DARK_GRID_R, DARK_GRID_G, DARK_GRID_B       = 255, 255, 255
)

// ColorScheme defines all colors used in the application
type ColorScheme struct {
	CPU       color.Color
	RAM       color.Color
	DISK      color.Color
	NET       color.Color
	Text      color.Color
	SubText   color.Color
	BG        color.Color
	PanelBG   color.Color
	Grid      color.Color
	Highlight color.Color
	Shadow    color.Color
	Success   color.Color
	Warning   color.Color
	Error     color.Color
}

// Pre-defined color schemes
var (
	EmptyRectangle = color.RGBA{FULL_ALPHA, FULL_ALPHA, FULL_ALPHA, TRANSPARENT_ALPHA}
	LightColors    = ColorScheme{
		CPU:       color.RGBA{LIGHT_GREEN_R, LIGHT_GREEN_G, LIGHT_GREEN_B, FULL_ALPHA},
		RAM:       color.RGBA{LIGHT_PURPLE_R, LIGHT_PURPLE_G, LIGHT_PURPLE_B, FULL_ALPHA},
		DISK:      color.RGBA{LIGHT_CYAN_R, LIGHT_CYAN_G, LIGHT_CYAN_B, FULL_ALPHA},
		NET:       color.RGBA{LIGHT_YELLOW_R, LIGHT_YELLOW_G, LIGHT_YELLOW_B, FULL_ALPHA},
		Text:      color.RGBA{LIGHT_TEXT_R, LIGHT_TEXT_G, LIGHT_TEXT_B, FULL_ALPHA},
		SubText:   color.RGBA{100, 100, 100, FULL_ALPHA},
		BG:        color.RGBA{LIGHT_BG_R, LIGHT_BG_G, LIGHT_BG_B, FULL_ALPHA},
		PanelBG:   color.RGBA{230, 230, 230, FULL_ALPHA},
		Grid:      color.RGBA{LIGHT_GRID_R, LIGHT_GRID_G, LIGHT_GRID_B, TRANSPARENT_ALPHA},
		Highlight: color.RGBA{220, 220, 220, FULL_ALPHA},
		Shadow:    color.RGBA{180, 180, 180, FULL_ALPHA},
		Success:   color.RGBA{40, 200, 40, FULL_ALPHA},
		Warning:   color.RGBA{200, 180, 40, FULL_ALPHA},
		Error:     color.RGBA{200, 40, 40, FULL_ALPHA},
	}

	DarkColors = ColorScheme{
		CPU:       color.RGBA{DARK_GREEN_R, DARK_GREEN_G, DARK_GREEN_B, FULL_ALPHA},
		RAM:       color.RGBA{DARK_PURPLE_R, DARK_PURPLE_G, DARK_PURPLE_B, FULL_ALPHA},
		DISK:      color.RGBA{DARK_CYAN_R, DARK_CYAN_G, DARK_CYAN_B, FULL_ALPHA},
		NET:       color.RGBA{DARK_YELLOW_R, DARK_YELLOW_G, DARK_YELLOW_B, FULL_ALPHA},
		Text:      color.RGBA{DARK_TEXT_R, DARK_TEXT_G, DARK_TEXT_B, FULL_ALPHA},
		SubText:   color.RGBA{200, 200, 200, FULL_ALPHA},
		BG:        color.RGBA{DARK_BG_R, DARK_BG_G, DARK_BG_B, FULL_ALPHA},
		PanelBG:   color.RGBA{30, 30, 30, FULL_ALPHA},
		Grid:      color.RGBA{DARK_GRID_R, DARK_GRID_G, DARK_GRID_B, TRANSPARENT_ALPHA},
		Highlight: color.RGBA{60, 60, 60, FULL_ALPHA},
		Shadow:    color.RGBA{20, 20, 20, FULL_ALPHA},
		Success:   color.RGBA{40, 220, 40, FULL_ALPHA},
		Warning:   color.RGBA{220, 200, 40, FULL_ALPHA},
		Error:     color.RGBA{220, 40, 40, FULL_ALPHA},
	}
)
