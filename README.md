# erebrus-dwifi

erebrus-dwifi is a Go program that tracks devices connected to your network using ARP scanning and packet capture.

## Overview

This program scans the local network to detect devices based on ARP responses. It logs connection and disconnection events, tracks total connected time, and retrieves device information such as MAC address, IP address, manufacturer, Host SSID, Gateway IP and network interface details.

## Features

- ARP scanning to detect connected devices
- Real-time logging of device connections and disconnections
- Calculation of total connected time per device
- Retrieval of device manufacturer and network interface details

## Installation

### Prerequisites

- Go programming language installed (version 1.16+ recommended)
- `arp-scan` and `nmcli` utilities installed for device information retrieval

### Building from Source

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/NetSepio/erebrus-dwifi.git
cd erebrus-dwifi

# Linux
GOOS=linux GOARCH=amd64 go build -o erebrus-dwifi-linux main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o erebrus-dwifi-macos main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o erebrus-dwifi.exe main.go

# Run the Binary according to your platform
./erebrus-dwifi-linux    # Linux
./erebrus-dwifi-macos    # macOS
erebrus-dwifi.exe       # Windows
