package ui

import (
	"fmt"
	"image/color"

	"go-dummy-monitor/ui/widgets"
)

// CPUDataProvider provides CPU-specific data for graphing
type CPUDataProvider struct {
	System MonitoringSystem
}

func (c *CPUDataProvider) GetData() []float64 {
	return c.System.GetCPUData()
}

func (c *CPUDataProvider) GetMaxValue() float64 {
	return 100.0 // CPU percentage is always 0-100
}

func (c *CPUDataProvider) GetCurrentValue() float64 {
	return c.System.GetCPUUsage()
}

func (c *CPUDataProvider) GetTitle() string {
	return "CPU"
}

func (c *CPUDataProvider) GetColor() color.Color {
	return c.System.GetColorScheme().CPU
}

// RAMDataProvider provides RAM-specific data for graphing
type RAMDataProvider struct {
	System MonitoringSystem
}

func (r *RAMDataProvider) GetData() []float64 {
	return r.System.GetRAMData()
}

func (r *RAMDataProvider) GetMaxValue() float64 {
	return 100.0 // RAM percentage is always 0-100
}

func (r *RAMDataProvider) GetCurrentValue() float64 {
	return r.System.GetRAMUsage()
}

func (r *RAMDataProvider) GetTitle() string {
	return "RAM"
}

func (r *RAMDataProvider) GetColor() color.Color {
	return r.System.GetColorScheme().RAM
}

// DiskDataProvider provides Disk-specific data for graphing
type DiskDataProvider struct {
	System MonitoringSystem
}

func (d *DiskDataProvider) GetReadData() []float64 {
	return d.System.GetDiskReadData()
}

func (d *DiskDataProvider) GetWriteData() []float64 {
	return d.System.GetDiskWriteData()
}

func (d *DiskDataProvider) GetMaxValue() float64 {
	return 100.0 // For graph scale
}

func (d *DiskDataProvider) GetCurrentReadValue() float64 {
	data := d.System.GetDiskReadData()
	if len(data) > 0 {
		return data[len(data)-1]
	}
	return 0
}

func (d *DiskDataProvider) GetCurrentWriteValue() float64 {
	data := d.System.GetDiskWriteData()
	if len(data) > 0 {
		return data[len(data)-1]
	}
	return 0
}

func (d *DiskDataProvider) GetTitle() string {
	return "Disk"
}

func (d *DiskDataProvider) GetColor() color.Color {
	return d.System.GetColorScheme().DISK
}

// NetworkDataProvider provides Network-specific data for graphing
type NetworkDataProvider struct {
	System MonitoringSystem
}

func (n *NetworkDataProvider) GetReadData() []float64 {
	return n.System.GetNetworkReadData()
}

func (n *NetworkDataProvider) GetWriteData() []float64 {
	return n.System.GetNetworkWriteData()
}

func (n *NetworkDataProvider) GetMaxValue() float64 {
	return n.System.GetMaxNetworkSpeed()
}

func (n *NetworkDataProvider) GetCurrentReadValue() float64 {
	data := n.System.GetNetworkReadData()
	if len(data) > 0 {
		return data[len(data)-1]
	}
	return 0
}

func (n *NetworkDataProvider) GetCurrentWriteValue() float64 {
	data := n.System.GetNetworkWriteData()
	if len(data) > 0 {
		return data[len(data)-1]
	}
	return 0
}

func (n *NetworkDataProvider) GetTitle() string {
	return "Net"
}

func (n *NetworkDataProvider) GetColor() color.Color {
	return n.System.GetColorScheme().NET
}

// GetDiskInfoProvider returns disk info functions
func GetDiskInfoProvider(system MonitoringSystem) []widgets.InfoRow {
	return []widgets.InfoRow{
		{
			Label: "Read",
			GetValue: func() string {
				data := system.GetDiskReadData()
				if len(data) > 0 {
					return fmt.Sprintf("%.2f MB/s", data[len(data)-1])
				}
				return "0.00 MB/s"
			},
		},
		{
			Label: "Write",
			GetValue: func() string {
				data := system.GetDiskWriteData()
				if len(data) > 0 {
					return fmt.Sprintf("%.2f MB/s", data[len(data)-1])
				}
				return "0.00 MB/s"
			},
		},
		{
			Label: "Usage",
			GetValue: func() string {
				return fmt.Sprintf("%.1f%%", system.GetDiskUsage())
			},
		},
	}
}

// GetRAMInfoProvider returns RAM info functions
func GetRAMInfoProvider(system MonitoringSystem) []widgets.InfoRow {
	return []widgets.InfoRow{
		{
			Label: "Total",
			GetValue: func() string {
				v := system.GetVirtualMemory()
				return fmt.Sprintf("%.1f GB", float64(v.Total)/(1024*1024*1024))
			},
		},
		{
			Label: "Used",
			GetValue: func() string {
				v := system.GetVirtualMemory()
				return fmt.Sprintf("%.1f GB (%.1f%%)",
					float64(v.Used)/(1024*1024*1024), system.GetRAMUsage())
			},
		},
		{
			Label: "Free",
			GetValue: func() string {
				v := system.GetVirtualMemory()
				return fmt.Sprintf("%.1f GB", float64(v.Free)/(1024*1024*1024))
			},
		},
	}
}

// GetCPUInfoProvider returns CPU info functions
func GetCPUInfoProvider(system MonitoringSystem) []widgets.InfoRow {
	return []widgets.InfoRow{
		{
			Label: "Model",
			GetValue: func() string {
				return system.GetCPUModelName()
			},
		},
		{
			Label: "Usage",
			GetValue: func() string {
				return fmt.Sprintf("%.2f%%", system.GetCPUUsage())
			},
		},
		{
			Label: "Cores",
			GetValue: func() string {
				return fmt.Sprintf("%d physical / %d logical",
					system.GetPhysicalCPUCount(), system.GetLogicalCPUCount())
			},
		},
	}
}

// GetNetworkInfoProvider returns network info functions
func GetNetworkInfoProvider(system MonitoringSystem) []widgets.InfoRow {
	return []widgets.InfoRow{
		{
			Label: "Interface",
			GetValue: func() string {
				return system.GetActiveNetInterfaceName()
			},
		},
		{
			Label: "Max Speed",
			GetValue: func() string {
				return fmt.Sprintf("%.1f MB/s", system.GetMaxNetworkSpeed())
			},
		},
	}
}
