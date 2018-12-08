package main

import (
	"./lib"
	"fmt"
	"github.com/tendermint/tendermint/rpc/client"
)

var (
	cli = client.NewHTTP("http://localhost:26657", "/websocket")
)

func main() {
	//initWallet()
	//issue()
	//transfer()
	//query("michael")
	query("britney")
}

func issue() {
	wallet := lib.LoadWallet("./wallet")
	tx := lib.NewTx(lib.NewIssuePayload(
		wallet.GetAddress("issuer"),
		wallet.GetAddress("michael"),
		1000))
	tx.Sign(wallet.GetPrivKey("issuer"))
	bz, err := lib.MarshalBinary(tx)
	if err != nil {
		panic(err)
	}
	ret, err := cli.BroadcastTxCommit(bz)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue ret => %+v\n", ret)
}

func transfer() {
	wallet := lib.LoadWallet("./wallet")
	tx := lib.NewTx(lib.NewTransferPayload(
		wallet.GetAddress("michael"),
		wallet.GetAddress("britney"),
		100))
	tx.Sign(wallet.GetPrivKey("michael"))
	bz, err := lib.MarshalBinary(tx)
	if err != nil {
		panic(err)
	}
	ret, err := cli.BroadcastTxCommit(bz)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue ret => %+v\n", ret)
}

func query(label string) {
	wallet := lib.LoadWallet("./wallet")
	ret, err := cli.ABCIQuery("", wallet.GetAddress(label))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ret => %+v\n", ret)
}

func initWallet() {
	wallet := lib.NewWallet()
	wallet.GenPrivKey("issuer")
	wallet.GenPrivKey("michael")
	wallet.GenPrivKey("britney")
	wallet.Save("./wallet")
}
