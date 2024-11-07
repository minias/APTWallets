package wallets

import "fmt"

// Wallet 인터페이스 정의
type Wallet interface {
	PrintWallet()
}

// 통합된 월렛 생성 함수
func CreateWallet(walletType string) (Wallet, error) {
	switch walletType {
	case "bitcoin":
		return NewBitcoinWallet()
	case "ethereum":
		return NewEthereumWallet()
	case "solana":
		return NewSolanaWallet()
	default:
		return nil, fmt.Errorf("unsupported wallet type: %s", walletType)
	}
}
