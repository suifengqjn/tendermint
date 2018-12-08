package lib

import (
	"bytes"
	"github.com/tendermint/tendermint/crypto"
	"time"
)

type Tx struct {
	Payload   Payload
	Signature []byte
	PubKey    crypto.PubKey
	Sequence  int64
}

func NewTx(payload Payload) *Tx {
	return &Tx{Payload: payload, Sequence: time.Now().Unix()}
}

func (tx *Tx) Verify() bool {
	signer := tx.Payload.GetSigner()
	signerFromKey := tx.PubKey.Address()
	if !bytes.Equal(signer, signerFromKey) {
		return false
	}
	data := tx.Payload.GetSignBytes()
	sig := tx.Signature
	valid := tx.PubKey.VerifyBytes(data, sig)
	if !valid {
		return false
	}
	return true
}

func (tx *Tx) Sign(priv crypto.PrivKey) error {
	data := tx.Payload.GetSignBytes()
	var err error
	tx.Signature, err = priv.Sign(data)
	tx.PubKey = priv.PubKey()
	return err
}
