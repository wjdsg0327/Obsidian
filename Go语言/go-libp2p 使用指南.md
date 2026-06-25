# go-libp2p 使用指南

> 项目地址：https://github.com/libp2p/go-libp2p
> 官方文档：https://docs.libp2p.io
> Go 包：https://pkg.go.dev/github.com/libp2p/go-libp2p

---

## 一、什么是 libp2p？

**libp2p** 是一个**模块化的 P2P 网络协议栈和库**，最初从 IPFS 项目中独立出来，现已成为众多去中心化项目的核心网络层。

**核心思想：**
- 将网络通信拆分为可插拔的模块（传输、发现、路由、安全、多路复用等）
- 应用可以按需选择模块，不必使用整个协议栈
- 天然支持多种传输协议（TCP、QUIC、WebSocket、WebRTC 等）

**go-libp2p** 是 libp2p 的 Go 语言实现，是整个 libp2p 生态中最重要的实现之一。

---

## 二、核心概念

### 2.1 Host（主机）
- libp2p 的核心对象，代表一个 P2P 节点
- 每个 Host 有一个唯一的 **PeerID**（基于公钥的多哈希）
- 包含：身份（密钥对）、监听地址、协议处理器、连接管理等

### 2.2 Multiaddr（多地址）
- libp2p 的地址格式，自描述地编码协议栈
- 例如：`/ip4/127.0.0.1/tcp/4001/p2p/QmPeerID`
- 可以叠加多种协议：IP → TCP → WebSocket → /p2p/PeerID

### 2.3 Stream（流）
- 节点间通信的基本单元，类似双向管道
- 基于协议 ID（如 `/echo/1.0.0`）进行路由
- 一个连接上可以多路复用多个 Stream

### 2.4 Peer Discovery（节点发现）
- mDNS（局域网发现）
- Kademlia DHT（分布式哈希表，广域网发现）
- Bootstrap nodes（引导节点）
- Rendezvous（约会点）

### 2.5 Protocol ID（协议标识）
- 自定义协议用字符串标识，如 `/myapp/chat/1.0.0`
- 通过 `SetStreamHandler` 注册处理函数

---

## 三、快速上手

### 3.1 环境要求
- Go >= 1.20

### 3.2 安装

```bash
mkdir -p ~/go-libp2p-demo && cd ~/go-libp2p-demo
go mod init mylibp2p
	go get github.com/libp2p/go-libp2p
```

### 3.3 最简示例：启动一个节点

```go
package main

import (
    "context"
    "fmt"

    "github.com/libp2p/go-libp2p"
)

func main() {
    // 创建一个默认配置的 libp2p 节点
    node, err := libp2p.New()
    if err != nil {
        panic(err)
    }
    defer node.Close()

    // 打印节点的监听地址
    fmt.Println("节点 ID:", node.ID())
    fmt.Println("监听地址:", node.Addrs())

    // 阻塞等待（实际使用中用 signal 优雅退出）
    select {}
}
```

运行后输出类似：
```
节点 ID: QmYo41GybZsFhR6g...
监听地址: [/ip4/127.0.0.1/tcp/57666 /ip6/::1/tcp/57667]
```

### 3.4 配置节点

```go
node, err := libp2p.New(
    // 指定监听地址和端口
    libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/4001"),
    // 禁用默认 ping 协议
    libp2p.Ping(false),
    // 使用 NAT 穿透
    libp2p.NATPortMap(),
    // 启用中继
    libp2p.EnableRelay(),
)
```

### 3.5 注册协议处理器（Stream Handler）

```go
import (
    "bufio"
    "fmt"
    "io"

    "github.com/libp2p/go-libp2p/core/network"
)

// 注册自定义协议
node.SetStreamHandler("/myapp/echo/1.0.0", func(s network.Stream) {
    defer s.Close()

    // 读取对方发来的消息
    buf := bufio.NewReader(s)
    str, _ := buf.ReadString('\n')
    fmt.Println("收到:", str)

    // 回复消息
    s.Write([]byte("echo: " + str))
})
```

### 3.6 连接其他节点并发送数据

