package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	// pub,prv 키 발행
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // 커브함수들 => elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	pbKey := &pvKey.PublicKey

	fmt.Printf("########## Key pair ##########\n")
	fmt.Printf("===== Private Key =====\n")
	fmt.Printf("Private Key: %x\n", pvKey.D)
	fmt.Printf("===== Public Key(X, Y) =====\n")
	fmt.Printf("X=%s Y=%s\n", pbKey.X, pbKey.Y)
	fmt.Printf("  Hex: X=%x Y=%x\n\n", pbKey.X.Bytes(), pbKey.Y.Bytes()) //키를 바이트로 출력

}
