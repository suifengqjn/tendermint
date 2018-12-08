package main

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

type CounterApp struct {
	types.BaseApplication
}

func NewCounterApp() *CounterApp {
	return &CounterApp{}
}

func (app *CounterApp) CheckTx(tx []byte) (rsp types.ResponseCheckTx) {
	if tx[0] < 0x04 {
		rsp.Log = "tx accepted"
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
