package libganache

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	PublicKeyString  string
	PrivateKeyString string
	PublicKey        common.Address
	PrivateKey       *ecdsa.PrivateKey
}

type AccountsFile struct {
	Addresses   map[string]string `json:"addresses"`
	PrivateKeys map[string]string `json:"private_keys"`
}

// This was used by ganache-cli
/*
type LegacyAccountsFile struct {
	Addresses map[string]struct {
		SecretKey struct {
			Type string `json:"type"`
			Data []byte `json:"data"`
		} `json:"secretKey"`
		PublicKey struct {
			Type string `json:"type"`
			Data []byte `json:"data"`
		} `json:"publicKey"`
		Address string `json:"address"`
		Account struct {
			Nonce     string `json:"nonce"`
			Balance   string `json:"balance"`
			StateRoot string `json:"stateRoot"`
			CodeHash  string `json:"codeHash"`
		} `json:"account"`
	} `json:"addresses"`
	PrivateKeys map[string]string `json:"private_keys"`
}
*/
