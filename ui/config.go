package ui

import (
	"image/color"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"go-dummy-monitor/constants"
)

// Constants for measurements
const (
	MaxNetworkSpeedDefault = 125 // Max network speed in MB/s (125 MB/s is the 1000Mbit/s or 1Gbps)
)

// Constants for UI sizing
const (
	GraphWidth      = float32(constants.GRAPH_WIDTH)
	GraphHeight     = float32(constants.GRAPH_HEIGHT)
	GraphPadding    = float32(constants.GRAPH_PADDING)
	LabelHeight     = float32(constants.LABEL_HEIGHT)
	ElementSpacing  = float32(constants.ELEMENT_SPACING)
	DefaultPadding  = float32(constants.DEFAULT_PADDING)
	TextPadding     = float32(constants.TEXT_PADDING)
	HeadingTextSize = float32(constants.HEADING_TEXT_SIZE)
	NormalTextSize  = float32(constants.NORMAL_TEXT_SIZE)
	SmallTextSize   = float32(constants.SMALL_TEXT_SIZE)
	StrokeWidth     = float32(constants.STROKE_WIDTH)

	FullAlpha            = constants.FULL_ALPHA
	TranslucentAlpha     = constants.TRANSLUCENT_ALPHA
	SemiTranslucentAlpha = constants.SEMI_TRANSLUCENT_ALPHA
	TransparentAlpha     = constants.TRANSPARENT_ALPHA
)

// ColorScheme is imported from constants package
type ColorScheme = constants.ColorScheme

// ComponentType defines different system component types
type ComponentType int

const (
	CPUComponent ComponentType = iota
	RAMComponent
	DiskComponent
	NetworkComponent
)

// SystemDataProvider is an interface for components that provide system data
type SystemDataProvider interface {
	GetCPUData() []float64
	GetCPUUsage() float64
	GetRAMData() []float64
	GetRAMUsage() float64
	GetDiskReadData() []float64
	GetDiskWriteData() []float64
	GetDiskUsage() float64
	GetNetworkReadData() []float64
	GetNetworkWriteData() []float64
	GetActiveNetInterfaceName() string
	GetPartitionInfo() []disk.PartitionStat
	GetVirtualMemory() *mem.VirtualMemoryStat
	GetPhysicalCPUCount() int
	GetLogicalCPUCount() int
	GetCPUModelName() string
	GetMaxNetworkSpeed() float64
}

// Theme defines the interface for theme-related functionality
type Theme interface {
	IsDarkMode() bool
	GetColorScheme() ColorScheme
	GetEmptyRectangle() color.Color
}

// MonitoringSystem defines the interface for the entire monitoring system
type MonitoringSystem interface {
	SystemDataProvider
	Theme
}
