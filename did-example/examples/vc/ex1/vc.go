package main

import (
	"fmt"
	"github.com/ryanlee5646/did-example/core"
	"os"
)

// Issuer에 의한 VC 발행 예시.
func main() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	issuerKeyEcdsa := core.NewEcdsa()

	// DID 생성.
	issuerDid, _ := core.NewDID("dgbds", issuerKeyEcdsa.PublicKeyBase58())

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)

	//verificationMethod := []core.VerificationMethod{
	//	{
	//		Id:                 verificationId,
	//		Type:               "EcdsaSecp256k1VerificationKey2019",
	//		Controller:         issuerDid.String(),
	//		PublicKeyMultibase: issuerKeyEcdsa.PublicKeyMultibase(),
	//	},
	//}

	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuerDid.String(),
		map[string]interface{}{
			"id": "1234567890",
			"alumniOf": map[string]interface{}{
				"id": "1234567",
				"name": []map[string]string{
					{
						"value": "Example University",
						"lang":  "en",
					}, {
						"value": "Exemple d'Université",
						"lang":  "fr",
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("VC 생성실패!!.")
		os.Exit(0)
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token := vc.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)

	// 생성된 VC를 검증한다.(public key를 사용해서 검증)
	res, _ := vc.VerifyJwt(token, issuerKeyEcdsa.PublicKey)

	if res {
		fmt.Println("VC is 검증완료!.")
	} else {
		fmt.Println("VC is 검증실패!.")
	}

}
