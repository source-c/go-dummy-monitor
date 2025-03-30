package ui

import (
	"fmt"
	"go-dummy-monitor/constants"
	"image/color"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	psnet "github.com/shirou/gopsutil/net"
)

// MonitorSystem implements the MonitoringSystem interface
type MonitorSystem struct {
	cpuData          []float64
	cpuUsage         float64
	ramData          []float64
	ramUsage         float64
	diskReadData     []float64
	diskWriteData    []float64
	diskUsage        float64
	networkReadData  []float64
	networkWriteData []float64
	prevDiskStats    map[string]disk.IOCountersStat
	maxNetworkSpeed  float64
	colorScheme      ColorScheme
	emptyRectangle   color.Color
	darkMode         bool
}

// NewMonitorSystem creates a new MonitorSystem
func NewMonitorSystem(
	maxNetworkSpeed float64,
	darkMode bool,
	lightColorScheme ColorScheme,
	darkColorScheme ColorScheme,
	emptyRectangle color.Color,
	dataPoints int,
) *MonitorSystem {
	system := &MonitorSystem{
		cpuData:          make([]float64, dataPoints),
		ramData:          make([]float64, dataPoints),
		diskReadData:     make([]float64, dataPoints),
		diskWriteData:    make([]float64, dataPoints),
		networkReadData:  make([]float64, dataPoints),
		networkWriteData: make([]float64, dataPoints),
		maxNetworkSpeed:  maxNetworkSpeed,
		emptyRectangle:   emptyRectangle,
		darkMode:         darkMode,
	}

	// Set initial color scheme based on mode
	if darkMode {
		system.colorScheme = darkColorScheme
	} else {
		system.colorScheme = lightColorScheme
	}

	// Initialize previous disk stats
	var err error
	system.prevDiskStats, err = disk.IOCounters()
	if err != nil {
		system.prevDiskStats = make(map[string]disk.IOCountersStat)
	}

	return system
}

// UpdateTheme updates the color scheme based on dark mode
func (s *MonitorSystem) UpdateTheme(darkMode bool, lightColorScheme ColorScheme, darkColorScheme ColorScheme) {
	s.darkMode = darkMode
	if darkMode {
		s.colorScheme = darkColorScheme
	} else {
		s.colorScheme = lightColorScheme
	}
}

// UpdateSystemStats collects and updates system stats
func (s *MonitorSystem) UpdateSystemStats() {
	cpuUsage, _ := cpu.Percent(0, false)
	memStats, _ := mem.VirtualMemory()
	diskStats, _ := disk.Usage("/")
	initialStats, _ := psnet.IOCounters(true)

	time.Sleep(time.Millisecond * 500)

	netStats, _ := psnet.IOCounters(true)

	activeNetInterfaceName := s.GetActiveNetInterfaceName()

	// Initialize with some values to make the graph visible
	netReadSpeed := 0.0
	netWriteSpeed := 0.0

	// Get the actual network usage if available
	if activeNetInterfaceName != "" {
		for i := 0; i < len(netStats); i++ {
			if strings.Contains(strings.ToLower(netStats[i].Name), strings.ToLower(activeNetInterfaceName)) {
				bytesReceived := netStats[i].BytesRecv - initialStats[i].BytesRecv
				mbReceived := float64(bytesReceived) / (1024 * 1024)
				netReadSpeed = mbReceived * 2 // Convert to per second (500ms interval)

				bytesSent := netStats[i].BytesSent - initialStats[i].BytesSent
				mbSent := float64(bytesSent) / (1024 * 1024)
				netWriteSpeed = mbSent * 2 // Convert to per second (500ms interval)

				// Debug print
				//fmt.Printf("Net speed: D=%.2f MB/s, U=%.2f MB/s\n", netReadSpeed, netWriteSpeed)
				break
			}
		}
	}

	ioStats, _ := disk.IOCounters()
	// No debug output needed in production

	var readSpeed, writeSpeed float64
	// Initialize ioStats if prev is empty
	if len(s.prevDiskStats) == 0 {
		s.prevDiskStats = ioStats
	}

	// Initialize with some values to make the graph visible
	readSpeed = 0.0
	writeSpeed = 0.0

	// Get real disk activity if available
	for diskName, stats := range ioStats {
		prevStats, exists := s.prevDiskStats[diskName]
		if exists {
			// The values are in bytes per interval (500ms), convert to MB/s
			bytesRead := float64(stats.ReadBytes - prevStats.ReadBytes)
			bytesWritten := float64(stats.WriteBytes - prevStats.WriteBytes)

			// Convert bytes to MB and adjust for time interval (500ms = 0.5s)
			readSpeed = (bytesRead / 1024 / 1024) * 2     // Convert to MB/s
			writeSpeed = (bytesWritten / 1024 / 1024) * 2 // Convert to MB/s

			// Avoid stopping at the first disk with no activity
			if readSpeed > 0 || writeSpeed > 0 {
				break
			}
		}
	}

	// Update system data
	s.cpuUsage = cpuUsage[0]
	s.ramUsage = memStats.UsedPercent
	s.diskUsage = diskStats.UsedPercent
	s.prevDiskStats = ioStats

	// Update historical data
	s.cpuData = append(s.cpuData[1:], s.cpuUsage)
	s.ramData = append(s.ramData[1:], s.ramUsage)
	// readSpeed and writeSpeed are already in MB/s
	s.diskReadData = append(s.diskReadData[1:], readSpeed)
	s.diskWriteData = append(s.diskWriteData[1:], writeSpeed)
	s.networkReadData = append(s.networkReadData[1:], netReadSpeed)
	s.networkWriteData = append(s.networkWriteData[1:], netWriteSpeed)

	// Values updated successfully
}

