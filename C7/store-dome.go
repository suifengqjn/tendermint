package main

import (
	"./lib"
	"fmt"
)

func main() {
	wallet := lib.LoadWallet("./wallet")
	fmt.Printf("wallet => %+v\n", wallet)
	store := lib.NewStore()
	fmt.Printf("store => %+v\n", store)

	addr := wallet.GetAddress("michael")
	store.SetBalance(addr, 3000)
	store.Commit()


	store.SetBalance(addr, 8000)
	store.Commit()

	ver := store.LastVersion
	fmt.Printf("latest commited versioin => %v\n", ver)

	store.SetBalance(addr, 12000)
	store.Commit()

	val, _ := store.GetBalance(addr)
	fmt.Printf("balance uncommited => %v\n", val)

	//val, _ = store.GetBalanceVersioned(addr, ver)
	//fmt.Printf("balance@%v => %v\n", ver, val)
	//
	//val, _ = store.GetBalanceVersioned(addr, ver-1)
	//fmt.Printf("balance@%v => %v\n", ver-1, val)

}
