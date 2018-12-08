package lib

import (
	"encoding/json"
	"github.com/tendermint/tendermint/crypto"
)

type Payload interface {
	GetSigner() crypto.Address
	GetSignBytes() []byte
	GetType() string
}

type IssuePayload struct {
	Issuer crypto.Address
	To     crypto.Address
	Value  int
}

func NewIssuePayload(issuer, to crypto.Address, value int) *IssuePayload {
	return &IssuePayload{issuer, to, value}
}

func (pld *IssuePayload) GetSigner() crypto.Address {
	return pld.Issuer
}

func (pld *IssuePayload) GetSignBytes() []byte {
	bz, err := json.Marshal(pld)
	if err != nil {
		return []byte{}
	}
	return bz
}

func (pld *IssuePayload) GetType() string {
	return "issue"
}

type TransferPayload struct {
	From  crypto.Address
	To    crypto.Address
	Value int
}

func NewTransferPayload(from, to crypto.Address, value int) *TransferPayload {
	return &TransferPayload{from, to, value}
}

func (pld *TransferPayload) GetSigner() crypto.Address {
	return pld.From
}

func (pld *TransferPayload) GetSignBytes() []byte {
	bz, err := json.Marshal(pld)
	if err != nil {
		return []byte{}
	}
	return bz
}

func (pld *TransferPayload) GetType() string {
	return "transfer"
}
