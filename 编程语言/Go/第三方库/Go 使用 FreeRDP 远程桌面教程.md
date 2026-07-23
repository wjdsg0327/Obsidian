# Go 使用 FreeRDP 远程桌面教程

> 适合场景：用 Go 程序启动、管理、封装 RDP 远程桌面连接。  
> 核心项目：<https://github.com/FreeRDP/FreeRDP>

## 1. 先说清楚：FreeRDP 不是 Go 包

`github.com/FreeRDP/FreeRDP` 是 FreeRDP 官方仓库，它本体主要是 C/C++ 项目，提供：

- `xfreerdp`：Linux X11 下常用的 RDP 客户端命令行程序
- `wlfreerdp`：Wayland 环境下的 RDP 客户端
- `sdl-freerdp`：SDL 客户端
- `libfreerdp` / `winpr`：底层 C 库
- server、proxy、shadow 等扩展组件

所以 Go 里面不能这样直接使用：

```bash
go get github.com/FreeRDP/FreeRDP
```

也不能这样 import：

```go
import "github.com/FreeRDP/FreeRDP"
```

Go 调 FreeRDP 通常有三条路线：

| 路线 | 适合场景 | 推荐度 |
|---|---|---|
| Go 启动 `xfreerdp` / `wlfreerdp` 命令 | 做远程桌面启动器、堡垒机辅助工具、运维小工具 | 推荐 |
| 使用第三方 Go wrapper，例如 `github.com/moatasemgamal/gofreerdp` | 想少写一些参数拼接代码 | 可用，但要看维护情况 |
| cgo 直接链接 `libfreerdp` | 想深度嵌入 RDP 协议栈、自己处理事件循环、画面、输入 | 高难度，不建议新手直接上 |

实际项目里，最稳的是第一种：Go 程序负责配置、权限、日志、进程管理，真正的 RDP 协议交给 FreeRDP 官方客户端。

## 2. 安装 FreeRDP

### Ubuntu / Debian

```bash
sudo apt update
sudo apt install -y freerdp3-x11
```

有些系统包名可能还是 FreeRDP 2.x：

```bash
sudo apt install -y freerdp2-x11
```

检查是否安装成功：

```bash
xfreerdp /version
```

### Fedora

```bash
sudo dnf install -y freerdp
xfreerdp /version
```

### Arch Linux

```bash
sudo pacman -S freerdp
xfreerdp /version
```

### WSL 注意

如果在 WSL 里启动 `xfreerdp`，需要能显示图形界面：

- Windows 11 WSLg 通常可以直接显示 Linux GUI
- 旧环境需要额外配置 X Server，例如 VcXsrv
- 如果只是想连 Windows 远程桌面，Windows 本机自带 `mstsc.exe`，但那就不是 FreeRDP 了

## 3. FreeRDP 命令行基础

最简单连接：

```bash
xfreerdp /v:192.168.1.10 /u:Administrator /p:YourPassword
```

常用参数：

```bash
xfreerdp \
  /v:192.168.1.10:3389 \
  /u:Administrator \
  /p:YourPassword \
  /size:1280x720 \
  /cert:tofu \
  +clipboard \
  +dynamic-resolution
```

参数解释：

| 参数 | 作用 |
|---|---|
| `/v:host[:port]` | 目标 RDP 地址 |
| `/u:username` | 用户名 |
| `/p:password` | 密码 |
| `/d:domain` | 域 |
| `/size:1280x720` | 指定窗口分辨率 |
| `/f` | 全屏 |
| `+clipboard` | 开启剪贴板同步 |
| `+dynamic-resolution` | 窗口大小变化时动态调整远程分辨率 |
| `/cert:tofu` | 第一次信任证书，后续证书变更时拒绝 |
| `/cert:ignore` | 忽略证书校验，不推荐生产环境使用 |
| `/drive:name,path` | 映射本地目录到远程 Windows |
| `/sound` | 声音重定向 |
| `/microphone` | 麦克风重定向 |
| `/log-level:INFO` | 日志级别 |

证书参数建议：

- 测试环境可以用 `/cert:ignore`
- 自己长期使用的机器建议用 `/cert:tofu`
- 生产环境尽量固定证书指纹或证书校验策略

## 4. Go 直接启动 xfreerdp

先建一个最小项目：

```bash
mkdir go-freerdp-demo
cd go-freerdp-demo
go mod init example.com/go-freerdp-demo
```

