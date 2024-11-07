package wallets

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"

	"golang.org/x/crypto/ed25519"
)

type SolanaWallet struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string
}

func NewSolanaWallet() (*SolanaWallet, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	address := GenerateSolanaAddress(pubKey)

	return &SolanaWallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
		Address:    address,
	}, nil
}

func GenerateSolanaAddress(pubKey []byte) string {
	return base58.Encode(pubKey)
}

func (w *SolanaWallet) GetPrivateKey() string {
	return fmt.Sprintf("%x", w.PrivateKey)
}

func (w *SolanaWallet) GetPublicKey() string {
	return fmt.Sprintf("%x", w.PublicKey)
}

func (w *SolanaWallet) GetAddress() string {
	return w.Address
}

// PrintWallet 정보를 출력하는 메서드
func (w *SolanaWallet) PrintWallet() {
	log.Println("Solana Wallet:")
	log.Println("Private Key:", w.GetPrivateKey())
	log.Println("Public Key:", w.GetPublicKey())
	log.Println("Address:", w.GetAddress())
}
