package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/cmd/tendermint/commands"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
)

func main() {
	root := commands.RootCmd
	root.AddCommand(commands.GenNodeKeyCmd)
	root.AddCommand(commands.GenValidatorCmd)
	root.AddCommand(commands.InitFilesCmd)
	root.AddCommand(commands.ResetAllCmd)
	root.AddCommand(commands.ShowNodeIDCmd)
	root.AddCommand(commands.TestnetFilesCmd)

	app := NewApp()
	nodeProvider := makeNodeProvider(app)
	root.AddCommand(commands.NewRunNodeCmd(nodeProvider))

	exec := cli.PrepareBaseCmd(root, "wiz", ".")
	exec.Execute()
}

type App struct {
	types.BaseApplication
	Value int
}

func NewApp() *App {
	return &App{}
}

func (app *App) CheckTx(tx []byte) (rsp types.ResponseCheckTx) {
	if tx[0] == 0x01 || tx[0] == 0x02 || tx[0] == 0x03 {
		rsp.Log = "tx accepted"
		return
	}
	rsp.Code = 1
	rsp.Log = "bad tx rejected"
	return
}

func (app *App) DeliverTx(tx []byte) (rsp types.ResponseDeliverTx) {
	fmt.Println("=================>")
	switch tx[0] {
	case 0x01:
		app.Value += 1
	case 0x02:
		app.Value -= 1
	case 0x03:
		app.Value = 0
	default:
		rsp.Log = "weird command"
	}
	return
}

func (app *App) Query(req types.RequestQuery) (rsp types.ResponseQuery) {
	rsp.Log = fmt.Sprintf("counter: %d", app.Value)
	return
}

/*
func (app *App) Commit() (rsp types.ResponseCommit) {
  bz,_ := json.Marshal(app.Value)
  rsp.Data = bz
  return
}
*/

func makeNodeProvider(app types.Application) node.NodeProvider {
	return func(config *cfg.Config, logger log.Logger) (*node.Node, error) {
		nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
		if err != nil {
			return nil, err
		}

		return node.NewNode(config,
			privval.LoadOrGenFilePV(config.PrivValidatorFile()),
			nodeKey,
			proxy.NewLocalClientCreator(app),
			node.DefaultGenesisDocProviderFunc(config),
			node.DefaultDBProvider,
			node.DefaultMetricsProvider(config.Instrumentation),
			logger,
		)
	}
}
