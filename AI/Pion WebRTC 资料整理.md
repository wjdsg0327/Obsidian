---
title: Pion WebRTC 资料整理
date: 2026-07-04
tags: [WebRTC, Go, Pion, 实时音视频, SFU]
source:
  - https://github.com/pion/webrtc
  - https://pkg.go.dev/github.com/pion/webrtc/v4
  - https://github.com/pion/example-webrtc-applications
  - https://github.com/pion/ion-sfu
---

# Pion WebRTC 资料整理

## 1. 一句话说明

**Pion WebRTC** 是一个用 **Go 语言实现的 WebRTC API 库**，可以让 Go 程序直接建立 WebRTC 连接，收发音视频流、DataChannel 数据、RTP/RTCP 包，并用于构建实时音视频、直播、会议、远程控制、P2P 数据传输、媒体网关、SFU 等系统。

官方仓库：<https://github.com/pion/webrtc>

截至 2026-07-04 查询：

- GitHub：`pion/webrtc`
- 语言：Go
- License：MIT
- Star：约 16.6k
- Fork：约 1.8k
- 当前主版本：`v4`
- 项目描述：Pure Go implementation of the WebRTC API

## 2. WebRTC 是什么

WebRTC，全称 **Web Real-Time Communication**，是浏览器和应用之间进行实时通信的一套标准技术。它主要解决：

- 实时音视频通话
- 浏览器之间 P2P 传输
- 低延迟直播/互动直播
- 数据通道通信
- NAT 穿透
- 加密传输

WebRTC 底层涉及很多协议：

- ICE：网络连通性检查、NAT 穿透
- STUN：发现公网地址
- TURN：中继转发，解决无法直连的问题
- DTLS：密钥协商和加密
- SRTP：加密音视频 RTP 流
- SCTP：DataChannel 数据传输
- RTP/RTCP：音视频媒体包和控制反馈

浏览器原生支持 WebRTC，但如果服务端也想直接参与 WebRTC，比如做 SFU、媒体转发、录制、机器人视频流、网关，就需要服务端 WebRTC 实现。Pion 就是 Go 生态里非常重要的服务端 WebRTC 库。

## 3. Pion WebRTC 主要能力

### 3.1 PeerConnection API

Pion 实现了类似浏览器的 `RTCPeerConnection` API，支持：

- 创建 PeerConnection
- 交换 SDP Offer/Answer
- 添加音视频 Track
- 接收远端 Track
- ICE Candidate 交换
- 连接状态监听
- 协商/重协商
- WebRTC stats

### 3.2 音视频收发

可以用 Go 程序：

- 向浏览器发送视频/音频
- 从浏览器接收摄像头/麦克风流
- 读取 RTP 包做分析或转发
- 保存媒体到文件
- 将 RTP/RTMP/SIP 等协议桥接到 WebRTC

支持或常见集成：

- Opus
- PCM
- H264
- VP8
- VP9
- IVF
- Ogg
- H264 文件
- Matroska/WebM
- ffmpeg
- GStreamer
- x264/libvpx

### 3.3 DataChannel

Pion 支持 WebRTC DataChannel，可用于浏览器与 Go 服务之间传输任意数据。

特点：

- 有序/无序
- 可靠/不可靠
- 低延迟
- 可用于游戏、控制指令、文件片段、机器人控制、P2P 消息等

### 3.4 ICE / NAT 穿透

支持：

- ICE Agent
- ICE Restart
- Trickle ICE
- STUN
- TURN
- UDP/TCP/TLS/DTLS TURN
- mDNS candidates
- 单端口 ICE
- TCP ICE

实际部署时，WebRTC 能否连通往往取决于 ICE/STUN/TURN 配置。

### 3.5 直接 RTP/RTCP 操作

Pion 相比很多高级封装库更底层、更灵活。它允许直接访问：

- RTP 包
- RTCP 包
- Sender/Receiver Report
- NACK
- PLI/FIR
- TWCC
- 带宽估计

这适合做 SFU、媒体服务器、录制服务、网关服务。

### 3.6 纯 Go

官方强调 **No Cgo**，因此：

- 跨平台好编译
- 部署简单
- 适合服务器、边缘设备、嵌入式场景
- 支持 Windows、macOS、Linux、FreeBSD、iOS、Android、WASM 等

## 4. Pion 适合做什么

### 4.1 浏览器实时音视频服务

例如：

- 浏览器摄像头推流到 Go 服务
- Go 服务转发给其他浏览器
- 实时会议
- 远程监控
- 互动直播

### 4.2 SFU / 会议系统

SFU 是 Selective Forwarding Unit，选择性转发单元。

传统多人会议如果每个人都互相 P2P，连接数量会爆炸。SFU 的模式是：

```text
用户 A ─┐
用户 B ─┼──> SFU ───> 其他用户
用户 C ─┘
```

每个用户只上传一份媒体流到 SFU，SFU 再转发给其他人。

Pion 可用于自己实现 SFU，也有相关项目：

- `pion/example-webrtc-applications/sfu-ws`
- `pion/ion-sfu`

### 4.3 协议网关

