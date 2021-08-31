package ecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/blake2b"
	"math/big"
)

/*
secp256k1
*/

func GenerateKey() ([]byte, []byte, error) {
	cruve := secp256k1.S256()

	priv, err := ecdsa.GenerateKey(cruve, rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("generate private key error: %v", err)
	}
	if len(priv.D.Bytes()) != 32 {
		for {
			priv, _ = ecdsa.GenerateKey(cruve, rand.Reader)
			if len(priv.D.Bytes()) == 32 {
				break
			}
		}
	}
	privBytes := priv.D.Bytes()

	comPub := secp256k1.CompressPubkey(priv.X, priv.Y)
	pub := blake2b.Sum256(comPub)

	return privBytes, pub[:], nil

}
func GenerateKeyBySeed(seed []byte) ([]byte, error) {
	cruve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.D = new(big.Int).SetBytes(seed)
	priv.X, priv.Y = cruve.ScalarBaseMult(priv.D.Bytes())
	comPub := secp256k1.CompressPubkey(priv.X, priv.Y)
	pub := blake2b.Sum256(comPub)
	return pub[:], nil
}

func Sign(privateKey, message []byte) ([]byte, error) {
	return secp256k1.Sign(message, privateKey)
}
