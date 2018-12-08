package main

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

type CounterApp struct {
	types.BaseApplication
	Value   int
	Version int64
	History map[int64]int
}

func NewCounterApp() *CounterApp {
	app := &CounterApp{History: map[int64]int{}}
	return app
}

func (app *CounterApp) Info(req types.RequestInfo) (rsp types.ResponseInfo) {
	rsp.LastBlockHeight = app.Version
	return
}

func (app *CounterApp) Commit() (rsp types.ResponseCommit) {
	app.Version += 1
	app.History[app.Version] = app.Value
	return
}

func (app *CounterApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	ver := req.Height
	if ver == 0 {
		ver = app.Version
	}
	rsp.Key = []byte("counter")
	value := app.History[ver]
	rsp.Value = []byte(fmt.Sprintf("%d", value))
	rsp.Height = ver
	rsp.Log = fmt.Sprintf("value@%d : %d", ver, value)
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
