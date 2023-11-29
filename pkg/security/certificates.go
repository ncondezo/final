package security

import (
	"crypto/rsa"
	"os"
	"sync"

	"github.com/golang-jwt/jwt"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	once       sync.Once
)

func LoadCertificates(privateKeyPath, publicKeyPath string) error {
	var err error
	once.Do(func() {
		err = loadCertificates(privateKeyPath, publicKeyPath)
	})
	return err
}

func loadCertificates(privateKeyPath, publicKeyPath string) error {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return err
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return err
	}
	return nil
}
