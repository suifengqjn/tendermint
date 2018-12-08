package lib

import (
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
)

type Wallet struct {
	Keys map[string]crypto.PrivKey
}

func NewWallet() *Wallet {
	return &Wallet{Keys: map[string]crypto.PrivKey{}}
}

func LoadWallet(wfn string) *Wallet {
	var wallet Wallet
	bz, err := ioutil.ReadFile(wfn)
	if err != nil {
		panic(err)
	}
	//err = json.Unmarshal(bz,&wallet)
	err = codec.UnmarshalJSON(bz, &wallet)
	if err != nil {
		panic(err)
	}
	return &wallet
}

func (wallet *Wallet) Save(wfn string) {
	//bz,err := json.Marshal(wallet)
	bz, err := codec.MarshalJSON(wallet)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(wfn, bz, 0644)
}

func (wallet *Wallet) GenPrivKey(label string) crypto.PrivKey {
	priv := kf.GenPrivKey()
	wallet.Keys[label] = priv
	return priv
}

func (wallet *Wallet) GetPrivKey(label string) crypto.PrivKey {
	return wallet.Keys[label]
}

func (wallet *Wallet) GetPubKey(label string) crypto.PubKey {
	priv := wallet.Keys[label]
	if priv == nil {
		panic("key not found")
	}
	return priv.PubKey()
}

func (wallet *Wallet) GetAddress(label string) crypto.Address {
	priv := wallet.Keys[label]
	if priv == nil {
		panic("key not found")
	}
	return priv.PubKey().Address()
}
