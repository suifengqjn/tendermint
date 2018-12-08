package main

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	"io/ioutil"
)

type CounterApp struct {
	types.BaseApplication
	Value   int
	Version int64
	History map[int64]int
}

func NewCounterApp() *CounterApp {
	app := &CounterApp{History: map[int64]int{}}
	bz, err := ioutil.ReadFile("./counter.state")
	if err != nil {
		return app
	}
	err = json.Unmarshal(bz, app)
	if err != nil {
		return app
	}
	return app
}

func (app *CounterApp) Commit() (rsp types.ResponseCommit) {
	app.Version += 1
	app.History[app.Version] = app.Value
	bz, err := json.Marshal(app)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("./counter.state", bz, 0644)
	return
}

func (app *CounterApp) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	ver := req.Height
	if ver == 0 {
		ver = app.Version
	}
	rsp.Key = []byte("counter")
	rsp.Value = []byte(fmt.Sprintf("%d", app.History[ver]))
	rsp.Height = ver
	rsp.Log = fmt.Sprintf("value:%d", app.History[ver])
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