创建 `main.go`：

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type RDPConfig struct {
	ClientPath string
	Host       string
	Port       int
	Username   string
	Password   string
	Domain     string
	Width      int
	Height     int
	Clipboard  bool
	DynamicRes bool
	CertMode   string // tofu, ignore, deny
}

func (c RDPConfig) Validate() error {
	if c.ClientPath == "" {
		return errors.New("ClientPath 不能为空，例如 xfreerdp")
	}
	if c.Host == "" {
		return errors.New("Host 不能为空")
	}
	if c.Username == "" {
		return errors.New("Username 不能为空")
	}
	if c.Password == "" {
		return errors.New("Password 不能为空")
	}
	return nil
}

func (c RDPConfig) Args() []string {
	port := c.Port
	if port == 0 {
		port = 3389
	}

	args := []string{
		fmt.Sprintf("/v:%s:%d", c.Host, port),
		fmt.Sprintf("/u:%s", c.Username),
		fmt.Sprintf("/p:%s", c.Password),
	}

	if c.Domain != "" {
		args = append(args, fmt.Sprintf("/d:%s", c.Domain))
	}
	if c.Width > 0 && c.Height > 0 {
		args = append(args, fmt.Sprintf("/size:%dx%d", c.Width, c.Height))
	}
	if c.CertMode != "" {
		args = append(args, fmt.Sprintf("/cert:%s", c.CertMode))
	}
	if c.Clipboard {
		args = append(args, "+clipboard")
	}
	if c.DynamicRes {
		args = append(args, "+dynamic-resolution")
	}

	args = append(args, "/log-level:INFO")
	return args
}

func StartRDP(ctx context.Context, cfg RDPConfig) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, cfg.ClientPath, cfg.Args()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Hour)
	defer cancel()

	cfg := RDPConfig{
		ClientPath: "xfreerdp",
		Host:       "192.168.1.10",
		Port:       3389,
		Username:   "Administrator",
		Password:   "YourPassword",
		Width:      1280,
		Height:     720,
		Clipboard:  true,
		DynamicRes: true,
		CertMode:   "tofu",
	}

	if err := StartRDP(ctx, cfg); err != nil {
		fmt.Fprintln(os.Stderr, "RDP 连接失败:", err)
		os.Exit(1)
	}
}
```

运行：

```bash
go run .
```

这个版本的思路很简单：Go 只做参数组装和进程管理，RDP 连接本身交给 `xfreerdp`。

## 5. 更安全地处理密码

上面的例子把密码放进命令行参数，适合本地测试，不适合长期使用。原因是命令行参数可能被 `ps`、任务管理器、日志系统看到。

更稳的做法有几种：

1. Go 程序从配置文件、环境变量、系统密钥库读取密码，不把密码写死到源码。
2. 不打印完整命令行，尤其不要打印 `/p:xxx`。
3. 使用 FreeRDP 的参数文件能力 `/args-from:file:xxx`，临时文件权限设为 `0600`，连接结束后删除。

示例：用临时 args 文件启动 FreeRDP。

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunWithArgsFile(ctx context.Context, client string, args []string) error {
	f, err := os.CreateTemp("", "freerdp-args-*.txt")
	if err != nil {
		return err
	}
	path := f.Name()
	defer os.Remove(path)

	if err := f.Chmod(0600); err != nil {
		f.Close()
		return err
	}

	content := strings.Join(args, "\n") + "\n"
	if _, err := f.WriteString(content); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, client, "/args-from:file:"+path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("启动 FreeRDP，参数文件:", path)
	return cmd.Run()
}
```

调用时：

```go
args := []string{
	"/v:192.168.1.10:3389",
	"/u:Administrator",
	"/p:YourPassword",
	"/cert:tofu",
	"+clipboard",
	"+dynamic-resolution",
}

err := RunWithArgsFile(context.Background(), "xfreerdp", args)
```

注意：临时文件依然会短暂存在，所以它不是绝对安全，只是比直接暴露在进程参数里更好。

## 6. 封装成一个可复用的小包

项目结构：

```text
go-freerdp-demo/
  go.mod
  main.go
  rdp/
    client.go
```

`rdp/client.go`：

