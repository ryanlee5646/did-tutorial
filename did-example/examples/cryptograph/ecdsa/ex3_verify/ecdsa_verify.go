package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s Signature) String() string {
	return s.R.String() + s.S.String()
}

func sign(digest []byte, pvKey *ecdsa.PrivateKey) (*Signature, error) {
	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err := ecdsa.Sign(rand.Reader, pvKey, digest)
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

func SignToString(digest []byte, pvKey *ecdsa.PrivateKey) (string, error) {
	signature, err := sign(digest, pvKey)
	if err != nil {
		return "", err
	}

	return signature.String(), nil
}

func verify(signature *Signature, digest []byte, pbKey *ecdsa.PublicKey) bool {
	return ecdsa.Verify(pbKey, digest, signature.R, signature.S)
}

func main() {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	msg := "Hello World."
	digest := sha256.Sum256([]byte(msg))
	signature, err := sign(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}

	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n", signature.R, signature.S)
	fmt.Printf("Signature: %+v\n", signature.String())

	msg = "Hello World" // 메세지를 조금만 변경해도 검증이 안되는거 확인
	digest = sha256.Sum256([]byte(msg))

	pbKey := &pvKey.PublicKey
	ret := verify(signature, digest[:], pbKey)

	fmt.Printf("########## Verification ##########\n")
	if ret {
		fmt.Printf("Signature verifies")
	} else {
		fmt.Printf("Signature does not verify")
	}

}
