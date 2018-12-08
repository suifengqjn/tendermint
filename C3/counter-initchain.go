package main

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

type CounterApp struct {
	types.BaseApplication
	Value int
}

func NewCounterApp() *CounterApp {
	return &CounterApp{}
}

func (app *CounterApp) InitChain(req types.RequestInitChain) (rsp types.ResponseInitChain) {
	var state map[string]int
	err := json.Unmarshal(req.AppStateBytes, &state)
	if err != nil {
		panic(err)
	}
	app.Value = state["counter"]
	return
}

func (app *CounterApp) DeliverTx(tx []byte) (rsp types.ResponseDeliverTx) {
	switch tx[0] {
	case 0x01:
		app.Value += 1
	case 0x02:
		app.Value -= 1
	case 0x03:
		app.Value = 0
	}
	rsp.Log = fmt.Sprintf("state updated : %d", app.Value)
	return
}

func (app *CounterApp) CheckTx(tx []byte) (rsp types.ResponseCheckTx) {
	if tx[0] < 0x04 {
		return
	}
	rsp.Code = 1
	rsp.Log = "bad tx rejected"
	return
}

func main() {
	app := NewCounterApp()
	svr, err := server.NewServer(":26658", "socket", app)
	if err != nil {
		panic(err)
	}
	svr.Start()
	defer svr.Stop()
	fmt.Println("abci server started.")
	select {}
}
