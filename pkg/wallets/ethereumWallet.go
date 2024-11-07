package wallets

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/sha3"
)

type EthereumWallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	Address    string
}

// NewEthereumWallet 생성 함수
func NewEthereumWallet() (Wallet, error) {
	return NewWallet()
}
func NewWallet() (*EthereumWallet, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	pubKey := privKey.PublicKey
	address := GenerateEthereumAddress(&pubKey)

	return &EthereumWallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
		Address:    address,
	}, nil
}

func GenerateEthereumAddress(pubKey *ecdsa.PublicKey) string {
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes)
	addressBytes := hash.Sum(nil)[len(hash.Sum(nil))-20:]
	return "0x" + hex.EncodeToString(addressBytes)
}

// PrintWallet 정보를 출력하는 메서드
func (w *EthereumWallet) PrintWallet() {
	log.Println("Ethereum Wallet:")
	log.Println("Private Key:", w.GetPrivateKey())
	log.Println("Public Key:", w.GetPublicKey())
	log.Println("Address:", w.GetAddress())
}
func (w *EthereumWallet) GetPrivateKey() string {
	//return hex.EncodeToString(w.PrivateKey.D.Bytes())
	return fmt.Sprintf("%x", w.PrivateKey.D)
}

func (w *EthereumWallet) GetPublicKey() string {
	return fmt.Sprintf("%x", w.PublicKey.X) // X, Y 결합해야 함
}

func (w *EthereumWallet) GetAddress() string {
	return w.Address
}
