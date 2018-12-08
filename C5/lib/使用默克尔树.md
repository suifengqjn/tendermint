# 代币案例：使用默克尔树

理解默克尔树的作用，学习如何在状态机中使用默克尔树。


目录文件组织：

- daemon.go: abci应用
- cli.go: 节点客户端
- lib: 公用代码目录
- merkle-hash.go：默克尔哈希测试代码
- merkle-proof.go：默克尔证据测试代码
- wallet：钱包文件

## 预置代码运行

### 1、计算默克尔哈希

在2#终端执行以下命令：

```
~/repo/go/src/hubwiz.com/c6$ go run merkle-hash.go
```

### 2、状态的默克尔证据

在2#终端执行以下命令：

```
~/repo/go/src/hubwiz.com/c6$ go run merkle-proof.go
```

### 3、ABCI应用

在2#终端启动ABCI应用：

```
~/repo/go/src/hubwiz.com/c6$ go run daemon.go
```

在1#终端重新初始化并启动tendermint

```
~$ tendermint unsafe_reset_all
~$ tendermint node
```

在3#终端执行客户端程序的子命令，例如：

发行代币：

```
~/repo/go/src/hubwiz.com/c6$ go run cli.go issue-tx
```

转账：

```
~/repo/go/src/hubwiz.com/c6$ go run cli.go transfer-tx
```

查询账户michael的余额：

```
~/repo/go/src/hubwiz.com/c6$ go run cli.go query michael
```

