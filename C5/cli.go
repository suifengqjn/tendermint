package main

import (
	"./lib"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
)

var (
	cli = client.NewHTTP("http://localhost:26657", "/websocket")
)

func main() {
	rootCmd := &cobra.Command{
		Use: "cli",
	}

	walletCmd := &cobra.Command{
		Use: "init-wallet",
		Run: func(cmd *cobra.Command, args []string) { initWallet() },
	}

	issueCmd := &cobra.Command{
		Use: "issue-tx",
		Run: func(cmd *cobra.Command, args []string) { issue() },
	}

	transferCmd := &cobra.Command{
		Use: "transfer-tx",
		Run: func(cmd *cobra.Command, args []string) { transfer() },
	}

	queryCmd := &cobra.Command{
		Use: "query",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("query who?")
			}
			label := args[0]
			query(label)
			return nil
		},
	}

	rootCmd.AddCommand(walletCmd)
	rootCmd.AddCommand(issueCmd)
	rootCmd.AddCommand(transferCmd)
	rootCmd.AddCommand(queryCmd)

	rootCmd.Execute()

	//initWallet()
	//issue()
	//transfer()
	//query("michael")
	//query("britney")
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

	fmt.Printf("ret => %+v\n", ret.Response)
}

func initWallet() {
	wallet := lib.NewWallet()
	wallet.GenPrivKey("issuer")
	wallet.GenPrivKey("michael")
	wallet.GenPrivKey("britney")
	wallet.Save("./wallet")
}
