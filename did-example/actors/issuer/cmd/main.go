package main

import (
	"github.com/ryanlee5646/did-example/actors/issuer"
	"github.com/ryanlee5646/did-example/config"
	protos "github.com/ryanlee5646/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	// New Issuer
	issr := new(issuer.Issuer)
	issr.GenerateDID()

	lis, err := net.Listen("tcp", config.SystemConfig.IssuerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	issuerServer := issuer.Server{}
	issuerServer.Issuer = issr

	s := grpc.NewServer()
	protos.RegisterSimpleIssuerServer(s, &issuerServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
