// core/ecdsa.go
package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/multiformats/go-multibase"
	"log"
	"math/big"
)

// 리턴값을 위에ㅓ 선언해주면 return이 없더라도 선언한 형식대로 return

// curve 정의(curve를 상수로 선언 ex: p256
const (
	ELLIPTIC_CURVE = "p256"
)

// ECDSAManager 구조체 (어차피 Privatekey안에 pub이 잇지만 편리를 위해 분리)
type ECDSAManager struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

// not use now.
type ECDSAInterface interface {
	Sign(msg string) ([]byte, error)
	SignToString(msg string) (string, error)
	Verify() bool
	Encode() string
	//Decode() string
}

// ECDSA 서명 구조체
type Signature struct {
	R *big.Int
	S *big.Int
}

// serialization ecdsa signature
func (s Signature) String() string {
	return s.R.String() + s.S.String()
}

// Generate new ecdsa keys
func NewEcdsa() (ecdsa *ECDSAManager) {
	ecdsa = new(ECDSAManager)
	err := ecdsa.Generate()
	if err != nil {
		log.Printf("Fail to ECDSA Generate.")
		return nil
	}
	// 리턴값을 위에 선언해주면 return이 없더라도 선언한 형식대로 return
	return
}

func (e *ECDSAManager) Generate() error {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		return errors.New("ECDSA Keypair generation was Fail!")
	}

	e.PrivateKey = pvKey
	e.PublicKey = &pvKey.PublicKey

	return nil
}

func (e *ECDSAManager) Sign(digest []byte) (*Signature, error) {
	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err := ecdsa.Sign(rand.Reader, e.PrivateKey, digest)
	if err != nil {
		return nil, err //errors.New("failed to sign to msg.")
	}

	// prepare a signature structure to marshal into json
	signature := &Signature{
		R: r,
		S: s,
	}
	/*
		signature := r.Bytes()
		signature = append(signature, s.Bytes()...)
	*/
	return signature, nil
}

func (e *ECDSAManager) Verify(signature *Signature, digest []byte) bool {
	return ecdsa.Verify(e.PublicKey, digest, signature.R, signature.S)
}

func (e *ECDSAManager) SignToString(digest []byte) (string, error) {
	signature, err := e.Sign(digest)
	if err != nil {
		return "", err
	}

	return signature.String(), nil
}

func (e *ECDSAManager) PublicKeyToString() (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&e.PublicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return "", err
	}

	publicKeyHash := sha256.Sum256(publicKeyBytes)

	return hex.EncodeToString(publicKeyHash[:]), nil
}

func (e *ECDSAManager) PublicKeyBase58() string {
	if e.PublicKey == nil {
		return ""
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}

	return base58.Encode(publicKeyBytes)
}

func (e *ECDSAManager) PublicKeyMultibase() string {
	if e.PublicKey == nil {
		return ""
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)

	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}

	// The current multibase table: https://github.com/multiformats/multibase/blob/master/multibase.csv
	str, err := multibase.Encode(multibase.Base58BTC, publicKeyBytes)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}
	return str
}

func (e *ECDSAManager) PrintPublicKey() {
	str, err := e.PublicKeyToString()
	if err != nil {

	}
	fmt.Printf("Public Key: %s\n", str)
}
