package wallets

import (
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// Wallet 구조체 정의
type BitcoinWallet struct {
	PrivateKey string
	Address    string
}

// NewWallet 생성 함수
func NewBitcoinWallet() (*BitcoinWallet, error) {
	// 비트코인 개인 키 생성
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	// 공개 키에서 비트코인 주소 생성
	pubKey := privKey.PubKey()
	address, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// 개인 키를 16진수 문자열로 변환
	privKeyHex := hex.EncodeToString(privKey.Serialize())

	return &BitcoinWallet{
		PrivateKey: privKeyHex,
		Address:    address.String(),
	}, nil
}

func (w *BitcoinWallet) GetAddress() string {
	return w.Address
}
func (w *BitcoinWallet) GetPrivateKey() string {
	return w.PrivateKey
}

// PrintWallet 정보를 출력하는 메서드
func (w *BitcoinWallet) PrintWallet() {
	log.Println("Bitcoin Wallet:")
	log.Println("Private Key:", w.GetPrivateKey())
	log.Println("Address:", w.GetAddress())
}
