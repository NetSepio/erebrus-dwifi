package dwifi

import (
	"fmt"
	"os/exec"
	"strings"
)

func getDefaultGateway() string {
	out, err := exec.Command("ip", "route", "show", "default").Output()
	if err != nil {
		return "Unknown"
	}
	parts := strings.Fields(string(out))
	if len(parts) > 2 {
		return parts[2]
	}
	return "Unknown"
}

func getManufacturer(mac string) string {
	out, err := exec.Command("sudo", "arp-scan", "--localnet").Output()
	if err != nil {
		return "Unknown Manufacturer"
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, mac) {
			fields := strings.Fields(line)
			if len(fields) > 2 {
				return strings.Join(fields[2:], " ")
			}
		}
	}
	return "Unknown Manufacturer"
}

func getHostSSID(interfaceName string) string {
	out, err := exec.Command("nmcli", "dev", "show", interfaceName).Output()
	if err != nil {
		return "Unknown"
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "GENERAL.CONNECTION") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				return fields[1]
			}
		}
	}
	return "Unknown"
}

func getWiFiPassword() (string, string, error) {
	output, err := exec.Command("nmcli", "-t", "-f", "TYPE,NAME", "connection", "show", "--active").Output()
	if err != nil {
		return "", "", fmt.Errorf("error getting active connections: %v", err)
	}

	connections := strings.Split(strings.TrimSpace(string(output)), "\n")
	var wifiConnection string

	for _, conn := range connections {
		parts := strings.Split(conn, ":")
		if len(parts) == 2 && parts[0] == "802-11-wireless" {
			wifiConnection = parts[1]
			break
		}
	}

	if wifiConnection == "" {
		return "", "", fmt.Errorf("no active WiFi connection found")
	}

	output, err = exec.Command("nmcli", "-s", "-g", "802-11-wireless-security.psk", "connection", "show", wifiConnection).Output()
	if err != nil {
		return "", "", fmt.Errorf("error getting password for %s: %v", wifiConnection, err)
	}

	password := strings.TrimSpace(string(output))
	return wifiConnection, password, nil
}