```go
package rdp

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

type Client struct {
	Bin    string
	Stdout io.Writer
	Stderr io.Writer
}

type Session struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Domain     string
	Width      int
	Height     int
	Fullscreen bool
	Clipboard  bool
	CertMode   string
	DriveName  string
	DrivePath  string
}

func (c Client) Start(ctx context.Context, s Session) error {
	bin := c.Bin
	if bin == "" {
		bin = "xfreerdp"
	}

	args := buildArgs(s)
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stderr

	return cmd.Run()
}

func buildArgs(s Session) []string {
	port := s.Port
	if port == 0 {
		port = 3389
	}

	args := []string{
		fmt.Sprintf("/v:%s:%d", s.Host, port),
		fmt.Sprintf("/u:%s", s.Username),
		fmt.Sprintf("/p:%s", s.Password),
	}

	if s.Domain != "" {
		args = append(args, fmt.Sprintf("/d:%s", s.Domain))
	}
	if s.Fullscreen {
		args = append(args, "/f")
	} else if s.Width > 0 && s.Height > 0 {
		args = append(args, fmt.Sprintf("/size:%dx%d", s.Width, s.Height))
	}
	if s.Clipboard {
		args = append(args, "+clipboard")
	}
	if s.CertMode != "" {
		args = append(args, fmt.Sprintf("/cert:%s", s.CertMode))
	}
	if s.DriveName != "" && s.DrivePath != "" {
		args = append(args, fmt.Sprintf("/drive:%s,%s", s.DriveName, s.DrivePath))
	}

	return args
}
```

`main.go`：

```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"example.com/go-freerdp-demo/rdp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Hour)
	defer cancel()

	client := rdp.Client{
		Bin:    "xfreerdp",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	session := rdp.Session{
		Host:      "192.168.1.10",
		Username:  "Administrator",
		Password:  os.Getenv("RDP_PASSWORD"),
		Width:     1440,
		Height:    900,
		Clipboard: true,
		CertMode:  "tofu",
		DriveName: "share",
		DrivePath: "/home/wang/share",
	}

	if err := client.Start(ctx, session); err != nil {
		log.Fatal(err)
	}
}
```

运行：

```bash
export RDP_PASSWORD='你的密码'
go run .
```

## 7. 适合运维工具的功能设计

如果你要做一个自己的远程桌面启动器，建议不要只封装一条命令，而是设计成这样：

```text
Go 程序
  ├─ 读取机器清单：服务器名、IP、端口、用户名
  ├─ 从密钥库或环境变量读取密码
  ├─ 生成 FreeRDP 参数
  ├─ 启动 xfreerdp 进程
  ├─ 捕获退出码和日志
  └─ 根据失败原因给出中文提示
```

机器清单示例 `servers.yaml`：

```yaml
servers:
  - name: win-dev
    host: 192.168.1.10
    port: 3389
    username: Administrator
    cert_mode: tofu
    width: 1440
    height: 900

  - name: win-prod
    host: 10.0.0.12
    port: 3389
    username: ops
    cert_mode: tofu
    fullscreen: true
```

命令行设计：

```bash
go-rdp list
go-rdp connect win-dev
go-rdp connect win-prod --fullscreen
go-rdp connect win-dev --drive share=/home/wang/share
```

这类工具的价值不在于重新实现 RDP，而是把常用连接参数、机器清单、安全策略、日志记录统一起来。

## 8. 使用第三方 Go wrapper

有人做过 Go wrapper：

```bash
go get github.com/moatasemgamal/gofreerdp
```

它的思路也是封装 FreeRDP 3.x 命令行，不是纯 Go RDP 协议实现。

示例风格大概是：

```go
package main

import "github.com/moatasemgamal/gofreerdp"

func main() {
	freeRDP, err := gofreerdp.Init(gofreerdp.DisplayServer_Xorg)
	if err != nil {
		panic(err)
	}

	conf := &gofreerdp.RDPConfig{
		Addr:     "192.168.1.10:3389",
		Username: "Administrator",
		Password: "YourPassword",
	}

	freeRDP.SetConfig(conf)
	freeRDP.Run()
}
```

使用这类 wrapper 前，要检查几点：

- 是否支持你机器上的 FreeRDP 版本，例如 2.x / 3.x
- 是否支持你需要的参数，比如剪贴板、目录映射、网关、证书策略
- 是否会把密码打印到日志里
- 项目是否还在维护

如果 wrapper 功能不够，直接用 `os/exec` 自己封装反而更清楚。

## 9. 纯 Go / WebAssembly 路线：grdpwasm

你提到的 `https://github.com/nakagami/grdpwasm` 是另一条路线，它和 FreeRDP 没有直接关系。

它的结构大概是：

```text
浏览器 WASM 客户端
  └─ WebSocket
      └─ Go proxy
          └─ TCP 3389
              └─ Windows / xrdp / 其他 RDP Server
```

浏览器不能直接打开原始 TCP 连接，所以 `grdpwasm` 做了两件事：

