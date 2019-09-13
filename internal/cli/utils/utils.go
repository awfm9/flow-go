package utils

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/dapperlabs/bamboo-node/pkg/crypto"
)

func Exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(1)
}

func DecodePrivateKey(derHex string) (crypto.PrKey, error) {
	salg, err := crypto.NewSignatureAlgo(crypto.ECDSA_P256)
	if err != nil {
		return nil, err
	}

	prKeyDer, err := hex.DecodeString(derHex)
	if err != nil {
		return nil, err
	}

	return salg.DecodePrKey(prKeyDer)
}

func EncodePrKey(sk crypto.PrKey) (string, error) {
	salg, err := crypto.NewSignatureAlgo(crypto.ECDSA_P256)
	if err != nil {
		return "", err
	}

	skBytes, err := salg.EncodePrKey(sk)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(skBytes), nil
}

func EncodePublicKey(pubKey crypto.PubKey) ([]byte, error) {
	salg, err := crypto.NewSignatureAlgo(crypto.ECDSA_P256)
	if err != nil {
		return nil, err
	}

	return salg.EncodePubKey(pubKey)
}
