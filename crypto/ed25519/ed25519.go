package ed25519

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/ed25519"
)

func GenerateKey() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return priv.Seed(), pub, nil
}
func GenerateKeyBySeed(seed []byte) ([]byte, error) {
	priv := ed25519.NewKeyFromSeed(seed)

	pub := priv[32:]
	if len(pub) != 32 {
		return nil, errors.New("public key length is not equal 32")
	}
	return pub, nil
}

func Sign(privateKey, message []byte) ([]byte, error) {
	privKey := ed25519.NewKeyFromSeed(privateKey)
	sig := ed25519.Sign(privKey, message)
	if len(sig) != 64 {
		return nil, errors.New("sign fail,sig length is not equal 64")
	}
	return sig, nil
}
