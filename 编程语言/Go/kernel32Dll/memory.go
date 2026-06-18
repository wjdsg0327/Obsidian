//go:build windows

package kernel32Dll

import "unsafe"

var (
	procVirtualAlloc       = kernel32.NewProc("VirtualAlloc")
	procVirtualFree        = kernel32.NewProc("VirtualFree")
	procVirtualProtect     = kernel32.NewProc("VirtualProtect")
	procVirtualQuery       = kernel32.NewProc("VirtualQuery")
	procVirtualQueryEx     = kernel32.NewProc("VirtualQueryEx")
	procReadProcessMemory  = kernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory = kernel32.NewProc("WriteProcessMemory")
	procGetProcessHeap     = kernel32.NewProc("GetProcessHeap")
	procHeapCreate         = kernel32.NewProc("HeapCreate")
	procHeapAlloc          = kernel32.NewProc("HeapAlloc")
	procHeapFree           = kernel32.NewProc("HeapFree")
	procHeapDestroy        = kernel32.NewProc("HeapDestroy")
	procGlobalAlloc        = kernel32.NewProc("GlobalAlloc")
	procGlobalLock         = kernel32.NewProc("GlobalLock")
	procGlobalUnlock       = kernel32.NewProc("GlobalUnlock")
	procGlobalFree         = kernel32.NewProc("GlobalFree")
	procGlobalSize         = kernel32.NewProc("GlobalSize")
)

const (
	MEM_COMMIT  = 0x00001000
	MEM_RESERVE = 0x00002000
	MEM_RELEASE = 0x00008000
	MEM_FREE    = 0x00010000
	MEM_PRIVATE = 0x00020000
	MEM_MAPPED  = 0x00040000
	MEM_IMAGE   = 0x01000000

	PAGE_NOACCESS          = 0x01
	PAGE_READONLY          = 0x02
	PAGE_READWRITE         = 0x04
	PAGE_WRITECOPY         = 0x08
	PAGE_EXECUTE           = 0x10
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_EXECUTE_WRITECOPY = 0x80
	PAGE_GUARD             = 0x100
	PAGE_NOCACHE           = 0x200
	PAGE_WRITECOMBINE      = 0x400

	HEAP_NO_SERIALIZE = 0x00000001
	HEAP_ZERO_MEMORY  = 0x00000008

	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
)

// MEMORY_BASIC_INFORMATION 表示 VirtualQuery 返回的内存区域信息。
type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	PartitionId       uint16
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

// VirtualAlloc 提交或保留虚拟内存。
// 参数address：期望分配地址，0 表示由系统选择。
// 参数size：分配大小，单位字节。
// 参数allocationType：分配类型，例如 MEM_COMMIT 或 MEM_RESERVE。
// 参数protect：内存保护属性，例如 PAGE_READWRITE。
// 返回值：返回分配到的内存地址；返回 0 表示分配失败。
func VirtualAlloc(address uintptr, size uintptr, allocationType uint32, protect uint32) uintptr {
	ret, _, _ := procVirtualAlloc.Call(address, size, uintptr(allocationType), uintptr(protect))
	return ret
}

// VirtualFree 释放或取消提交虚拟内存。
// 参数address：要释放的内存地址。
// 参数size：释放大小；MEM_RELEASE 时必须为 0。
// 参数freeType：释放类型，例如 MEM_RELEASE。
// 返回值：true 表示释放成功，false 表示释放失败。
func VirtualFree(address uintptr, size uintptr, freeType uint32) bool {
	ret, _, _ := procVirtualFree.Call(address, size, uintptr(freeType))
	return ret != 0
}

// VirtualProtect 修改虚拟内存保护属性。
// 参数address：内存起始地址。
// 参数size：内存大小，单位字节。
// 参数newProtect：新的内存保护属性。
// 返回值：返回旧保护属性和是否修改成功；true 表示 oldProtect 有效。
func VirtualProtect(address uintptr, size uintptr, newProtect uint32) (oldProtect uint32, ok bool) {
	ret, _, _ := procVirtualProtect.Call(address, size, uintptr(newProtect), uintptr(unsafe.Pointer(&oldProtect)))
	return oldProtect, ret != 0
}

// VirtualQuery 查询虚拟内存区域信息。
// 参数address：要查询的内存地址。
// 返回值：返回内存区域信息和是否查询成功；true 表示结构体内容有效。
func VirtualQuery(address uintptr) (MEMORY_BASIC_INFORMATION, bool) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, _ := procVirtualQuery.Call(address, uintptr(unsafe.Pointer(&mbi)), unsafe.Sizeof(mbi))
	return mbi, ret != 0
}

// VirtualQueryEx 查询指定进程的虚拟内存区域信息。
// 参数process：目标进程句柄，需要 PROCESS_QUERY_INFORMATION 或 PROCESS_QUERY_LIMITED_INFORMATION 权限。
// 参数address：目标进程中要查询的内存地址。
// 返回值：返回内存区域信息和是否查询成功；true 表示结构体内容有效，false 表示查询失败或地址超出范围。
func VirtualQueryEx(process uintptr, address uintptr) (MEMORY_BASIC_INFORMATION, bool) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, _ := procVirtualQueryEx.Call(process, address, uintptr(unsafe.Pointer(&mbi)), unsafe.Sizeof(mbi))
	return mbi, ret != 0
}

