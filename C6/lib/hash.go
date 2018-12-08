package lib

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type Balance int

func (b Balance) Hash() []byte {
	v, _ := codec.MarshalBinaryBare(b)
	return tmhash.Sum(v)
}

func (app *TokenApp) stateToHasherMap() map[string][]byte {
	hashers := map[string][]byte{}
	for addr, val := range app.Accounts {
		balance := int(val)
		hashers[addr] = IntToBytes(balance)
	}
	return hashers
}

func (app *TokenApp) getRootHash() []byte {
	hashers := app.stateToHasherMap()
	return merkle.SimpleHashFromMap(hashers)
}

func (app *TokenApp) getProofBytes(addr string) []byte {
	hashers := app.stateToHasherMap()
	_, proofs, _ := merkle.SimpleProofsFromMap(hashers)

	bz, err := codec.MarshalBinaryBare(proofs[addr])
	if err != nil {
		return nil
	}
	return bz
}

func (app *TokenApp)getProof(addr string) *merkle.Proof  {
	hashers := app.stateToHasherMap()
	_, proofs, _ := merkle.SimpleProofsFromMap(hashers)

	addByte, _ := hex.DecodeString(addr)
	simpleValueOp := merkle.NewSimpleValueOp(addByte,proofs[addr])
	pop := simpleValueOp.ProofOp()
	return &merkle.Proof{Ops:[]merkle.ProofOp{pop}}

}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}