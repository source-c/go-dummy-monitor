package utils

// Default network speed in MB/s (1 Gbps)
// Can be overridden at build time with -ldflags "-X go-dummy-monitor/utils.DefaultNetworkSpeed=250"
var DefaultNetworkSpeed float64 = 125 // 1 Gbps in MB/s

// Network speed conversion factors
const (
	MbpsToMBs = 8.0                 // Convert Mbps to MB/s
	BpsToMBs  = 8 * 1024.0 * 1024.0 // Convert bps to MB/s (bits to bytes * bytes to MB)
)
