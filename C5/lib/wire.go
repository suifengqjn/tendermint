package lib

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

var codec = amino.NewCodec()

func init() {
	codec.RegisterInterface((*Payload)(nil), nil)
	codec.RegisterConcrete(&IssuePayload{}, "tx/issue", nil)
	codec.RegisterConcrete(&TransferPayload{}, "tx/transfer", nil)
	codec.RegisterInterface((*crypto.PubKey)(nil), nil)
	codec.RegisterConcrete(&kf.PubKeySecp256k1{}, "secp256k1/pubkey", nil)
	codec.RegisterInterface((*crypto.PrivKey)(nil), nil)
	codec.RegisterConcrete(&kf.PrivKeySecp256k1{}, "secp256k1/privkey", nil)
}

func MarshalBinary(o interface{}) ([]byte, error) {
	return codec.MarshalBinaryBare(o)
}

func UnmarshalBinary(bz []byte, ptr interface{}) error {
	return codec.UnmarshalBinaryBare(bz, ptr)
}

func MarshalJSON(o interface{}) ([]byte, error) {
	return codec.MarshalJSON(o)
}

func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return codec.UnmarshalJSON(bz, ptr)
}