Pion 可以作为 WebRTC 与其他协议之间的桥梁：

- RTMP → WebRTC
- RTP → WebRTC
- SIP → WebRTC
- WebRTC → RTMP/Twitch
- WebRTC → 文件录制
- 摄像头/设备视频 → 浏览器

### 4.4 边缘设备/机器人/IoT

因为 Go 编译部署简单，Pion 很适合：

- 摄像头设备推流到浏览器
- 机器人实时视频回传
- 远程控制指令 DataChannel
- 内网设备穿透通信

### 4.5 服务端媒体处理

例如：

- 收到浏览器视频后做 OpenCV/GoCV 分析
- 运动检测
- 截图
- 录制 WebM
- 实时特效
- 音视频统计分析

## 5. Pion 的基本工作流程

典型 WebRTC 建连流程：

```text
浏览器/客户端 A                         Go 服务端 Pion
       |                                     |
       | -------- SDP Offer --------------> |
       |                                     |
       | <------- SDP Answer ---------------|
       |                                     |
       | ---- ICE Candidates -------------> |
       | <--- ICE Candidates -------------- |
       |                                     |
       | ===== DTLS/SRTP/DataChannel ====== |
       |                                     |
```

注意：Pion 只负责 WebRTC 本身，不规定信令方式。

信令可以用：

- HTTP
- WebSocket
- SSE
- gRPC
- Redis/PubSub
- MQTT
- 手动复制 SDP
- 任意业务自己的协议

这也是 Pion 的特点：**底层灵活，但应用层要自己设计信令。**

## 6. 最小使用示例思路

Go 代码中一般这样开始：

```go
import "github.com/pion/webrtc/v4"
```

创建 PeerConnection：

```go
peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
    ICEServers: []webrtc.ICEServer{
        {URLs: []string{"stun:stun.l.google.com:19302"}},
    },
})
```

接收远端 DataChannel：

```go
peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
    d.OnMessage(func(msg webrtc.DataChannelMessage) {
        println(string(msg.Data))
    })
})
```

接收远端媒体 Track：

```go
peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
    for {
        pkt, _, err := track.ReadRTP()
        if err != nil {
            return
        }
        _ = pkt
    }
})
```

真实项目还需要：

- SDP Offer/Answer 交换
- ICE Candidate 交换
- 连接状态处理
- Track 管理
- TURN 服务器配置
- 断线重连和重协商

## 7. 官方示例分类

官方 `pion/webrtc/examples` 包含大量基础例子：

### Media API

- `reflect`：收到什么媒体就原样发回
- `play-from-disk`：从磁盘文件向浏览器发送视频
- `save-to-disk`：接收摄像头流并保存
- `broadcast`：一个上传者，多浏览器观看
- `rtp-forwarder`：转发 RTP 音视频流
- `rtp-to-webrtc`：RTP 包进入 Pion，再发给浏览器
- `simulcast`：多层视频流处理
- `swap-tracks`：动态切换媒体流
- `rtcp-processing`：处理 RTCP 统计和控制信息

### DataChannel API

- `data-channels`：浏览器和 Go 互发消息
- `data-channels-detach`：以更接近 Go 网络编程的方式使用 DataChannel
- `data-channels-flow-control`：DataChannel 流控
- `pion-to-pion`：两个 Go Pion 实例直接通信

### ICE / 网络

- `ice-restart`：网络切换后 ICE 重启
- `ice-single-port`：多个连接共用单 UDP 端口
- `ice-tcp`：TCP ICE
- `trickle-ice`：边收集候选边连接
- `vnet`：虚拟网络测试

## 8. 更完整的应用示例

仓库：<https://github.com/pion/example-webrtc-applications>

包含更接近真实项目的例子：

- `sfu-ws`：基于 WebSocket 信令的 SFU 会议系统
- `rtmp-to-webrtc`：RTMP 转 WebRTC
- `sip-to-webrtc`：SIP 转 WebRTC
- `sip-over-websocket-to-webrtc`
- `save-to-webm`：保存 WebM
- `gocv-receive`：接收视频并用 GoCV 处理
- `gocv-to-webrtc`：GoCV 摄像头/处理结果推给浏览器
- `gstreamer-send/receive`：和 GStreamer 集成
- `twitch`：WebRTC 转 RTMP 推 Twitch
- `snapshot`：视频帧截图并通过 HTTP 提供
- `ebiten-game`：WebRTC 游戏示例

## 9. Pion 与 ion-sfu

`ion-sfu` 是基于 Pion 生态的 SFU 项目，仓库：

<https://github.com/pion/ion-sfu>

特点：

- 聚焦 SFU 引擎
- 信令较少，方便嵌入到自己的服务
- 支持 json-rpc 信令
- 支持 gRPC 信令
- 可用 Docker 启动
- 支持实时媒体处理扩展

适合想快速搭会议/多人音视频转发的人参考。

## 10. Pion 的优点

### 优点

1. **Go 原生**
   - 部署简单，单二进制友好。

2. **非常灵活**
   - 不强绑定信令、业务模型、媒体架构。

