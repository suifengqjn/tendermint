package main

import (
	"./lib"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "token-wallet",
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init wallet",
		Run:   func(cmd *cobra.Command, args []string) { initWallet() },
	}

	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "load wallet",
		Run:   func(cmd *cobra.Command, args []string) { loadWallet() },
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.Execute()
}

func initWallet() {
	wallet := lib.NewWallet()
	wallet.GenPrivKey("issuer")
	wallet.GenPrivKey("michael")
	wallet.GenPrivKey("britney")
	fmt.Printf("wallet => %+v\n", wallet)
	wallet.Save("./wallet")
}

func loadWallet() {
	wallet := lib.LoadWallet("./wallet")
	fmt.Printf("wallet => %+v\n", wallet)
	priv := wallet.GetPrivKey("michael")
	fmt.Printf("michael's private key => %v\n", priv)
	pub := wallet.GetPubKey("michael")
	fmt.Printf("michael's public key => %v\n", pub)
	addr := wallet.GetAddress("michael")
	fmt.Printf("michael's address => %v\n", addr)
}
