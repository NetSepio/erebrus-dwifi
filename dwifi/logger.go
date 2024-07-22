package dwifi

import (
	"fmt"

	"github.com/fatih/color"
)

func logDeviceInfo(event string, device *deviceInfo) {

	fmt.Println("----------------------------------------")
	// fmt.Printf("%s : %s\n", color.HiGreenString(event), color.HiBlueString(time.Now().Format("2006-01-02 15:04:05")))

	fmt.Printf("  MAC: %s\n", color.HiMagentaString(device.MACAddress))
	fmt.Printf("  IP: %s\n", color.HiMagentaString(device.IPAddress))
	fmt.Printf("  Manufacturer: %s\n", color.HiRedString(device.Manufacturer))
	fmt.Printf("  Interface: %s\n", color.HiGreenString(device.InterfaceName))
	fmt.Printf("  Host SSID: %s\n", color.HiBlueString(device.HostSSID))
	fmt.Printf("  Total connected time: %s minutes\n", color.HiCyanString(fmt.Sprintf("%.2f", device.TotalConnectedTime.Minutes())))
	// if event == "Device disconnected" {
	// 	fmt.Printf("  Connected At: %s\n", color.HiCyanString(device.ConnectedAt.Format("2006-01-02 15:04:05")))
	// }
	fmt.Println("----------------------------------------")
}