```go
import (
    "context"
    "fmt"

    "github.com/libp2p/go-libp2p/core/peer"
    ma "github.com/multiformats/go-multiaddr"
)

// 解析对方的多地址
targetAddr, _ := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/4001/p2p/QmTargetPeerID")
addrInfo, _ := peer.AddrInfoFromP2pAddr(targetAddr)

// 连接
err := node.Connect(context.Background(), *addrInfo)

// 打开一个流
s, err := node.NewStream(context.Background(), addrInfo.ID, "/myapp/echo/1.0.0")

// 发送消息
s.Write([]byte("Hello libp2p!\n"))

// 读取回复
buf := bufio.NewReader(s)
reply, _ := buf.ReadString('\n')
fmt.Println("收到回复:", reply)
```

---

## 四、核心功能详解

### 4.1 Ping 协议

```go
import "github.com/libp2p/go-libp2p/p2p/protocol/ping"

// 创建 ping 服务
pingService := &ping.PingService{Host: node}
node.SetStreamHandler(ping.ID, pingService.PingHandler)

// 向其他节点 ping
ch := pingService.Ping(context.Background(), peerID)
res := <-ch
fmt.Println("RTT:", res.RTT)
```

### 4.2 Kademlia DHT 节点发现

```go
import dht "github.com/libp2p/go-libp2p-kad-dht"

// 创建 DHT 实例
kdht, err := dht.New(context.Background(), node)

// 引导 DHT（连接到引导节点）
err = kdht.Bootstrap(context.Background())

// 提供内容路由
kdht.Provide(context.Background(), contentCID, true)

// 查找提供某内容的节点
peers, err := kdht.FindProviders(context.Background(), contentCID)
```

### 4.3 mDNS 局域网发现

```go
import mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"

// 创建 mDNS 服务
mdnsService := mdns.NewMdnsService(node, "my-app-rendezvous", &mdnsNotifee{})

// 实现发现回调
type mdnsNotifee struct{}

func (n *mdnsNotifee) HandlePeerFound(pi peer.AddrInfo) {
    node.Connect(context.Background(), pi)
    fmt.Println("发现节点:", pi.ID)
}
```

### 4.4 GossipSub 消息广播

```go
import pubsub "github.com/libp2p/go-libp2p-pubsub"

// 创建 GossipSub 实例
ps, err := pubsub.NewGossipSub(context.Background(), node)

// 加入主题
topic, _ := ps.Join("my-topic")
sub, _ := topic.Subscribe()

// 发布消息
topic.Publish(context.Background(), []byte("hello everyone"))

// 接收消息
for {
    msg, _ := sub.Next(context.Background())
    fmt.Printf("来自 %s: %s\n", msg.ReceivedFrom, string(msg.Data))
}
```

---

## 五、完整示例：P2P Echo 服务

```go
package main

import (
    "bufio"
    "context"
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p/core/network"
    "github.com/libp2p/go-libp2p/core/peer"
    ma "github.com/multiformats/go-multiaddr"
)

const protocolID = "/myecho/1.0.0"

func main() {
    listenPort := flag.Int("port", 0, "监听端口（0=随机）")
    target := flag.String("target", "", "目标节点的多地址")
    flag.Parse()

    // 创建节点
    node, err := libp2p.New(
        libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *listenPort)),
    )
    if err != nil {
        panic(err)
    }
    defer node.Close()

    // 注册 echo 协议处理器
    node.SetStreamHandler(protocolID, func(s network.Stream) {
        defer s.Close()
        rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
        str, _ := rw.ReadString('\n')
        fmt.Printf("收到: %s", str)
        rw.WriteString("echo: " + str)
        rw.Flush()
    })

    // 打印节点信息
    addrInfo := peer.AddrInfo{ID: node.ID(), Addrs: node.Addrs()}
    addrs, _ := peer.AddrInfoToP2pAddrs(&addrInfo)
    fmt.Println("本节点地址:", addrs[0])

    // 如果指定了目标，连接并发送消息
    if *target != "" {
        maddr, _ := ma.NewMultiaddr(*target)
        pi, _ := peer.AddrInfoFromP2pAddr(maddr)

        if err := node.Connect(context.Background(), *pi); err != nil {
            panic(err)
        }

        s, err := node.NewStream(context.Background(), pi.ID, protocolID)
        if err != nil {
            panic(err)
        }

        rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
        rw.WriteString("Hello from echo client!\n")
        rw.Flush()

        reply, _ := rw.ReadString('\n')
        fmt.Printf("收到回复: %s", reply)
        s.Close()
    } else {
        // 等待信号退出
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        <-ch
        fmt.Println("\n正在关闭...")
    }
}
```

**使用方式：**

