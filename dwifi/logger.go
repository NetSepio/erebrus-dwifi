package dwifi

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func logDeviceInfo(event string, device *deviceInfo) {
	fmt.Println("----------------------------------------")
	fmt.Printf("%s : %s\n", color.HiGreenString(event), color.HiBlueString(time.Now().Format("2006-01-02 15:04:05")))

	fmt.Printf("  MAC: %s\n", color.HiMagentaString(device.macAddress.String()))
	fmt.Printf("  IP: %s\n", color.HiMagentaString(device.ipAddress.String()))
	fmt.Printf("  Manufacturer: %s\n", color.HiRedString(device.manufacturer))
	fmt.Printf("  Gateway: %s\n", color.HiYellowString(device.defaultGateway))
	fmt.Printf("  Interface: %s\n", color.HiYellowString(device.interfaceName))
	fmt.Printf("  Host SSID: %s\n", color.HiYellowString(device.HostSSID))
	fmt.Printf("  Total connected time: %v\n", color.HiCyanString(device.totalConnectedTime.String()))
}
