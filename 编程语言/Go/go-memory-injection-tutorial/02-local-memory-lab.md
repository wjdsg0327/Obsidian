# 第 1 阶段：本进程内存实验

> 目标：在**当前进程**完成“分配 → 写入 → 读回 →（可选）改保护”。  
> 这是最安全的起点，不涉及跨进程。  
> 建议用时：半小时到 2 小时。

---

## 1. 本阶段你要完成什么

做一个 Go 小程序：

1. 调用 `VirtualAlloc` 分配 4KB 内存  
2. 写入一段 ASCII 字符串，例如 `hello-from-go-lab`  
3. 从同一地址读回并打印  
4. （可选）用 `VirtualProtect` 修改页属性  
5. 用 `VirtualFree` 释放  

成功标准：控制台输出分配地址，并校验读写一致。

---

## 2. 创建项目

```bash
mkdir -p phase1_local
cd phase1_local
go mod init example.com/phase1_local
go get golang.org/x/sys/windows
```

---

## 3. 示例代码（读写验证版）

把下面保存为 `main.go`：

```go
package main

import (
	"fmt"
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	memCommit  = 0x1000
	memReserve = 0x2000
	memRelease = 0x8000
)

func main() {
	size := uintptr(4096)

	addr, err := windows.VirtualAlloc(
		0,
		size,
		memCommit|memReserve,
		windows.PAGE_READWRITE,
	)
	if err != nil {
		log.Fatalf("VirtualAlloc failed: %v", err)
	}
	fmt.Printf("[+] allocated at: 0x%X\n", addr)

	// 确保失败路径也会释放；成功路径最后也释放
	defer func() {
		if freeErr := windows.VirtualFree(addr, 0, memRelease); freeErr != nil {
			log.Printf("VirtualFree failed: %v", freeErr)
		} else {
			fmt.Println("[+] memory freed")
		}
	}()

	payload := []byte("hello-from-go-lab\x00")
	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(payload))
	copy(dst, payload)
	fmt.Printf("[+] wrote %d bytes\n", len(payload))

	// 读回校验
	got := string(dst[:len(payload)-1]) // 去掉尾 0
	fmt.Printf("[+] read back: %q\n", got)
	if got != "hello-from-go-lab" {
		log.Fatal("data mismatch")
	}

	// 可选：修改保护为只读，验证保护确实生效
	var oldProtect uint32
	if err := windows.VirtualProtect(addr, size, windows.PAGE_READONLY, &oldProtect); err != nil {
		log.Fatalf("VirtualProtect failed: %v", err)
	}
	fmt.Printf("[+] protect changed, old=0x%X new=PAGE_READONLY\n", oldProtect)

	fmt.Println("[+] phase1 local lab OK")
}
```

运行：

```bash
go run .
```

预期类似：

```text
[+] allocated at: 0x...
[+] wrote 18 bytes
[+] read back: "hello-from-go-lab"
[+] protect changed, old=0x04 new=PAGE_READONLY
[+] memory freed
[+] phase1 local lab OK
```

---

## 4. 你要看懂的每一行

### 4.1 `VirtualAlloc` 参数

```text
VirtualAlloc(
  lpAddress = 0,          // 让系统挑地址
  dwSize    = 4096,       // 大小，通常按页对齐
  flAllocationType = MEM_COMMIT|MEM_RESERVE,
  flProtect = PAGE_READWRITE
)
```

### 4.2 为什么用 `unsafe.Slice`

`VirtualAlloc` 返回的是原始地址。  
要在 Go 里当字节数组用，需要把它解释成 `[]byte` 视图。  
这里只是教学演示；真实工程要更谨慎地处理生命周期与边界。

### 4.3 `VirtualProtect`

写数据和执行代码常常需要不同保护。  
本实验先改成 `PAGE_READONLY`，帮助你确认 API 可用。

---

## 5. 扩展练习（仍在本进程）

按顺序做，做完一项打勾：

- [ ] 把 size 改成 1 字节，观察系统实际提交是否仍按页对齐（可用工具看）  
- [ ] 分配后不写满整页，只写前 16 字节，读回确认  
- [ ] 打印 `oldProtect` 的数值，对照微软文档常量  
- [ ] 尝试对只读页再次 `copy` 写入，观察是否触发访问异常（可在子实验中用 recover 思路或单独小程序）  
- [ ] 查阅 `golang.org/x/sys/windows` 里还有哪些内存相关封装  

---

## 6. （选修）本进程“执行”概念实验

> 警告：执行自定义机器码涉及架构细节。入门阶段**可以跳过**。  
> 若做，请只执行你完全理解的极短 stub，并在虚拟机中进行。

更推荐的学习替代：

- 不执行机器码  
- 而是理解“可执行页”和“函数指针调用”在概念上的关系  
- 把真正的执行验证放到 **DLL 注入** 阶段（结果更清晰）

若你强行要做本地执行，学习重点应是：

1. 代码页需要执行权限  
2. 指令必须符合当前 CPU 架构  
3. 调用约定要匹配  
4. 错误的机器码会直接让进程崩溃  

---

## 7. 常见错误

| 错误 | 原因 | 处理 |
|------|------|------|
| `The parameter is incorrect` | size/类型/保护常量错误 | 对照文档检查 flags |
| 读回乱码 | 长度算错 / 没按字节解释 | 打印 hex |
| 程序崩溃 | 越界写 / 错误执行 | 先只做读写，别执行 |
| import 失败 | 未 `go get` | 安装 `x/sys/windows` |

---

## 8. 实验记录模板

复制到你的 `notes/observations.md`：

```markdown
## Phase1 记录

- 日期：
- Go 版本：
- OS：
- 分配地址：
- 写入内容：
- 读回内容：
- VirtualProtect 旧属性：
- 遇到的错误：
- 我的理解（3 句话）：
```

---

## 9. 完成后去哪里

确认本进程分配与读写无误后：

→ [03-remote-process-lab.md](./03-remote-process-lab.md)
