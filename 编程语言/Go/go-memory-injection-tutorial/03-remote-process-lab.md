# 第 2 阶段：远程进程打开与内存读写

> 目标：对**自己启动的 target 进程**完成跨进程分配与读写。  
> 还不要执行远程代码。  
> 建议用时：2～4 小时。

---

## 1. 为什么这一阶段重要

很多“注入失败”其实死在执行之前：

- 打不开进程  
- 分不了内存  
- 写不进去  

先把 **Attach + Allocate + Write/Read** 练稳，后面 `CreateRemoteThread` 才好排错。

---

## 2. 先写一个 Target

### 2.1 目标程序要求

- 长期运行（不要立刻退出）  
- 定期打印 PID 与心跳  
- 自身不需要特殊权限  

`target/main.go` 示例：

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("[target] pid=%d started\n", os.Getpid())
	for i := 0; ; i++ {
		fmt.Printf("[target] heartbeat=%d time=%s\n", i, time.Now().Format(time.RFC3339))
		time.Sleep(2 * time.Second)
	}
}
```

编译：

```bash
cd target
go mod init example.com/target
go build -o target.exe
./target.exe
```

记下输出的 PID。

---

## 3. 远程读写程序要做什么

```text
解析命令行 PID
  → OpenProcess
  → VirtualAllocEx
  → WriteProcessMemory(字符串)
  → ReadProcessMemory
  → 打印读回结果
  → VirtualFreeEx（可选）
  → CloseHandle
```

---

## 4. 示例代码框架（教学版）

`phase2_remote_rw/main.go`：

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	memCommit  = 0x1000
	memReserve = 0x2000
	memRelease = 0x8000
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <pid>", os.Args[0])
	}
	pid64, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("invalid pid: %v", err)
	}
	pid := uint32(pid64)

	access := uint32(windows.PROCESS_VM_OPERATION |
		windows.PROCESS_VM_WRITE |
		windows.PROCESS_VM_READ |
		windows.PROCESS_QUERY_INFORMATION)

	hProcess, err := windows.OpenProcess(access, false, pid)
	if err != nil {
		log.Fatalf("OpenProcess failed: %v", err)
	}
	defer windows.CloseHandle(hProcess)
	fmt.Printf("[+] opened pid=%d\n", pid)

	size := uintptr(4096)
	remoteAddr, err := virtualAllocEx(hProcess, 0, size, memCommit|memReserve, windows.PAGE_READWRITE)
	if err != nil {
		log.Fatalf("VirtualAllocEx failed: %v", err)
	}
	fmt.Printf("[+] remote alloc: 0x%X\n", remoteAddr)

	defer func() {
		_ = virtualFreeEx(hProcess, remoteAddr, 0, memRelease)
	}()

	payload := []byte("hello-remote-lab\x00")
	var written uintptr
	if err := writeProcessMemory(hProcess, remoteAddr, payload, &written); err != nil {
		log.Fatalf("WriteProcessMemory failed: %v", err)
	}
	fmt.Printf("[+] wrote %d bytes\n", written)

	buf := make([]byte, len(payload))
	var read n uintptr
	if err := readProcessMemory(hProcess, remoteAddr, buf, &read); err != nil {
		log.Fatalf("ReadProcessMemory failed: %v", err)
	}
	fmt.Printf("[+] read %d bytes: %q\n", read, string(buf[:len(buf)-1]))

	fmt.Println("[+] phase2 remote rw OK")
}

// 下面用 LazyDLL 演示未完全封装时的调用方式。
// 你也可以优先查找 x/sys/windows 是否已有同名封装并改用它。

var (
	modKernel32            = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx     = modKernel32.NewProc("VirtualAllocEx")
	procVirtualFreeEx      = modKernel32.NewProc("VirtualFreeEx")
	procWriteProcessMemory = modKernel32.NewProc("WriteProcessMemory")
	procReadProcessMemory  = modKernel32.NewProc("ReadProcessMemory")
)

func virtualAllocEx(h windows.Handle, addr uintptr, size uintptr, allocType, protect uint32) (uintptr, error) {
	r1, _, e1 := procVirtualAllocEx.Call(
		uintptr(h),
		addr,
		size,
		uintptr(allocType),
		uintptr(protect),
	)
	if r1 == 0 {
		if e1 != nil {
			return 0, e1
		}
		return 0, fmt.Errorf("VirtualAllocEx returned NULL")
	}
	return r1, nil
}

func virtualFreeEx(h windows.Handle, addr uintptr, size uintptr, freeType uint32) error {
	r1, _, e1 := procVirtualFreeEx.Call(uintptr(h), addr, size, uintptr(freeType))
	if r1 == 0 {
		return e1
	}
	return nil
}

func writeProcessMemory(h windows.Handle, addr uintptr, data []byte, written *uintptr) error {
	r1, _, e1 := procWriteProcessMemory.Call(
		uintptr(h),
		addr,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(written)),
	)
	if r1 == 0 {
		return e1
	}
	return nil
}

func readProcessMemory(h windows.Handle, addr uintptr, buf []byte, read *uintptr) error {
	r1, _, e1 := procReadProcessMemory.Call(
		uintptr(h),
		addr,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(read)),
	)
	if r1 == 0 {
		return e1
	}
	return nil
}
```

### 使用方式

终端 1：

```bash
./target.exe
# 假设输出 pid=12345
```

终端 2：

```bash
go run . 12345
```

---

## 5. 用工具验证（强烈建议）

打开 Process Hacker / System Informer：

1. 找到 `target.exe`  
2. 查看 Memory 标签  
3. 找你刚分配的地址附近区域  
4. 确认 Protection 为读写，内容可见（若工具支持 hex 查看）  

把截图或记录写进笔记。

---

## 6. 练习任务

- [ ] 故意传错误 PID，观察报错  
- [ ] 去掉 `PROCESS_VM_WRITE`，确认写入失败  
- [ ] 写入 UTF-8 中文，读回打印  
- [ ] 分配两块内存，分别写入不同字符串  
- [ ] 在 target 退出后再注入，观察失败模式  
- [ ] 把错误信息翻译成人话记到 `notes/errors.md`

---

## 7. 权限与安全边界（必读）

本阶段只允许：

- 你自己编译的 `target.exe`  
- 你自己会话中启动的进程  

不要拿去实验：

- `lsass.exe`、系统服务、安全软件  
- 他人已经在跑的日常软件  

原因：容易触发系统保护，也容易越界到违规用途。

---

## 8. 完成后去哪里

远程读写稳定后：

→ [04-classic-injection.md](./04-classic-injection.md)  
或直接更推荐：

→ [05-dll-injection.md](./05-dll-injection.md)（结果更好观察）
