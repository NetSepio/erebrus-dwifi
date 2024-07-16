package main

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/NetSepio/erebrus-dwifi/dwifi"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	color.NoColor = false

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	db, err := dwifi.InitDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dwifi.InitDevicesDatabase(db)

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
