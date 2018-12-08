package lib

import (
	"errors"
	"github.com/tendermint/iavl"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/db"
)

type Store struct {
	tree        *iavl.MutableTree
	LastVersion int64
	LastHash    []byte
}

func NewStore() *Store {
	gdb, err := db.NewGoLevelDB("account", ".")
	if err != nil {
		panic(err)
	}
	tree := iavl.NewMutableTree(gdb, 128)
	ver, err := tree.Load()
	if err != nil {
		panic(err)
	}
	hash := tree.Hash()
	return &Store{tree, ver, hash}
}

func (store *Store) GetBalance(addr crypto.Address) (int, error) {
	_, bz := store.tree.Get(addr)
	if bz == nil {
		return 0, errors.New("account not found")
	}
	var val int
	err := codec.UnmarshalBinaryBare(bz, &val)
	if err != nil {
		return 0, errors.New("decode error")
	}
	return val, nil
}

func (store *Store) GetBalanceVersioned(addr crypto.Address, version int64) (int, error) {
	_, bz := store.tree.GetVersioned(addr, version)
	if bz == nil {
		return 0, errors.New("account not found on this version")
	}
	var val int
	err := codec.UnmarshalBinaryBare(bz, &val)
	if err != nil {
		return 0, errors.New("decode error")
	}
	return val, nil
}

func (store *Store) SetBalance(addr crypto.Address, value int) error {
	bz, err := codec.MarshalBinaryBare(value)
	if err != nil {
		return err
	}
	store.tree.Set(addr, bz)
	return nil
}

func (store *Store) Commit() {
	hash, ver, err := store.tree.SaveVersion()
	if err != nil {
		panic(err)
	}
	store.LastVersion = ver
	store.LastHash = hash
}
