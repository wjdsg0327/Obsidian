---
title: SS 链接转 Clash 配置
date: 2026-07-20
tags:
  - Clash
  - Shadowsocks
  - 科学上网
  - 配置
---

# SS 链接转 Clash 配置

## 原始链接

```text
ss://YWVzLTI1Ni1nY206ZzFnJWRha0QyQGNkVjhTS1kqQDQzLjE2MC4yNDcuNTA6MjMzMzY=
```

## 解码结果

```text
aes-256-gcm:g1g%dakD2@cdV8SKY*@43.160.247.50:23336
```

- 加密方式：`aes-256-gcm`
- 密码：`g1g%dakD2@cdV8SKY*`
- 服务器：`43.160.247.50`
- 端口：`23336`

## Clash 节点写法

```yaml
proxies:
  - name: ss-1
    type: ss
    server: 43.160.247.50
    port: 23336
    cipher: aes-256-gcm
    password: "g1g%dakD2@cdV8SKY*"
    udp: true
```

## 说明

- 这是单个 Shadowsocks 节点转成 Clash 节点配置后的写法。
- 如果后面有 `plugin` 参数，还需要补 `plugin` 和 `plugin-opts`。
- 如果你要做成完整 Clash 配置文件，可以把这段放进 `proxies:` 里，再配 `proxy-groups:` 和 `rules:`。