// ReadProcessMemory 读取指定进程内存。
// 参数process：目标进程句柄，需要 PROCESS_VM_READ 权限。
// 参数baseAddress：目标进程中要读取的内存地址。
// 参数buffer：接收读取数据的字节切片。
// 返回值：返回实际读取字节数和是否读取成功；true 表示 buffer 中前 n 个字节有效，false 表示读取失败。
func ReadProcessMemory(process uintptr, baseAddress uintptr, buffer []byte) (n uintptr, ok bool) {
	var bytesRead uintptr
	var bufPtr uintptr
	if len(buffer) > 0 {
		bufPtr = uintptr(unsafe.Pointer(&buffer[0]))
	}
	ret, _, _ := procReadProcessMemory.Call(process, baseAddress, bufPtr, uintptr(len(buffer)), uintptr(unsafe.Pointer(&bytesRead)))
	return bytesRead, ret != 0
}

// WriteProcessMemory 写入指定进程内存。
// 参数process：目标进程句柄，需要 PROCESS_VM_WRITE 和 PROCESS_VM_OPERATION 权限。
// 参数baseAddress：目标进程中要写入的内存地址。
// 参数buffer：要写入的字节切片。
// 返回值：返回实际写入字节数和是否写入成功；true 表示写入调用成功，false 表示写入失败。
func WriteProcessMemory(process uintptr, baseAddress uintptr, buffer []byte) (n uintptr, ok bool) {
	var bytesWritten uintptr
	var bufPtr uintptr
	if len(buffer) > 0 {
		bufPtr = uintptr(unsafe.Pointer(&buffer[0]))
	}
	ret, _, _ := procWriteProcessMemory.Call(process, baseAddress, bufPtr, uintptr(len(buffer)), uintptr(unsafe.Pointer(&bytesWritten)))
	return bytesWritten, ret != 0
}

// GetProcessHeap 获取当前进程默认堆。
// 返回值：返回默认进程堆句柄；返回 0 表示获取失败。
func GetProcessHeap() uintptr {
	ret, _, _ := procGetProcessHeap.Call()
	return ret
}

// HeapCreate 创建私有堆。
// 参数options：堆选项标志。
// 参数initialSize：初始提交大小。
// 参数maximumSize：最大堆大小，0 表示可增长。
// 返回值：返回堆句柄；返回 0 表示创建失败。
func HeapCreate(options uint32, initialSize, maximumSize uintptr) uintptr {
	ret, _, _ := procHeapCreate.Call(uintptr(options), initialSize, maximumSize)
	return ret
}

// HeapAlloc 从堆中分配内存。
// 参数heap：堆句柄。
// 参数flags：分配标志，例如 HEAP_ZERO_MEMORY。
// 参数size：分配大小，单位字节。
// 返回值：返回分配到的内存地址；返回 0 表示分配失败。
func HeapAlloc(heap uintptr, flags uint32, size uintptr) uintptr {
	ret, _, _ := procHeapAlloc.Call(heap, uintptr(flags), size)
	return ret
}

// HeapFree 释放堆内存。
// 参数heap：堆句柄。
// 参数flags：释放标志，通常为 0。
// 参数memory：要释放的内存地址。
// 返回值：true 表示释放成功，false 表示释放失败。
func HeapFree(heap uintptr, flags uint32, memory uintptr) bool {
	ret, _, _ := procHeapFree.Call(heap, uintptr(flags), memory)
	return ret != 0
}

// HeapDestroy 销毁私有堆。
// 参数heap：堆句柄。
// 返回值：true 表示堆销毁成功，false 表示销毁失败。
func HeapDestroy(heap uintptr) bool {
	ret, _, _ := procHeapDestroy.Call(heap)
	return ret != 0
}

// GlobalAlloc 分配全局内存块。
// 参数flags：分配标志，例如 GMEM_MOVEABLE。
// 参数size：分配大小，单位字节。
// 返回值：返回全局内存句柄；返回 0 表示分配失败。
func GlobalAlloc(flags uint32, size uintptr) uintptr {
	ret, _, _ := procGlobalAlloc.Call(uintptr(flags), size)
	return ret
}

// GlobalLock 锁定全局内存块并取得指针。
// 参数memory：全局内存句柄。
// 返回值：返回内存指针；返回 0 表示锁定失败。
func GlobalLock(memory uintptr) uintptr {
	ret, _, _ := procGlobalLock.Call(memory)
	return ret
}

// GlobalUnlock 解锁全局内存块。
// 参数memory：全局内存句柄。
// 返回值：true 表示内存仍处于锁定状态，false 表示已解锁或调用失败；需要结合 GetLastError 判断。
func GlobalUnlock(memory uintptr) bool {
	ret, _, _ := procGlobalUnlock.Call(memory)
	return ret != 0
}

// GlobalFree 释放全局内存块。
// 参数memory：全局内存句柄。
// 返回值：返回 0 表示释放成功；非 0 表示释放失败并返回原句柄。
func GlobalFree(memory uintptr) uintptr {
	ret, _, _ := procGlobalFree.Call(memory)
	return ret
}

// GlobalSize 获取全局内存块大小。
// 参数memory：全局内存句柄。
// 返回值：返回内存块大小，单位字节；返回 0 表示获取失败或大小为 0。
func GlobalSize(memory uintptr) uintptr {
	ret, _, _ := procGlobalSize.Call(memory)
	return ret
}
