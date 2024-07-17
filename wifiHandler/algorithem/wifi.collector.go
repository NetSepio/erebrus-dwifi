package wifiAlgorithem

import "github.com/NetSepio/erebrus-dwifi/connections/database/wifiDbOperations"

//craete function

func WifiCollector(data wifiDbOperations.DeviceInfo) {
	wifiDbOperations.CreateDevice(&data)
}
