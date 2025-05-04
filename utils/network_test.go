package utils

import (
	"math"
	"testing"
)

func TestGetMaxNetworkSpeed(t *testing.T) {
	// Test that we get a valid network speed
	speed := GetMaxNetworkSpeed()
	if speed <= 0 {
		t.Errorf("Expected positive network speed, got %f", speed)
	}

	// Test that we get at least the default speed
	if speed < DefaultNetworkSpeed {
		t.Errorf("Expected speed >= %f, got %f", DefaultNetworkSpeed, speed)
	}
}

func TestNetworkSpeedConversion(t *testing.T) {
	// Test Mbps to MB/s conversion
	mbps := 1000.0 // 1 Gbps
	mbs := mbps / MbpsToMBs
	if mbs != 125.0 {
		t.Errorf("Expected 125 MB/s for 1 Gbps, got %f", mbs)
	}

	// Test bps to MB/s conversion
	bps := 1000000000.0 // 1 Gbps in bps
	mbs = bps / BpsToMBs
	expected := 119.209
	actual := math.Round(mbs*1000) / 1000 // Round to 3 decimal places
	if actual != expected {
		t.Errorf("Expected %.3f MB/s for 1 Gbps in bps, got %.3f", expected, actual)
	}
}
