package lib

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	SYSTEM_ISSUER = crypto.Address("KING_OF_TOKEN")
)

type TokenApp struct {
	types.BaseApplication
	Accounts map[string]int
}

func NewTokenApp() *TokenApp {
	return &TokenApp{Accounts: map[string]int{}}
}

func (app *TokenApp) Commit() (rsp types.ResponseCommit) {
	rsp.Data = app.getRootHash()
	return
}

func (app *TokenApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	addr := crypto.Address(req.Data)
	rsp.Key = req.Data
	rsp.Value, _ = codec.MarshalBinaryBare(app.Accounts[addr.String()])
	rsp.Height = req.Height
	rsp.Proof = app.getProof(addr.String())
	//rsp.Proof = app.getProofBytes(addr.String())
	fmt.Println("proof", app.getProofBytes(addr.String()))
	return
}

func (app *TokenApp) DeliverTx(raw []byte) (rsp types.ResponseDeliverTx) {
	tx, _ := app.decodeTx(raw)
	switch tx.Payload.GetType() {
	case "issue":
		pld := tx.Payload.(*IssuePayload)
		err := app.issue(pld.Issuer, pld.To, pld.Value)
		if err != nil {
			rsp.Log = err.Error()
		}
		rsp.Info = "issue tx applied"
	case "transfer":
		pld := tx.Payload.(*TransferPayload)
		err := app.transfer(pld.From, pld.To, pld.Value)
		if err != nil {
			rsp.Log = err.Error()
		}
		rsp.Info = "transfer tx applied"
	}
	return
}

func (app *TokenApp) CheckTx(raw []byte) (rsp types.ResponseCheckTx) {
	tx, err := app.decodeTx(raw)
	if err != nil {
		rsp.Code = 1
		rsp.Log = "decode error"
		return
	}
	if !tx.Verify() {
		rsp.Code = 2
		rsp.Log = "verify failed"
		return
	}
	return
}

func (app *TokenApp) decodeTx(raw []byte) (*Tx, error) {
	var tx Tx
	err := codec.UnmarshalBinaryBare(raw, &tx)
	return &tx, err
}

func (app *TokenApp) transfer(from, to crypto.Address, value int) error {
	if app.Accounts[from.String()] < value {
		return errors.New("balance low")
	}
	app.Accounts[from.String()] -= value
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp) issue(issuer, to crypto.Address, value int) error {
	wallet := LoadWallet("./wallet")
	SYSTEM_ISSUER = wallet.GetAddress("issuer")
	if !bytes.Equal(issuer, SYSTEM_ISSUER) {
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp) Dump() {
	fmt.Printf("state => %v\n", app.Accounts)
}
