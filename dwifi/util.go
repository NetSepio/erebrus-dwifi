package dwifi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
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
	// Check if this is the MAC of the current device
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if iface.HardwareAddr.String() == mac {
				// Use dmidecode to get the manufacturer for the current device
				out, err := exec.Command("sudo", "dmidecode", "-s", "system-manufacturer").Output()
				if err == nil {
					return strings.TrimSpace(string(out))
				}
				// If dmidecode fails, fall back to a generic description
				return "This Device"
			}
		}
	}

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

func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func getLocation() (string, error) {
	resp, err := http.Get("https://ipinfo.io/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var location struct {
		City    string `json:"city"`
		Region  string `json:"region"`
		Country string `json:"country"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s, %s, %s", location.City, location.Region, location.Country), nil
}

func PrintFancyBanner() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	banner := `
	 ______          _                      _____            _  ____  _
	|  ____|        | |                    |  __ \          (_)|  __)(_)
	| |__   _ __ ___| |__  _ __ _   _ ___  | |  | |_      _  _ | |_   _
	|  __| | '__/ _ \ '_ \| '__| | | / __| | |  | \ \ /\ / /| ||  _| | |
	| |____| | |  __/ |_) | |  | |_| \__ \ | |__| |\ V  V / | || |   | |
	|______|_|  \___|_.__/|_|   \__,_|___/ |_____/  \_/\_/  |_||_|   |_|
																  
    `
	lines := strings.Split(banner, "\n")
	colors := []*color.Color{
		color.New(color.FgRed).Add(color.Bold),
		color.New(color.FgYellow).Add(color.Bold),
		color.New(color.FgGreen).Add(color.Bold),
		color.New(color.FgCyan).Add(color.Bold),
		color.New(color.FgBlue).Add(color.Bold),
		color.New(color.FgMagenta).Add(color.Bold),
	}

	for i, line := range lines {
		colors[i%len(colors)].Println(line)
	}

	fmt.Println()
	color.New(color.FgHiWhite).Add(color.Bold, color.Underline).Println("Welcome to Erebrus DWiFi!")
	fmt.Println()
	fmt.Printf(" Wallet Address: %s\n", color.HiGreenString(os.Getenv("WALLET_ADDRESS")))
	fmt.Printf(" Chain : %s\n", color.HiBlueString(os.Getenv("CHAIN")))
	fmt.Println(" Price Per Minute : ", color.HiMagentaString(os.Getenv("PRICE_PER_MIN")))

}

func getOwnIP(iface *net.Interface) string {
	addrs, err := iface.Addrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
