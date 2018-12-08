# 初识tendermint

了解tendermint应用开发流程，学习tendermint工具链的使用。

按以下步骤执行示例代码：

## 1、初始化节点

进入1#终端，执行以下命令初始化节点：

```
~$ tendermint init
```

## 2、启动节点

进入1#终端，执行以下命令启动节点：

```
~$ tendermint node
```

按ctrl+c终止节点运行，或者在另一个终端输入如下命令：

```
~$ pkill -9 tendermint
```


## 3、启动abci应用

进入2#终端，执行以下命令启动应用：

```
~$ cd ~/repo/go/src/hubwiz.com/c2
~/repo/go/src/hubwiz.com/c2$ go run mini-app.go
```

## 4、提交交易

进入3#终端，执行以下命令提交交易：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x12345678
```

## 5、查看区块信息

进入3#终端，执行以下命令查看指定高度的区块：

```
~$ curl localhost:26657/block?height=1
```
