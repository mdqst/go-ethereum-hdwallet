package main

import (
	"flag"
	"fmt"
	"log"

	hdwallet "github.com/stephenlacy/go-ethereum-hdwallet"
)

func main() {
	var mnemonic string
	var hdpath string

	flag.StringVar(&mnemonic, "mnemonic", "", "Mnemonic")
	flag.StringVar(&hdpath, "path", "", "HD path")
	flag.Parse()

	// Check for empty values
	if mnemonic == "" {
		log.Fatal("Mnemonic phrase is required")
	}

	if hdpath == "" {
		log.Fatal("HD path is required")
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath(hdpath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("public address:", account.Address.Hex())
	fmt.Println("private key:", privateKey)
}
