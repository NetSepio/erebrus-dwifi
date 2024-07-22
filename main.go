package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"sync"

	"github.com/NetSepio/erebrus-dwifi/dwifi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/joho/godotenv"

	node "github.com/NetSepio/erebrus-dwifi/node"
)

func main() {
	color.NoColor = false

	dwifi.PrintFancyBanner()
	pricePerMinStr := os.Getenv("PRICE_PER_MIN")
	if pricePerMinStr == "" {
		log.Fatal("PRICE_PER_MIN environment variable is not set")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_NODE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(9990)) // Use the correct chain ID for your network
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := node.NewNode(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	user := common.HexToAddress(os.Getenv("WALLET_ADDRESS"))
	deviceId := os.Getenv("DEVICE_ID")
	peaqDid := common.HexToAddress(os.Getenv("PEAQ_DID"))
	ssid := os.Getenv("SSID")
	location := os.Getenv("LOCATION")
	pricePerMin, ok := new(big.Int).SetString(os.Getenv("PRICE_PER_MIN"), 10)
	if !ok {
		log.Fatal("Failed to parse PRICE_PER_MIN")
	}

	tx, err := instance.DelegateRegisterWifiNode(auth, user, deviceId, peaqDid, ssid, location, pricePerMin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(color.HiYellowString(" tx sent: %s\n", tx.Hash().Hex()))

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(color.HiGreenString(" tx mined: %s\n", receipt.TxHash.Hex()))
	fmt.Println("____________________________________________________________________________________")
	fmt.Println(color.HiBlueString(" Dwifi Registered SUccessfully!!!"))
	fmt.Println("____________________________________________________________________________________")

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	walletAddress := os.Getenv("WALLET_ADDRESS")
	if walletAddress == "" {
		log.Fatal("WALLET_ADDRESS environment variable is not set")
	}

	db, err := dwifi.InitDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dwifi.InitDevicesDatabase(db, walletAddress, os.Getenv("PRICE_PER_MIN"))

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
				// log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
	}
	wg.Wait()

	dwifi.PrintTotalConnectedTime()
}