// IsDarkMode returns whether the system is in dark mode
func (s *MonitorSystem) IsDarkMode() bool {
	return s.darkMode
}

// GetColorScheme returns the current color scheme
func (s *MonitorSystem) GetColorScheme() ColorScheme {
	return s.colorScheme
}

// GetEmptyRectangle returns the empty rectangle color
func (s *MonitorSystem) GetEmptyRectangle() color.Color {
	return s.emptyRectangle
}

// GetCPUData returns the CPU usage data
func (s *MonitorSystem) GetCPUData() []float64 {
	return s.cpuData
}

// GetCPUUsage returns the current CPU usage
func (s *MonitorSystem) GetCPUUsage() float64 {
	return s.cpuUsage
}

// GetRAMData returns the RAM usage data
func (s *MonitorSystem) GetRAMData() []float64 {
	return s.ramData
}

// GetRAMUsage returns the current RAM usage
func (s *MonitorSystem) GetRAMUsage() float64 {
	return s.ramUsage
}

// GetDiskReadData returns the disk read data
func (s *MonitorSystem) GetDiskReadData() []float64 {
	return s.diskReadData
}

// GetDiskWriteData returns the disk write data
func (s *MonitorSystem) GetDiskWriteData() []float64 {
	return s.diskWriteData
}

// GetDiskUsage returns the current disk usage
func (s *MonitorSystem) GetDiskUsage() float64 {
	return s.diskUsage
}

// GetNetworkReadData returns the network read data
func (s *MonitorSystem) GetNetworkReadData() []float64 {
	return s.networkReadData
}

// GetNetworkWriteData returns the network write data
func (s *MonitorSystem) GetNetworkWriteData() []float64 {
	return s.networkWriteData
}

// GetMaxNetworkSpeed returns the max network speed
func (s *MonitorSystem) GetMaxNetworkSpeed() float64 {
	return s.maxNetworkSpeed
}

// GetActiveNetInterfaceName returns the active network interface name
func (s *MonitorSystem) GetActiveNetInterfaceName() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Don't need debug output in production

	for _, iface := range interfaces {
		// Skip interfaces that are down or loopback interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Check if the interface has an IP address
		addrs, err := iface.Addrs()
		if err == nil && len(addrs) > 0 {
			// Found a usable interface

			// For macOS, we need to handle different interface names
			if runtime.GOOS == "darwin" {
				// TODO: On macOS, we look for en0 (WiFi) or en1 (Ethernet)
				if iface.Name == "en0" || iface.Name == "en1" {
					return iface.Name
				}
			} else {
				// Always accept these common types
				if strings.Contains(strings.ToLower(iface.Name), "wi-fi") ||
					strings.Contains(strings.ToLower(iface.Name), "wlan") ||
					strings.Contains(strings.ToLower(iface.Name), "eth") ||
					strings.Contains(strings.ToLower(iface.Name), "en") ||
					strings.Contains(strings.ToLower(iface.Name), "wlp") {
					return iface.Name
				}
			}

			// Return first valid interface if no specific match
			return iface.Name
		}
	}
	return ""
}

// getCPUInfoDarwin gets CPU info on Darwin systems
func getCPUInfoDarwin() (string, int, int) {
	// Get CPU model name
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	output, err := cmd.Output()
	modelName := "Apple Silicon M1" // FIXME: this is a subject to change in future to handle M{1,2,3,4}* silicons!!
	if err == nil {
		modelName = strings.TrimSpace(string(output))
	}

	// Get CPU cores
	cmd = exec.Command("sysctl", "-n", "hw.ncpu")
	output, err = cmd.Output()
	logicalCores := 8
	physicalCores := 4
	if err == nil {
		count := strings.TrimSpace(string(output))
		_, err := fmt.Sscanf(count, "%d", &logicalCores)
		if err != nil {
			return "", 0, 0
		}
		physicalCores = logicalCores / 2
	}

	return modelName, logicalCores, physicalCores
}

// GetCPUModelName returns the CPU model name
func (s *MonitorSystem) GetCPUModelName() string {
	if runtime.GOOS == "darwin" {
		modelName, _, _ := getCPUInfoDarwin()
		return modelName
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return constants.UNKNOWN_CPU
	}

	if len(cpuInfo) > 0 {
		return cpuInfo[0].ModelName
	}

	return constants.UNKNOWN_CPU
}

// GetPhysicalCPUCount returns the number of physical CPU cores
func (s *MonitorSystem) GetPhysicalCPUCount() int {
	if runtime.GOOS == "darwin" {
		_, _, physicalCores := getCPUInfoDarwin()
		return physicalCores
	}

	count, err := cpu.Counts(false)
	if err != nil {
		return 1
	}

	return count
}

// GetLogicalCPUCount returns the number of logical CPU cores
func (s *MonitorSystem) GetLogicalCPUCount() int {
	if runtime.GOOS == "darwin" {
		_, logicalCores, _ := getCPUInfoDarwin()
		return logicalCores
	}

	count, err := cpu.Counts(true)
	if err != nil {
		return 1
	}

	return count
}

// GetVirtualMemory returns the virtual memory stats
func (s *MonitorSystem) GetVirtualMemory() *mem.VirtualMemoryStat {
	v, err := mem.VirtualMemory()
	if err != nil {
		return &mem.VirtualMemoryStat{}
	}

	return v
}

// GetPartitionInfo returns disk partition information
func (s *MonitorSystem) GetPartitionInfo() []disk.PartitionStat {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return []disk.PartitionStat{}
	}

	return partitions
}
