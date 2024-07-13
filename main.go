package main

import (
	"log"
	"net"
	"sync"

	"github.com/NetSepio/erebrus-dwifi/dwifi"
	"github.com/fatih/color"
)

func main() {
	color.NoColor = false

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		go func(iface net.Interface) {
			defer wg.Done()
			if err := dwifi.Scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
	}
	wg.Wait()

	dwifi.PrintTotalConnectedTime()
}
