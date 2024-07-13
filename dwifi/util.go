package dwifi

import (
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