```bash
# 终端 1：启动监听节点
go run main.go -port 4001
# 输出: 本节点地址: /ip4/0.0.0.0/tcp/4001/p2p/QmXxx...

# 终端 2：连接并 echo
go run main.go -target "/ip4/127.0.0.1/tcp/4001/p2p/QmXxx..."
# 输出: 收到回复: echo: Hello from echo client!
```

---

## 六、实际使用案例（Notable Users）

go-libp2p 已被大量顶级区块链和去中心化项目采用：

| 项目 | 类型 | 说明 |
|------|------|------|
| **Kubo (IPFS)** | 去中心化存储 | Go 语言的 IPFS 实现，go-libp2p 的诞生地 |
| **Lotus (Filecoin)** | 去中心化存储 | Filecoin 协议的核心实现 |
| **Prysm (Ethereum)** | 区块链 | 以太坊 Beacon Chain 共识客户端 |
| **Celestia Node** | 模块化区块链 | 数据可用性节点实现 |
| **Polygon Edge** | Layer 2 | 以太坊兼容网络框架 |
| **Flow (Dapper Labs)** | 区块链 | 支持游戏、NFT 的区块链 |
| **Berty** | 即时通讯 | 开源、离线优先、端到端加密的 P2P 消息应用 |
| **Status** | 即时通讯/Web3 | 去中心化通讯 + 加密钱包 |
| **Swarm Bee** | 去中心化存储 | 以太坊 Swarm 网络客户端 |
| **Mina Protocol** | 区块链 | 恒定大小的零知识证明区块链 |
| **IOTA Wasp** | 智能合约 | IOTA 智能合约节点 |

---

## 七、典型应用场景

### 7.1 P2P 聊天应用
- 无需中心服务器，节点直接通信
- 结合 mDNS 实现局域网发现
- 结合 DHT 实现广域网通信

### 7.2 区块链网络层
- 节点发现和连接管理
- 交易和区块广播（GossipSub）
- 数据可用性采样

### 7.3 去中心化文件存储
- IPFS 的核心网络层
- 内容寻址和路由（DHT + Bitswap）
- 文件分发和复制

### 7.4 去中心化即时通讯
- 端到端加密消息传递
- 离线优先架构
- NAT 穿透和中继

### 7.5 IoT 设备通信
- 设备间直接通信
- 轻量级协议选择
- 自动 NAT 穿透

---

## 八、常用模块速查

```go
// 核心
"github.com/libp2p/go-libp2p"                    // 主入口
"github.com/libp2p/go-libp2p/core/host"           // Host 接口
"github.com/libp2p/go-libp2p/core/peer"           // PeerID、AddrInfo
"github.com/libp2p/go-libp2p/core/network"        // Stream、Conn
"github.com/libp2p/go-libp2p/core/protocol"       // Protocol ID

// 协议
"github.com/libp2p/go-libp2p/p2p/protocol/ping"   // Ping
"github.com/libp2p/go-libp2p/p2p/protocol/relay"   // 中继

// 发现
"github.com/libp2p/go-libp2p/p2p/discovery/mdns"  // mDNS 局域网发现
"github.com/libp2p/go-libp2p-kad-dht"             // Kademlia DHT

// 扩展
"github.com/libp2p/go-libp2p-pubsub"              // GossipSub 广播
"github.com/multiformats/go-multiaddr"             // 多地址解析
"github.com/libp2p/go-libp2p/p2p/net/swarm"       // Swarm 网络层

// 安全与多路复用（通常自动选择）
"github.com/libp2p/go-libp2p/p2p/security/noise"  // Noise 加密
"github.com/libp2p/go-libp2p/p2p/muxer/yamux"     // Yamux 多路复用
"github.com/libp2p/go-libp2p/p2p/transport/quic"  // QUIC 传输
"github.com/libp2p/go-libp2p/p2p/transport/tcp"   // TCP 传输
"github.com/libp2p/go-libp2p/p2p/transport/ws"    // WebSocket 传输
```

---

## 九、参考资源

- 官方文档：https://docs.libp2p.io
- Go 入门教程：https://libp2p.io/docs/getting-started-go
- 官方示例：https://github.com/libp2p/go-libp2p/tree/master/examples
- 讨论论坛：https://discuss.libp2p.io
- ProtoSchool 教程：https://proto.school/introduction-to-libp2p
- libp2p 规范：https://github.com/libp2p/specs

---

*最后更新：2025-06-25*
