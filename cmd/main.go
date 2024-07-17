package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"abi/api"
	"abi/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress("<SMART_CONTRACT>")

	oracle, err := api.NewApi(contractAddress, client)
	if err != nil {
		log.Fatalf("Error initializing contract: %v", err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	chainID := big.NewInt(31337)
	auth, err := createAuth(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	url := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.12/dados?formato=json"
	data, err := json.FetchData(url)
	if err != nil {
		log.Fatalf("Failed to fetch data from API: %v", err)
	}

	for _, entry := range data {

		layout := "02/01/2006"
		date, err := time.Parse(layout, entry.Data)
		if err != nil {
			log.Printf("Failed to parse date %s: %v", entry.Data, err)
			continue
		}

		timestamp := big.NewInt(date.Unix())

		value, ok := new(big.Float).SetString(entry.Valor)
		if !ok {
			log.Printf("Invalid value format for %s", entry.Valor)
			continue
		}

		scale := new(big.Float).SetFloat64(1e6)
		value.Mul(value, scale)

		intValue := new(big.Int)
		value.Int(intValue)

		updatedAt := big.NewInt(time.Now().Unix())

		fmt.Printf("Timestamp for date %s: %d\n", entry.Data, timestamp.Int64())
		fmt.Printf("Value for date %s: %s\n", entry.Data, value.String())

		tx, err := oracle.SaveIndicator(auth, timestamp, intValue, updatedAt, 0)
		if err != nil {
			log.Printf("Failed to save indicator for date %s: %v", entry.Data, err)
			continue
		}

		receipt, err := waitForReceipt(client, tx.Hash())
		if err != nil {
			log.Printf("Failed to get transaction receipt for date %s: %v", entry.Data, err)
			continue
		}

		err = saveCSV(timestamp.String(), receipt.GasUsed)
		if err != nil {
			log.Printf("Failed to save transaction details to CSV: %v", err)
		}

		fmt.Printf("Transaction sent for date %s: %s\n", entry.Data, tx.Hash().Hex())
		fmt.Printf("Transaction receipt for date %s: %+v\n", entry.Data, receipt)
		fmt.Printf("Gas used for date %s: %d\n", entry.Data, receipt.GasUsed)
	}
}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(time.Second)
	}
}

func createAuth(privateKey string, chainID *big.Int) (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}
	return auth, nil
}

func saveCSV(timestamp string, gasUsed uint64) error {
	file, err := os.OpenFile("gasreport.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{timestamp, fmt.Sprintf("%d", gasUsed)})
	if err != nil {
		return err
	}

	return nil
}