- 前端：Go 编译成 WebAssembly，在浏览器 canvas 里显示远程桌面画面，处理键盘、鼠标、音频。
- 后端：Go proxy 把浏览器 WebSocket 流量转发到 RDP Server 的 TCP 端口。

它底层依赖的是 `github.com/nakagami/grdp`，这是一个纯 Go RDP 协议客户端，不依赖 FreeRDP 的 `xfreerdp` 命令行，也不是 `libfreerdp` 的 cgo binding。

### grdpwasm 适合什么

适合：

- 想做一个网页 RDP 客户端。
- 想在浏览器里打开 Windows 远程桌面，不让用户本地安装客户端。
- 想研究纯 Go RDP 协议实现。
- 想做内网运维面板、远程桌面入口、轻量堡垒机原型。

不太适合：

- 只想从 Go 程序里启动一个远程桌面窗口。
- 想要 FreeRDP 那种成熟、完整、长期打磨的桌面客户端能力。
- 想暴露到公网直接给多人使用，但又不准备做认证、TLS、审计和权限隔离。

### 和 FreeRDP 路线的区别

| 对比项 | Go + xfreerdp | grdpwasm / grdp |
|---|---|---|
| 底层实现 | FreeRDP 官方客户端 | 纯 Go RDP 协议实现 |
| 是否依赖 FreeRDP | 依赖 `xfreerdp` | 不依赖 |
| 使用方式 | Go 启动本地图形客户端 | 浏览器打开网页客户端 |
| 画面承载 | 本地 X11 / Wayland 窗口 | 浏览器 canvas |
| 网络方式 | 客户端直接连 RDP Server | 浏览器经 WebSocket proxy 转 TCP |
| 成熟度 | 更成熟 | 更适合研究和二次开发 |
| 适合方向 | 运维工具、启动器、本机客户端 | Web RDP、浏览器远程桌面、协议研究 |

### grdpwasm 的优点

- 纯 Go 技术栈，适合 Go 项目二次开发。
- 可以跑在浏览器里，不要求用户安装 RDP 客户端。
- 支持键盘、鼠标、滚轮输入。
- README 提到支持远程音频，通过浏览器 Web Audio API 播放。
- `grdp` 包暴露了 `RdpClient`、`OnBitmap`、`OnAudio`、`KeyDown`、`MouseMove`、`MouseDown`、`MouseWheel` 等 API，说明它不是简单命令封装，而是真在 Go 里实现和处理 RDP 数据。

### grdpwasm 的风险点

- 许可证是 GPLv3，商用闭源项目要特别注意。
- 浏览器到 proxy 之间会传账号密码，必须上 HTTPS/WSS。
- proxy 默认如果不加认证，不能直接暴露公网。
- RDP 协议很复杂，兼容性通常不如 FreeRDP 官方客户端。
- 如果要做生产级堡垒机，还需要补认证、会话隔离、审计、限流、权限控制、日志脱敏。

### 快速运行思路

仓库 README 的构建方式大致是：

```bash
git clone https://github.com/nakagami/grdpwasm.git
cd grdpwasm
make all
make serve
```

构建后会生成：

| 文件 | 作用 |
|---|---|
| `static/main.wasm` | 跑在浏览器里的 Go WASM 客户端 |
| `static/wasm_exec.js` | Go WebAssembly 运行时 JS |
| `proxy/proxy` | WebSocket-to-TCP proxy 和静态文件服务 |

启动后打开：

```text
http://localhost:8080
```

然后填 RDP Host、Port、Domain、User、Password、Width、Height 连接。

### 我的判断

如果你的目标是“用 Go 做一个网页远程桌面客户端”，`grdpwasm` 比 FreeRDP 命令行封装更贴近目标。

如果你的目标是“Go 程序帮我打开远程桌面，稳定能用”，还是优先用 `xfreerdp`。

如果你的目标是“研究 RDP 协议，拿到画面帧、音频、鼠标键盘事件，自己做产品”，可以重点看 `nakagami/grdp` 和 `grdpwasm`，但要提前接受兼容性、安全和 GPLv3 许可证成本。

## 10. cgo 直接调用 libfreerdp 的思路

这条路线适合想把 RDP 客户端深度嵌进 Go 程序的人，比如：

- 自己接收远程画面帧
- 自己处理键盘、鼠标输入
- 自己做远程桌面网关或录屏
- 想接 FreeRDP 的底层通道能力

但这不是简单的 Go import。你需要：

