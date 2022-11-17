package key

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func GenerateKey(encrypted bool) (string, *keystore.KeyStore) {
	d := "./keys"
	newKs := keystore.NewPlaintextKeyStore
	if encrypted {
		newKs = func(kd string) *keystore.KeyStore { return keystore.NewKeyStore(kd, 2, 1) }
	}
	return d, newKs(d)
}
