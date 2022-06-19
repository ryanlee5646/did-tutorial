package main

import (
	"fmt"
	core "github.com/ryanlee5646/did-example/core"
)

func main() {
	did, _ := core.NewDID("dgbds", "12345")
	fmt.Printf("DID: [%s]", did.String())
}