1. 安装 FreeRDP 开发包：头文件和动态库。
2. 用 `pkg-config` 找到 `freerdp3`、`winpr3` 之类的编译参数。
3. 用 cgo 写 C 桥接层。
4. 在 C 侧处理 FreeRDP 的回调、事件循环、内存释放。
5. 小心 Go 和 C 之间的指针传递规则。

开发包安装示例：

```bash
sudo apt install -y libfreerdp-dev libwinpr-dev pkg-config
pkg-config --libs --cflags freerdp3
```

如果系统还是 FreeRDP 2.x，名字可能是：

```bash
pkg-config --libs --cflags freerdp2
```

cgo 文件通常长这样：

```go
package freerdp

/*
#cgo pkg-config: freerdp3 winpr3
#include <freerdp/freerdp.h>
#include <freerdp/settings.h>
*/
import "C"

func Version() string {
	return C.GoString(C.freerdp_get_version_string())
}
```

这个例子只适合说明“怎么链接到 FreeRDP C 库”。真正建立连接还要设置 `freerdp_context_new`、`freerdp_connect`、事件循环、图形回调、输入回调等，复杂度明显高于命令行封装。

我的建议：

- 如果只是启动远程桌面窗口：用 `os/exec` 调 `xfreerdp`。
- 如果要做堡垒机、运维面板：还是先用 `os/exec`，把连接管理做好。
- 如果要拿到画面帧、做录制、做嵌入式客户端：再研究 cgo + `libfreerdp`。

## 11. 常见问题

### 连接时报证书错误

测试时可以先用：

```bash
/cert:ignore
```

长期使用建议改成：

```bash
/cert:tofu
```

`tofu` 是 Trust On First Use，第一次连接信任，之后如果证书变了会提醒或拒绝。

### Windows 远程桌面连不上

检查目标 Windows：

- 是否开启“远程桌面”
- 防火墙是否允许 3389
- 用户是否有远程登录权限
- Windows 家庭版通常不能作为 RDP Server
- 账号是否允许密码登录，空密码一般不行

### FreeRDP 版本不同导致参数不兼容

先看版本：

```bash
xfreerdp /version
xfreerdp /help
```

Go 程序里也可以启动时检查：

```go
out, err := exec.Command("xfreerdp", "/version").CombinedOutput()
if err != nil {
	return err
}
fmt.Println(string(out))
```

### 在服务器上运行没有图形界面

`xfreerdp` 是图形客户端，需要 X11 或 Wayland。纯服务器环境没有显示器时，可以考虑：

- 用带 GUI 的桌面环境运行
- 用 Xvfb 虚拟显示
- 改用 FreeRDP 的无画面/测试参数做认证或压力测试
- 如果只是管理 Windows，考虑 WinRM、SSH、PowerShell Remoting，而不是 RDP

### 中文输入或键盘布局异常

可以查看 FreeRDP 支持的键盘布局：

```bash
xfreerdp /list:kbd
```

连接时指定键盘布局：

```bash
xfreerdp /v:192.168.1.10 /u:Administrator /p:YourPassword /kbd:layout:0x00000409
```

`0x00000409` 是英文美国键盘，中文环境需要按实际情况调整。

## 12. 推荐落地方案

如果你现在要用 Go 做一个 FreeRDP 工具，推荐从这个版本开始：

```text
第一阶段：Go + os/exec + xfreerdp
  - 支持机器清单
  - 支持剪贴板、分辨率、证书策略
  - 支持日志和错误提示
  - 密码从环境变量或密钥库读取

第二阶段：做成 CLI 工具
  - go-rdp list
  - go-rdp connect <server>
  - go-rdp test <server>

第三阶段：如果想做浏览器 RDP，研究 grdpwasm / grdp

第四阶段：如果真需要 FreeRDP 底层能力，再研究 cgo + libfreerdp
```

不要一开始就尝试用 Go 重写 RDP 协议，也不要急着 cgo 直连 `libfreerdp`。FreeRDP 官方客户端已经处理了大量协议细节，Go 更适合站在它上面做配置管理、流程控制和工具化封装。

## 参考资料

- FreeRDP 官方仓库：<https://github.com/FreeRDP/FreeRDP>
- FreeRDP 官网：<https://www.freerdp.com/>
- FreeRDP API 文档：<https://pub.freerdp.com/api/>
- xfreerdp man page：<https://www.mankier.com/1/xfreerdp>
- gofreerdp wrapper：<https://pkg.go.dev/github.com/moatasemgamal/gofreerdp>
