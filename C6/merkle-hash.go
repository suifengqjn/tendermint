package main

import (
	"fmt"
	"github.com/tendermint/tendermint/crypto/merkle"
)


func sliceDemo() {

	data := [][]byte{[]byte("one"),[]byte("two"),[]byte("three"),[]byte("four")}
		//[]hasher{&sh{"one"}, &sh{"two"}, &sh{"three"}, &sh{"four"}}
	hash := merkle.SimpleHashFromByteSlices(data)
	fmt.Printf("root hash => %x\n", hash)
}

func mapDemo() {
	data := map[string][]byte{
		"tom":   []byte("actor"),
		"mary":  []byte("teacher"),
		"linda": []byte("scientist"),
		"luke":  []byte("fisher")}

	hash := merkle.SimpleHashFromMap(data)

	fmt.Printf("root hash => %x\n", hash)
}

func main() {
	sliceDemo()
	mapDemo()
}
