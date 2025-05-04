# Go System Monitor

A lightweight, cross-platform system resource monitoring application built with Go and Fyne UI toolkit.

![Screenshot of Go System Monitor](icon.png)

## Features

- Real-time monitoring of system resources:
  - CPU usage
  - Memory (RAM) usage
  - Disk usage and read/write speeds
  - Network upload/download speeds
- Responsive UI that adapts to window size
- Light and dark theme support
- Cross-platform compatibility (macOS, Linux, Windows)
- Historical data visualization with simple graphs

## Installation

### Prerequisites

- Go 1.23.3 or later

### From Source

Clone the repository and build the application:

```bash
git clone https://github.com/source-c/go-dummy-monitor.git
cd go-dummy-monitor
make build
```

The built executable will be available in the `build` directory.

## Usage

Simply run the executable to launch the application:

```bash
./build/go-dummy-monitor
```

### Interface

- Toggle between light and dark themes using the "Toggle Theme" button
- The UI automatically adapts to the window size
- Monitor CPU, RAM, disk, and network usage in real-time

## Development

### Project Structure

- `main.go`: Entry point and main application logic
- `constants/`: Application-wide constants and color definitions
- `ui/`: UI components, widgets, and monitoring system
  - `widgets/`: Custom widgets for displaying system metrics
- `utils/`: Utility functions for collecting system information

### Make Commands

The project includes a comprehensive Makefile for common development tasks:

```bash
make help               # Show available make commands
make build              # Build for current platform
make build-all          # Build for all supported platforms
make test               # Run all tests
make clean              # Clean build artifacts
make lint               # Run linter
make fmt                # Format code
make debug              # Run with debugger
```

### Dependencies

- [Fyne](https://fyne.io/) - Cross-platform GUI toolkit
- [gopsutil](https://github.com/shirou/gopsutil) - Process and system monitoring library

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.