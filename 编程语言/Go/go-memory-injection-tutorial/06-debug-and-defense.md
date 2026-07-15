# 第 6 阶段：调试、排错与防御视角

> 目标：会排错，会观察，会从防御方理解注入行为。  
> 建议用时：与前面实验穿插进行，不必单独囤到最后。

---

## 1. 调试方法论

### 1.1 一次只改变一个变量

例如远程写失败时，不要同时改：

- 权限  
- 架构  
- 路径  
- payload  

先固定 target，再只改注入器一点。

### 1.2 强制记录现场

每次失败至少记录：

```text
时间
PID
注入器位数 / 目标位数
OpenProcess access 掩码
每个 API 的返回值与错误
DLL 绝对路径
远程分配地址
```

### 1.3 分层验证

```text
L1 本进程 VirtualAlloc 是否正常？
L2 OpenProcess 是否正常？
L3 VirtualAllocEx + Write/Read 是否正常？
L4 CreateRemoteThread 是否创建成功？
L5 DLL 是否真的出现在模块列表？
L6 DllMain 副作用是否出现？
```

哪一层失败，就只在那一层修。

---

## 2. 推荐观察工具怎么用

### Process Hacker / System Informer

关注：

- **Processes**：PID、完整性、命令行  
- **Memory**：新分配区、Protection  
- **Threads**：起始地址是否落在异常区域  
- **Modules**：是否出现陌生 DLL  
- **Handles**：是否打开了其他进程句柄  

### x64dbg

适合：

- 附加到 target  
- 在 `LoadLibraryW` 下断  
- 看参数是否为你的 DLL 路径  
- 观察返回的模块基址  

### API Monitor（可选）

观察注入器调用了哪些 API 及参数。

---

## 3. 常见错误码速查（学习向）

实际以 `FormatMessage` / 文档为准，下面是常见直觉：

| 情况 | 可能原因 |
|------|----------|
| Access is denied | 权限不足、受保护进程、安全软件 |
| Invalid handle | 句柄已关、Open 失败仍继续用 |
| The parameter is incorrect | flags/大小/地址不对 |
| 部分成功后目标退出 | payload/DllMain 崩溃 |

在 Go 里把 `error` 完整打印，不要吞掉。

---

## 4. Go 侧常见坑

1. **相对路径**：目标当前目录不是注入器目录  
2. **ANSI/UTF-16 混用**：`LoadLibraryW` 配了窄字符串  
3. **切片底层数组失效**：极少见但要有意识；保持 buffer 存活到 API 返回  
4. **忘记 defer CloseHandle**  
5. **把注入器本地地址当成远程地址传参**  
6. **32/64 混编**  

---

## 5. 防御视角：你在学什么检测点

### 5.1 行为链检测（经典）

短时间内出现：

```text
OpenProcess
 + VirtualAllocEx
 + WriteProcessMemory
 + CreateRemoteThread
```

并且源进程与目标进程关系异常 → 高可疑。

### 5.2 内存特征

- 匿名内存（不属于任何模块）且可执行  
- 线程起始地址落在匿名 RX/RWX 区  

### 5.3 模块特征

- 从临时目录加载的异常 DLL  
- 无签名、路径怪异、父子进程不匹配  

### 5.4 作为学习者的练习

写一页《我的检测备忘录》：

1. 若我是蓝队，会先看哪些工具  
2. 哪些事件能证明“发生了注入”而不是普通加载  
3. 误报可能来自哪里（调试器、IDE、合法热更新）  

---

## 6. 进阶技术只建立地图，不急着实现

| 名称 | 一句话 | 你现在要做什么 |
|------|--------|----------------|
| APC 注入 | 把回调排队到现有线程 | 读原理，画对比图 |
| 线程劫持 | 改线程上下文的 RIP | 读原理，理解风险 |
| 手动映射 | 自己实现 PE 加载 | 先精通正常 LoadLibrary |
| Process hollowing 等 | 更复杂进程替换手法 | 作为恶意软件分析词条学习 |

建议阅读顺序：

1. 先能稳定做 DLL 注入  
2. 再读恶意软件分析书的 injection 分类  
3. 最后才考虑实现第二种技术  

---

## 7. 实验卫生

- 使用虚拟机快照  
- 实验目录固定，如 `C:\Users\Public\injection-lab\`  
- 每次实验后记录是否需要回滚快照  
- 不要在宿主机对重要进程做实验  
- 不要关闭安全软件去“方便攻击测试”（学习阶段没必要）  

---

## 8. 阶段验收问题

能口头回答即合格：

1. 注入四步是什么？  
2. 为什么 DLL 路径要写到目标进程？  
3. 如何证明注入成功（至少 3 个证据）？  
4. 最常见的 5 个失败点？  
5. 蓝队可能如何发现 CreateRemoteThread 注入？  

---

## 9. 回到主线

- 资料总表：[resources.md](./resources.md)  
- 打卡清单：[lab-checklist.md](./lab-checklist.md)  
- 目录首页：[README.md](./README.md)
