package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
)

var (
	SYSTEM_ISSUER = crypto.Address("KING_OF_TOKEN")
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
	a1 := crypto.Address("TEST_ADDR1")
	a2 := crypto.Address("TEST_ADDR2")

	app.issue(SYSTEM_ISSUER, a1, 1000)
	app.transfer(a1, a2, 100)
	app.Dump()
}
