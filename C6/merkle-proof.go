package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type sh struct {
	value string
}

func (h sh) Hash() []byte {
	return tmhash.Sum([]byte(h.value))
}

func sliceDemo() {

	data := [][]byte{[]byte("one"),[]byte("two"),[]byte("three"),[]byte("four")}
	root := merkle.SimpleHashFromByteSlices(data)
	root1, proofs :=merkle.SimpleProofsFromByteSlices(data)

	if hex.EncodeToString(root) != hex.EncodeToString(root1) {
		panic("error")
	}

	for i, v := range data {
		itemHash := tmhash.Sum(v)
		proof := proofs[i]
		valid := proof.Verify(root, itemHash)
		fmt.Println("check hash",valid)
	}


	for i, _ := range data {
		proof := proofs[i]
		lefthash := proof.LeafHash
		valid := proof.Verify(root, lefthash)
		fmt.Println("check lefthash",valid)
	}

}

func mapDemo() {

	data := map[string][]byte{
		"tom":   []byte("actor"),
		"mary":  []byte("teacher"),
		"linda": []byte("scientist"),
		"luke":  []byte("fisher")}
	root, proofs, keys := merkle.SimpleProofsFromMap(data)
	fmt.Printf("root hash => %x\n", root)
	fmt.Printf("proof for tom => %+v\n", proofs["tom"])
	fmt.Printf("keys sorted => %v\n", keys)


	valid := proofs["tom"].Verify(root, proofs["tom"].LeafHash)
	fmt.Printf("data[\"tom\"] is valid? => %t\n", valid)
}

func main() {
	sliceDemo()
	mapDemo()
}
