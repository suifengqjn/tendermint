package main

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

type EzApp struct {
	types.BaseApplication
}

func NewEzApp() *EzApp {
	return &EzApp{}
}

func (app *EzApp) InitChain(req types.RequestInitChain) (rsp types.ResponseInitChain) {
	fmt.Printf("initchian => %+v\n", req)
	return
}

func (app *EzApp) Info(req types.RequestInfo) (rsp types.ResponseInfo) {
	fmt.Printf("info => %+v\n", req)
	return
}

func (app *EzApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	fmt.Printf("query => %+v\n", req)
	return
}

func (app *EzApp) CheckTx(tx []byte) (rsp types.ResponseCheckTx) {

	rsp = types.ResponseCheckTx{Code: 1, GasWanted: 1}
	fmt.Printf("checktx => %x\n", tx)
	return
}

func (app *EzApp) BeginBlock(req types.RequestBeginBlock) (rsp types.ResponseBeginBlock) {
	fmt.Printf("beginblock => %+v\n", req)
	return
}

func (app *EzApp) DeliverTx(tx []byte) (rsp types.ResponseDeliverTx) {
	fmt.Printf("delivertx => %v+\n", tx)
	return
}

func (app *EzApp) EndBlock(req types.RequestEndBlock) (rsp types.ResponseEndBlock) {
	fmt.Printf("endblock => %+v\n", req)
	return
}

func (app *EzApp) Commit() (rsp types.ResponseCommit) {
	fmt.Printf("commit => \n")
	return
}

func main() {
	app := NewEzApp()
	svr, err := server.NewServer(":26658", "socket", app)
	if err != nil {
		panic(err)
	}
	svr.Start()
	defer svr.Stop()
	fmt.Println("abci server started.")
	select {}
}
