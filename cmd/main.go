package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/xllwhoami/etherix/internal/database"
	"github.com/xllwhoami/etherix/pkg/ethereum"
)

func saveWallet(address string, seedPhrase string, privateKeyHex string, resultFile string) {
	file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Content to append to the file
	content := "This is some new content.\n"

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func check(address string, db *database.Database, resultFile string) (bool, error) {
	result, err := db.SelectAddress(address)

	return result, err
}

func process(counter int, db *database.Database, resultFile string) {
	seedPhrase := ethereum.NewSeedPhrase()
	address, privateKeyHex, err := ethereum.ExtractAddressAndPrivateKey(seedPhrase)

	if err != nil {
		log.Panic(err)
	}

	result, err := check(address, db, resultFile)

	if err != nil {
		log.Panic(err)
	}

	if result {
		saveWallet(address, seedPhrase, privateKeyHex, resultFile)

		message := fmt.Sprintf("%d. Address founded in database", counter)

		color.Green(message)
	} else {
		message := fmt.Sprintf("%d. Address not found", counter)

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

	for {
		go process(counter, db, resultFile)
		counter++
	} //
}
