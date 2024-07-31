# Erebrus-dwifi

erebrus-dwifi is a Go program that tracks devices connected to your network using ARP scanning and packet capture, built on the PEAQ network for Decentralized Physical Infrastructure Networks (DePIN).

## Overview

This program scans the local network to detect devices based on ARP responses. It logs connection and disconnection events, tracks total connected time, and retrieves device information such as MAC address, IP address, manufacturer, Host SSID, Gateway IP and network interface details. Leveraging the PEAQ network, erebrus-dwifi contributes to the decentralized management of physical network infrastructure.

## Features

- ARP scanning to detect connected devices
- Real-time logging of device connections and disconnections
- Calculation of total connected time per device
- Retrieval of device manufacturer and network interface details
- Integration with PEAQ network for decentralized infrastructure management

## PEAQ Network and DePIN

erebrus-dwifi is built on the PEAQ network, a blockchain platform designed for Decentralized Physical Infrastructure Networks (DePIN). This integration allows for:

- Decentralized management of network data
- Enhanced security and privacy in device tracking
- Potential for tokenized incentives for network participants
- Contribution to a broader ecosystem of decentralized infrastructure solutions

## Installation

### Prerequisites

- Go programming language installed (version 1.16+ recommended)
- `arp-scan` and `nmcli` utilities installed for device information retrieval
- PEAQ network node or access point (refer to PEAQ documentation for setup)

### Building from Source

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/NetSepio/erebrus-dwifi.git
cd erebrus-dwifi

# Linux
GOOS=linux GOARCH=amd64 go build -o erebrus-dwifi-linux main.go

# Run the Binary according to your platform
./erebrus-dwifi-linux    # Linux