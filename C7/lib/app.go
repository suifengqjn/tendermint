package lib

import (
	"bytes"
	"errors"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	SYSTEM_ISSUER = crypto.Address("KING_OF_TOKEN")
)

type TokenApp struct {
	types.BaseApplication
	store *Store
}

func NewTokenApp() *TokenApp {
	store := NewStore()
	return &TokenApp{store: store}
}

func (app *TokenApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	addr := crypto.Address(req.Data)
	rsp.Key = req.Data
	bal, _ := app.store.GetBalance(addr)
	bz, _ := codec.MarshalBinaryBare(bal)
	rsp.Value = bz
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
	fromBalance, _ := app.store.GetBalance(from)
	if fromBalance < value {
		return errors.New("balance low")
	}
	app.store.SetBalance(from, fromBalance-value)
	toBalance, _ := app.store.GetBalance(to)
	app.store.SetBalance(to, toBalance+value)
	return nil
}

func (app *TokenApp) issue(issuer, to crypto.Address, value int) error {
	wallet := LoadWallet("./wallet")
	SYSTEM_ISSUER = wallet.GetAddress("issuer")
	if !bytes.Equal(issuer, SYSTEM_ISSUER) {
		return errors.New("invalid issuer")
	}
	toBalance, _ := app.store.GetBalance(to)
	app.store.SetBalance(to, toBalance+value)
	return nil
}
