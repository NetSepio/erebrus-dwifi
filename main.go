package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/NetSepio/erebrus-dwifi/connections/database"
	"github.com/NetSepio/erebrus-dwifi/dwifi"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database.Connect()
}

func main() {
	color.NoColor = false

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	walletAddress := os.Getenv("WALLET_ADDRESS")
	if walletAddress == "" {
		log.Fatal("WALLET_ADDRESS environment variable is not set")
	}

	pricePerMinStr := os.Getenv("PRICE_PER_MIN")
	if pricePerMinStr == "" {
		log.Fatal("PRICE_PER_MIN environment variable is not set")
	}
	pricePerMin, err := strconv.ParseFloat(pricePerMinStr, 64)
	if err != nil {
		log.Fatalf("Invalid PRICE_PER_MIN value: %v", err)
	}

	db, err := dwifi.InitDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// dwifi.InitDevicesDatabase(db)
	dwifi.InitDevicesDatabase(db, walletAddress, pricePerMin)

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
