package main

import (
	"./lib"
	"bytes"
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	issuer        = kf.GenPrivKey()
	SYSTEM_ISSUER = issuer.PubKey().Address()
)

type TokenApp struct {
	Accounts map[string]int
}

func NewTokenApp() *TokenApp {
	return &TokenApp{Accounts: map[string]int{}}
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
	if !bytes.Equal(issuer, SYSTEM_ISSUER) {
		return errors.New("invalid issuer")
	}
	app.Accounts[to.String()] += value
	return nil
}

func (app *TokenApp) Dump() {
	fmt.Printf("state => %v\n", app.Accounts)
}

func main() {
	app := NewTokenApp()
	p1 := kf.GenPrivKey()
	p2 := kf.GenPrivKey()

	app.issue(SYSTEM_ISSUER, p1.PubKey().Address(), 1000)
	app.transfer(p1.PubKey().Address(), p2.PubKey().Address(), 100)
	app.Dump()

	txIssue := lib.NewTx(lib.NewIssuePayload(
		issuer.PubKey().Address(),
		p1.PubKey().Address(),
		1000))
	txIssue.Sign(issuer)

	rawtx, err := lib.MarshalBinary(txIssue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue tx encoded => %v\n", rawtx)

	var txReceived lib.Tx
	err = lib.UnmarshalBinary(rawtx, &txReceived)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue tx decoded => %+v\n", txReceived)
	fmt.Printf("validated => %t\n", txReceived.Verify())
}
