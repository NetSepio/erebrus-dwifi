package contract

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/joho/godotenv"

	contract "github.com/NetSepio/erebrus-dwifi/contractcall/contract"
)

func GeneratePeaqDID(length int) (string, error) {
	if length <= 0 {
		length = 51
	}
	const validChars = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}
		result[i] = validChars[randomIndex.Int64()]
	}
	return fmt.Sprintf("did:pq:%s", string(result)), nil
}

func ErebrusRegistry() (string, *big.Int, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_NODE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	erebrusRegistry, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate the contract: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}

	chainID, ok := new(big.Int).SetString(os.Getenv("CHAIN_ID"), 10)
	if !ok {
		log.Fatal("Failed to parse CHAIN_ID")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	peaqDid, err := GeneratePeaqDID(51)
	fmt.Println(color.HiGreenString(fmt.Sprintf(" Peaq DID: %s", peaqDid)))
	if err != nil {
		log.Fatalf("Failed to generate PEAQ DID: %v", err)
	}

	deviceId := os.Getenv("DEVICE_ID")
	ssid := os.Getenv("SSID")
	location := os.Getenv("LOCATION")

	reader := bufio.NewReader(os.Stdin)
	time.Sleep(1 * time.Second)
	fmt.Print(" Enter price per minute in Agung: ")
	pricePerMinuteStr, _ := reader.ReadString('\n')
	pricePerMinuteStr = strings.TrimSpace(pricePerMinuteStr)

	pricePerMinuteFloat, err := strconv.ParseFloat(pricePerMinuteStr, 64)
	if err != nil {
		log.Fatal("Failed to parse price per minute as float:", err)
	}

	pricePerMinuteWei := new(big.Int).Mul(big.NewInt(int64(pricePerMinuteFloat*math.Pow10(18))), big.NewInt(1))

	tx, err := erebrusRegistry.RegisterWifiNode(auth, deviceId, peaqDid, ssid, location, pricePerMinuteWei)
	if err != nil {
		log.Fatalf("Failed to register WiFi node: %v", err)
	}

	// fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

	// fmt.Printf("Transaction mined in block %d\n", receipt.BlockNumber)
	return tx.Hash().Hex(), receipt.BlockNumber, nil
}
