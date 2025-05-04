package ui

import (
	"testing"

	"go-dummy-monitor/constants"
)

func TestNewMonitorSystem(t *testing.T) {
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Test initial state
	if system.maxNetworkSpeed != 100.0 {
		t.Errorf("Expected maxNetworkSpeed to be 100.0, got %f", system.maxNetworkSpeed)
	}
	if system.darkMode != false {
		t.Error("Expected darkMode to be false")
	}
	if len(system.cpuData) != 60 {
		t.Errorf("Expected cpuData length to be 60, got %d", len(system.cpuData))
	}
	if len(system.ramData) != 60 {
		t.Errorf("Expected ramData length to be 60, got %d", len(system.ramData))
	}
	if len(system.diskReadData) != 60 {
		t.Errorf("Expected diskReadData length to be 60, got %d", len(system.diskReadData))
	}
	if len(system.diskWriteData) != 60 {
		t.Errorf("Expected diskWriteData length to be 60, got %d", len(system.diskWriteData))
	}
	if len(system.networkReadData) != 60 {
		t.Errorf("Expected networkReadData length to be 60, got %d", len(system.networkReadData))
	}
	if len(system.networkWriteData) != 60 {
		t.Errorf("Expected networkWriteData length to be 60, got %d", len(system.networkWriteData))
	}
}

func TestUpdateTheme(t *testing.T) {
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Test updating to dark mode
	system.UpdateTheme(true, constants.LightColors, constants.DarkColors)
	if !system.darkMode {
		t.Error("Expected darkMode to be true after update")
	}
	if system.colorScheme != constants.DarkColors {
		t.Error("Expected colorScheme to be DarkColors after update")
	}

	// Test updating to light mode
	system.UpdateTheme(false, constants.LightColors, constants.DarkColors)
	if system.darkMode {
		t.Error("Expected darkMode to be false after update")
	}
	if system.colorScheme != constants.LightColors {
		t.Error("Expected colorScheme to be LightColors after update")
	}
}

func TestGetColorScheme(t *testing.T) {
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Test initial color scheme
	if system.GetColorScheme() != constants.LightColors {
		t.Error("Expected initial color scheme to be LightColors")
	}

	// Test color scheme after theme update
	system.UpdateTheme(true, constants.LightColors, constants.DarkColors)
	if system.GetColorScheme() != constants.DarkColors {
		t.Error("Expected color scheme to be DarkColors after theme update")
	}
}

func TestGetEmptyRectangle(t *testing.T) {
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Test empty rectangle color
	if system.GetEmptyRectangle() != constants.EmptyRectangle {
		t.Error("Expected empty rectangle color to match constants.EmptyRectangle")
	}
}

func TestGetMaxNetworkSpeed(t *testing.T) {
	system := NewMonitorSystem(
		100.0, // maxNetworkSpeed
		false, // darkMode
		constants.LightColors,
		constants.DarkColors,
		constants.EmptyRectangle,
		60, // dataPoints
	)

	// Test max network speed
	if system.GetMaxNetworkSpeed() != 100.0 {
		t.Errorf("Expected max network speed to be 100.0, got %f", system.GetMaxNetworkSpeed())
	}
}
