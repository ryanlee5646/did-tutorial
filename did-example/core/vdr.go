package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/ryanlee5646/did-example/config"
	"github.com/ryanlee5646/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"time"
)

// DID를 레지스터에 등록
func RegisterDid(did string, didDocument string) error {
	conn, err := grpc.Dial(config.SystemConfig.RegistrarAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Registrar not connect: %v\n", err)
		return errors.New(fmt.Sprintf("Registrar not connect: %v", err))
	}
	defer conn.Close()

	client := protos.NewRegistrarClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.RegisterDid(ctx, &protos.RegistrarRequest{Did: did, DidDocument: didDocument})
	if err != nil {
		log.Println("레지스터에 DID 등록 실패.")
		return errors.New("DID 등록실패")
	}

	fmt.Printf("Registrar Response: %s\n", res)

	return nil
}

// 레지스터에 있는 DIDDocument를 호출
func ResolveDid(did string) (string, error) {
	conn, err := grpc.Dial(config.SystemConfig.ResolverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Resolver not connect: %v\n", err)
		return "", errors.New(fmt.Sprintf("Resolver not connect: %v", err))
	}
	defer conn.Close()

	client := protos.NewResolverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ResolveDid(ctx, &protos.ResolverRequest{Did: did})
	if err != nil {
		log.Fatalf("Failed to resolve DID.")
	}

	fmt.Printf("Result: %s\n", res)

	return res.DidDocument, nil
}
