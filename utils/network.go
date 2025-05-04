package utils

import (
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// GetMaxNetworkSpeed returns the maximum network speed in MB/s for the current interface.
// Returns DefaultNetworkSpeed as a fallback if the actual speed cannot be determined.
func GetMaxNetworkSpeed() float64 {
	// Get the primary network interface
	ifaces, err := net.Interfaces()
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Find the first non-loopback interface
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			switch runtime.GOOS {
			case "darwin":
				return getMacNetworkSpeed(iface.Name)
			case "linux":
				return getLinuxNetworkSpeed(iface.Name)
			case "windows":
				return getWindowsNetworkSpeed(iface.Name)
			default:
				return DefaultNetworkSpeed
			}
		}
	}

	return DefaultNetworkSpeed
}

func getMacNetworkSpeed(iface string) float64 {
	// Use networksetup to get the hardware port for the interface
	cmd := exec.Command("networksetup", "-listallhardwareports")
	output, err := cmd.Output()
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Find the hardware port for our interface
	lines := strings.Split(string(output), "\n")
	var hardwarePort string
	for i, line := range lines {
		if strings.Contains(line, iface) && i > 0 {
			// The hardware port is in the previous line
			hardwarePort = strings.TrimSpace(strings.TrimPrefix(lines[i-1], "Hardware Port:"))
			break
		}
	}

	if hardwarePort == "" {
		return DefaultNetworkSpeed
	}

	// Get the current media type and speed
	cmd = exec.Command("networksetup", "-getmedia", hardwarePort)
	output, err = cmd.Output()
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Parse the output to find the speed
	lines = strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Speed:") {
			// Extract the speed value
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				continue
			}
			speedStr := strings.TrimSpace(parts[1])
			speedStr = strings.TrimSuffix(speedStr, "Mbps")
			speed, err := strconv.ParseFloat(strings.TrimSpace(speedStr), 64)
			if err != nil {
				continue
			}
			// Convert Mbps to MB/s
			return speed / MbpsToMBs
		}
	}

	return DefaultNetworkSpeed
}

func getLinuxNetworkSpeed(iface string) float64 {
	// Read the speed from /sys/class/net/<iface>/speed
	speedPath := "/sys/class/net/" + iface + "/speed"
	speedBytes, err := os.ReadFile(speedPath)
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Parse the speed value
	speedStr := strings.TrimSpace(string(speedBytes))
	speed, err := strconv.ParseFloat(speedStr, 64)
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Convert Mbps to MB/s
	return speed / MbpsToMBs
}

func getWindowsNetworkSpeed(iface string) float64 {
	// Use wmic to get the network adapter speed
	cmd := exec.Command("wmic", "nic", "where", "NetConnectionID='"+iface+"'", "get", "Speed")
	output, err := cmd.Output()
	if err != nil {
		return DefaultNetworkSpeed
	}

	// Parse the output to find the speed
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line == "Speed" {
			continue
		}
		speed, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}
		// Convert bps to MB/s
		return speed / BpsToMBs
	}

	return DefaultNetworkSpeed
}