3. **适合服务端 WebRTC**
   - SFU、网关、录制、转码、媒体处理都能做。

4. **底层能力强**
   - RTP/RTCP、ICE、DTLS、SRTP 等都能深入控制。

5. **示例多**
   - 官方示例覆盖了大量常见场景。

6. **社区成熟**
   - Star 高、项目活跃、文档和例子较多。

## 11. Pion 的缺点/注意点

### 11.1 学习曲线高

WebRTC 本身复杂，Pion 又比较底层。需要理解：

- SDP
- ICE Candidate
- STUN/TURN
- Track
- RTP/RTCP
- 编解码器
- NAT 穿透
- 信令流程

### 11.2 信令要自己做

Pion 不提供固定的信令服务器。你需要自己设计：

- 用户加入房间
- Offer/Answer 交换
- Candidate 交换
- 重连
- 鉴权
- 房间状态
- 多人媒体路由

### 11.3 生产部署要配 TURN

仅用 STUN 在很多网络下不够，尤其是：

- 公司网络
- 对称 NAT
- 移动网络
- 防火墙严格环境

生产环境通常需要部署 TURN 服务，如 coturn。

### 11.4 做 SFU 不只是转发

真正可用的 SFU 还需要处理：

- 带宽估计
- NACK/PLI
- Simulcast/SVC
- 订阅关系
- 房间管理
- 断线重连
- 流量控制
- 录制/旁路直播
- 权限控制

Pion 能提供底层能力，但完整产品还要大量业务工程。

## 12. 与其他方案对比

| 方案 | 定位 | 适合场景 |
|---|---|---|
| Pion WebRTC | Go WebRTC 库 | 自研服务端 WebRTC、媒体网关、SFU、嵌入式实时通信 |
| LiveKit | 完整实时音视频平台 | 快速搭会议/直播/房间系统 |
| mediasoup | Node.js/C++ SFU | 高性能自研会议系统，JS 生态 |
| Janus Gateway | C 语言 WebRTC 网关 | 插件式媒体网关、传统协议接入 |
| ion-sfu | Pion 生态 SFU | Go SFU 引擎/参考实现 |
| WebRTC 浏览器 API | 浏览器端能力 | 浏览器 P2P、前端音视频采集播放 |

如果目标是“快速上线会议系统”，LiveKit 更省事。  
如果目标是“学习 WebRTC 或深度控制媒体链路”，Pion 很合适。  
如果目标是“Go 服务端自己掌控 WebRTC”，Pion 是首选之一。

## 13. 推荐学习路线

### 第一步：先理解 WebRTC 基础

重点概念：

- Offer / Answer
- SDP
- ICE Candidate
- STUN / TURN
- PeerConnection
- Track
- DataChannel

推荐资料：

- WebRTC for the Curious：<https://webrtcforthecurious.com/>

### 第二步：跑官方例子

从简单到复杂：

1. `data-channels`
2. `play-from-disk`
3. `save-to-disk`
4. `broadcast`
5. `rtp-to-webrtc`
6. `sfu-ws`

### 第三步：自己做一个小项目

建议项目：

- 浏览器摄像头推流到 Go 服务
- Go 服务保存 WebM
- Go 服务转发给另一个浏览器
- DataChannel 控制消息

### 第四步：研究 SFU

看：

- `example-webrtc-applications/sfu-ws`
- `ion-sfu`
- LiveKit / mediasoup 的架构思路

## 14. 常见项目结构建议

一个基于 Pion 的简单服务端项目可以这样组织：

```text
my-webrtc-server/
├── cmd/server/main.go
├── internal/signaling/     # WebSocket/HTTP 信令
├── internal/room/          # 房间、用户、订阅关系
├── internal/webrtc/        # PeerConnection、Track 管理
├── internal/media/         # RTP/RTCP、录制、转发
├── web/                    # 前端测试页面
├── go.mod
└── README.md
```

## 15. 什么时候该用 Pion

适合：

- 你用 Go 写后端
- 需要服务端参与 WebRTC
- 要做音视频网关、SFU、录制、转发
- 需要直接处理 RTP/RTCP
- 希望部署简单、跨平台
- 想深入学习 WebRTC 底层

不太适合：

- 只想快速做普通视频会议，不想研究 WebRTC 细节
- 不想自己写信令、房间、重连、TURN 配置
- 需要开箱即用的完整产品后台

这种情况可以优先看 LiveKit、Jitsi、Janus、mediasoup 等更完整方案。

## 16. 总结

Pion WebRTC 可以理解成：

> **Go 语言里的 WebRTC 引擎。**

它不是一个完整的视频会议产品，而是构建实时音视频系统的底层库。它把浏览器 WebRTC 能力搬到了 Go 服务端，使 Go 程序可以直接参与实时音视频和数据通道通信。

如果老王要研究实时音视频、P2P、SFU、远程控制、摄像头推流、SIP/RTMP/WebRTC 网关，Pion 是非常值得看的项目。  
如果只是想快速搭成品系统，可以把 Pion 当底层技术参考，同时评估 LiveKit、ion-sfu、Janus、mediasoup 这类更完整框架。
