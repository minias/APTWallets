package main

import (
	"log"

	"github.com/minias/APTWallets/pkg/wallets"
)

func main() {
	b, err := wallets.CreateWallet("bitcoin")
	if err != nil {
		log.Printf("err: %v", err)
	}
	b.PrintWallet()
	e, err := wallets.CreateWallet("ethereum")
	if err != nil {
		log.Printf("err: %v", err)
	}
	e.PrintWallet()
	s, err := wallets.CreateWallet("solana")
	if err != nil {
		log.Printf("err: %v", err)
	}
	s.PrintWallet()
}
