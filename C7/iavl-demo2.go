package main

import (
	"fmt"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/iavl"
	"github.com/tendermint/tendermint/libs/db"
	"os"
)

func main() {
	os.RemoveAll("./acc.db")
	gdb, err := db.NewGoLevelDB("acc", ".")
	defer gdb.Close()
	if err != nil {
		panic(err)
	}
	tree := iavl.NewMutableTree(gdb, 128)
	fmt.Println("version", tree.Version())
	// 初次加载到工作区（工作区内才能修改）
	//Load()最后（新）版本的状态库到工作区
	//LoadVersion()载入指定版本的状态库到工作区。
	ver, err := tree.Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("tree version => %v\n", ver)

	codec := amino.NewCodec()
	//

	//操作
    fmt.Println("第一次写入")
	GetBalance("a", tree, codec)
	GetBalance("b", tree, codec)

	SetBalance("a", 1, tree, codec)
	SetBalance("b", 2, tree, codec)
	// 保存工作区
	hash, ver, err := tree.SaveVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(hash,"version", ver)

	// 读取
	GetBalance("a", tree, codec)
	GetBalance("b", tree, codec)



	fmt.Println("第二次写入")

	SetBalance("a", 11, tree, codec)
	SetBalance("a", 12, tree, codec)
	// 保存工作区
	hash, ver, err = tree.SaveVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(hash,"version", ver)

	// 读取
	GetBalance("a", tree, codec)
	GetBalance("b", tree, codec)
	fmt.Println("version", tree.Version())

	//读取指定版本
	GetBalanceVersioned("a", 1, tree, codec)
	GetBalanceVersioned("a", 2, tree, codec)


}

func SetBalance(name string, balance int, tree *iavl.MutableTree, codec *amino.Codec) {
	fmt.Printf("set %v's balance => %v\n", name, balance)
	bz, err := codec.MarshalBinaryBare(balance)
	if err != nil {
		panic(err)
	}
	tree.Set([]byte(name), bz)
}

func GetBalanceVersioned(name string, version int64, tree *iavl.MutableTree, codec *amino.Codec) int {
	var val int
	_, bz := tree.GetVersioned([]byte(name), version)
	if bz == nil {
		val = 0
	} else {
		err := codec.UnmarshalBinaryBare(bz, &val)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%v's balance@%v => %v\n", name, version, val)
	return val
}

func GetBalance(name string, tree *iavl.MutableTree, codec *amino.Codec) int {
	var val int
	_, bz := tree.Get([]byte(name))
	if bz == nil {
		val = 0
	} else {
		err := codec.UnmarshalBinaryBare(bz, &val)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%v's balance@workspace => %v\n", name, val)
	return val
}



func TestVersionedRandomTree() {


	d, _ := db.NewGoLevelDB("acc", ".")
	defer d.Close()

	tree := iavl.NewMutableTree(d, 100)
	versions := 50
	keysPerVersion := 30

	// Create a tree of size 1000 with 100 versions.
	for i := 1; i <= versions; i++ {
		for j := 0; j < keysPerVersion; j++ {
			k := []byte("s")
			v := []byte("a")
			tree.Set(k, v)
		}
		tree.SaveVersion()
	}

	fmt.Println(tree.Version())


	for i := 1; i < versions; i++ {
		tree.DeleteVersion(int64(i))
	}


}