package main

import (
	"fmt"
	"github.com/ryanlee5646/did-example/core"
	"log"
)

func main() {
	var method = "dgbds"
	// 키 생성 ECDSAManger 객체 생성
	kms := new(core.ECDSAManager)
	// 키 발급
	kms.Generate()

	// DID를 생성 (method:키 형태)
	did, err := core.NewDID(method, kms.PublicKeyMultibase())

	// 에러 발생했는지 확인
	if err != nil {
		log.Printf("DID발급하는데 에러발생, error: %v\n", err)
	}

	// DID Document 생성. (Verification ID는 고유해야하므로 did+#key-1 하게되면 고유할수밖에 업승ㅁ)
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)

	fmt.Println("### New DID")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

	RegisterDid(did.String(), didDocument)

	//Resolve한다.
	didDocumentStr, err := core.ResolveDid(did.String())
	if err != nil {
		log.Printf("DID 리졸브 실패!.\nError: %x\n", err)
	}

	fmt.Printf("DID Document 리졸브 성공! => %+v\n", didDocumentStr)
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil
}
