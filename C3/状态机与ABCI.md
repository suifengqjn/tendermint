# 状态机与ABCI

了解状态机与状态机复制的概念，掌握tendermint的ABCI应用实现方法。

## 运行预置代码

### 1、查看ABCI消息

在2#终端启动msg-dump.go：

```
$ go run msg-dump.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端提交任意交易：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x787878
```

返回2#终端查看屏幕输出。

### 2、交易检查 - CheckTx

在2#终端启动counter-checktx.go：

```
$ go run counter-checktx.go
```

在1#终端重新初始化启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端提交合规交易与不合规交易，观察输出的区别：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x01
~$ curl localhost:26657/broadcast_tx_commit?tx=0x78
```

### 3、交易执行 - DeliverTx

在2#终端启动counter-delivertx.go：

```
$ go run counter-delivertx.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端提交递增交易，观察输出中的信息：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0102
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0103
```

### 4、状态初始化 - InitChain

在2#终端启动counter-initchain.go：

```
$ go run counter-initchain.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端提交递增交易，观察输出中的信息：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
```

### 5、应用状态查询 - Query

在2#终端启动counter-query.go：

```
$ go run counter-query.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端执行查询，观察输出中的信息：

```
~$ curl localhost:26657/abci_query
```


### 6、应用状态的历史

在2#终端启动counter-history.go：

```
$ go run counter-history.go
```

在1#终端重新初始化启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端执行若干递增交易，并查询执行区块的历史

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0102
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0103
~$ curl localhost:26657/abci_query?height=1
```

### 7、应用/区块链握手机制

在2#终端启动counter-handshake.go：

```
$ go run counter-handshake.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端执行若干递增交易：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0102
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0103
```

在1#终端重新启动tendermint，观察是否存在交易重放。


### 8、应用状态的哈希值

在2#终端启动counter-hash.go：

```
$ go run counter-hash.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端执行递增交易，查看所在区块头的app_hash：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
```

### 9、应用状态持久化

在2#终端启动counter-persist.go：

```
$ go run counter-persist.go
```

在1#终端重新初始化并启动tendermint：

```
~$ tendermint node
```

在3#终端执行若干递增交易，记录最后的状态值：

```
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0101
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0102
~$ curl localhost:26657/broadcast_tx_commit?tx=0x0103
```

在3#终端查看状态文件内容：

```
~$ cat ~/repo/go/src/hubwiz.com/c3/counter.state
```

重新启动tendermint和abci应用，查看状态值是否与记录一致：

```
~$ curl localhost:26657/abci_query
```
