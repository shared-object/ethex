package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/xllwhoami/etherix/internal/database"
	"github.com/xllwhoami/etherix/pkg/ethereum"
	"golang.org/x/sync/semaphore"
)

func saveWallet(address string, seedPhrase string, privateKeyHex string, resultFile string) {
	file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Content to append to the file
	content := fmt.Sprintf("Address: %s\nPrivate Key: %s\nSeed Phrase: %s\n\n", address, privateKeyHex, seedPhrase)

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func check(address string, db *database.Database) (bool, error) {
	result, err := db.SelectAddress(address)

	return result, err
}

func process(counter int, db *database.Database, resultFile string) {
	seedPhrase := ethereum.NewSeedPhrase()
	address, privateKeyHex, err := ethereum.ExtractAddressAndPrivateKey(seedPhrase)

	if err != nil {
		log.Panic(err)
	}

	result, err := check(address, db)

	if err != nil {
		log.Panic(err)
	}

	if result {
		saveWallet(address, seedPhrase, privateKeyHex, resultFile)

		message := fmt.Sprintf("%d. Address %s founded in database. Saved.", counter, address)

		color.Green(message)
	} else {
		message := fmt.Sprintf("%d. Address %s not found", counter, address)

		color.Red(message)
	}

}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseFilePath, ok := os.LookupEnv("DATABASE_FILE_PATH")

	if !ok {
		log.Fatal("DATABASE_FILE_PATH not specified")
	}

	resultFile, ok := os.LookupEnv("RESULT_FILE")

	if !ok {
		log.Fatal("RESULT_FILE not specified")
	}

	db, err := database.NewDatabase(databaseFilePath)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	sem := semaphore.NewWeighted(50)
	ctx := context.Background()

	for {
		if err := sem.Acquire(ctx, 10); err != nil {
			log.Panic(err)
		}
		if counter > 1000 {
			break
		}

		go func() {
			process(counter, db, resultFile)

			sem.Release(10)

			counter++
		}()

	}
}
