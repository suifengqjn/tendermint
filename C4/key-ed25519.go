package main

import (
	"fmt"
	kf "github.com/tendermint/tendermint/crypto/ed25519"
)

func main() {
	priv := kf.GenPrivKey()
	fmt.Printf("private key => %v\n", priv)

	pub := priv.PubKey()
	fmt.Printf("public key => %v\n", pub)

	addr := pub.Address()
	fmt.Printf("address => %v\n", addr)
}
