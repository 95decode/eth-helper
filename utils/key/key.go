package key

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	StandardScryptN = 1 << 18
	StandardScryptP = 1
)

func GenerateKey(encrypted bool) (string, *keystore.KeyStore) {
	d := "./keys"
	newKs := keystore.NewPlaintextKeyStore

	if encrypted {
		newKs = func(kd string) *keystore.KeyStore { return keystore.NewKeyStore(kd, StandardScryptN, StandardScryptP) }
	}

	return d, newKs(d)
}
