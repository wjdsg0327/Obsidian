### 场景 1：基础操作（内存模式，测试利器）

不需要真的在硬盘创建文件夹，代码运行完内存自动释放

```go
package main

import (
	"fmt"
	"github.com/spf13/afero"
)

func main() {
	// 初始化一个内存文件系统
	appFS := afero.NewMemMapFs()

	// 创建文件
	afero.WriteFile(appFS, "test.txt", []byte("hello afero"), 0644)

	// 读取文件
	content, _ := afero.ReadFile(appFS, "test.txt")
	fmt.Println(string(content)) // hello afero
}
```

### 场景 2：只读文件系统（保护配置文件）

如果你想确保某些敏感目录（如 `/etc/config`）在程序运行中绝对不会被意外修改。

```go
func main() {
	base := afero.NewOsFs() // 本地硬盘
	// 包装成只读模式
	roFS := afero.NewReadOnlyFs(base)

	err := roFS.Remove("config.yaml")
	if err != nil {
		fmt.Println("报错了，不能删除只读文件:", err)
	}
}
```

### 场景 3：限制根目录（Chroot）

类似于 Linux 的 `chroot`，让你的程序只能看到某个特定文件夹下的内容，增加安全性。

```go
package main

import (
	"github.com/spf13/afero"
)

func main() {
	base := afero.NewOsFs()
	// 把 /tmp/my-app 映射为该 FS 的根目录 "/"
	restrictedFS := afero.NewBasePathFs(base, "./tmp")

	// 此时创建 "data.log"，实际路径是 "/tmp/my-app/data.log"
	afero.WriteFile(restrictedFS, "data.log", []byte("log data"), 0644)
}

```

### 场景 4：层叠文件系统（Overlay）

经典的“覆盖”逻辑：先从内存找，找不到再去硬盘找。常用于“默认配置 + 用户自定义配置”的场景。

```go
func main() {
	base := afero.NewOsFs()      // 硬盘（底层）
	layer := afero.NewMemMapFs() // 内存（上层）

	// 组合起来
	compositeFS := afero.NewCopyOnWriteFs(base, layer)

	// 此时写入操作会写到 layer (内存)，读取操作会先看内存再看硬盘
	_ = afero.WriteFile(compositeFS, "patch.txt", []byte("override"), 0644)
}
```

### 场景 5：使用 Afero 的辅助工具 (Afero Struct)

Afero 提供了一个结构体封装，让调用更加丝滑（类似于 `os` 包的替代品）。

```go
func main() {
	fs := afero.NewMemMapFs()
	afs := &afero.Afero{Fs: fs}

	// 相比于 afero.WriteFile(fs, ...)，这种写法更简洁
	exists, _ := afs.Exists("some/path")
	if !exists {
		afs.MkdirAll("some/path", 0755)
	}
	
	afs.WriteReader("file.txt", someReader)
}
```