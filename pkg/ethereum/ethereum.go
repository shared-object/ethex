package ethereum

import (
	"encoding/hex"
	"log"

	"github.com/tyler-smith/go-bip39"
	"github.com/xllwhoami/etherix/pkg/hdwallet"
)

func NewSeedPhrase() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}

func ExtractAddressAndPrivateKey(seedPhrase string) (string, string, error) {
	wallet, err := hdwallet.NewFromMnemonic(seedPhrase)

	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", "", nil
	}

	privateKeyBytes, _ := wallet.PrivateKeyBytes(account)

	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	return account.Address.Hex(), privateKeyHex, nil
}
