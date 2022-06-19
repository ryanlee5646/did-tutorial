package main

import (
	"github.com/ryanlee5646/did-example/actors/verifier"
	"github.com/ryanlee5646/did-example/config"
	"github.com/ryanlee5646/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// New Issuer
	vrfr := new(verifier.Verifier)
	//issuer.generateDID()

	lis, err := net.Listen("tcp", config.SystemConfig.VerifierAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	verifierServer := verifier.Server{}
	verifierServer.Verifier = vrfr

	s := grpc.NewServer()
	protos.RegisterVerifierServer(s, &verifierServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
