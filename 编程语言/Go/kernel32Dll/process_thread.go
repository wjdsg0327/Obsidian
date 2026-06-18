//go:build windows

package kernel32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procGetCurrentProcess   = kernel32.NewProc("GetCurrentProcess")
	procGetCurrentProcessId = kernel32.NewProc("GetCurrentProcessId")
	procOpenProcess         = kernel32.NewProc("OpenProcess")
	procTerminateProcess    = kernel32.NewProc("TerminateProcess")
	procGetExitCodeProcess  = kernel32.NewProc("GetExitCodeProcess")
	procCreateProcessW      = kernel32.NewProc("CreateProcessW")
	procGetCurrentThread    = kernel32.NewProc("GetCurrentThread")
	procGetCurrentThreadId  = kernel32.NewProc("GetCurrentThreadId")
	procOpenThread          = kernel32.NewProc("OpenThread")
	procCreateThread        = kernel32.NewProc("CreateThread")
	procExitThread          = kernel32.NewProc("ExitThread")
	procGetExitCodeThread   = kernel32.NewProc("GetExitCodeThread")
	procSuspendThread       = kernel32.NewProc("SuspendThread")
	procResumeThread        = kernel32.NewProc("ResumeThread")
	procWaitForSingleObject = kernel32.NewProc("WaitForSingleObject")
	procWaitForMultipleObj  = kernel32.NewProc("WaitForMultipleObjects")
	procSleep               = kernel32.NewProc("Sleep")
)

const (
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_ALL_ACCESS                = 0x001F0FFF

	THREAD_TERMINATE         = 0x0001
	THREAD_SUSPEND_RESUME    = 0x0002
	THREAD_GET_CONTEXT       = 0x0008
	THREAD_SET_CONTEXT       = 0x0010
	THREAD_QUERY_INFORMATION = 0x0040
	THREAD_ALL_ACCESS        = 0x001F03FF

	STARTF_USESHOWWINDOW = 0x00000001
	STILL_ACTIVE         = 259
)

// STARTUPINFO 表示 CreateProcessW 使用的启动参数。
type STARTUPINFO struct {
	Cb              uint32
	LpReserved      uintptr
	LpDesktop       uintptr
	LpTitle         uintptr
	DwX             uint32
	DwY             uint32
	DwXSize         uint32
	DwYSize         uint32
	DwXCountChars   uint32
	DwYCountChars   uint32
	DwFillAttribute uint32
	DwFlags         uint32
	WShowWindow     uint16
	CbReserved2     uint16
	LpReserved2     uintptr
	HStdInput       uintptr
	HStdOutput      uintptr
	HStdError       uintptr
}

// PROCESS_INFORMATION 表示 CreateProcessW 返回的进程和线程信息。
type PROCESS_INFORMATION struct {
	HProcess    uintptr
	HThread     uintptr
	DwProcessId uint32
	DwThreadId  uint32
}

// GetCurrentProcess 获取当前进程伪句柄。
// 返回值：返回当前进程伪句柄；该句柄不需要 CloseHandle。
func GetCurrentProcess() uintptr {
	ret, _, _ := procGetCurrentProcess.Call()
	return ret
}

// GetCurrentProcessId 获取当前进程 ID。
// 返回值：返回当前进程 ID。
func GetCurrentProcessId() uint32 {
	ret, _, _ := procGetCurrentProcessId.Call()
	return uint32(ret)
}

// OpenProcess 打开指定进程。
// 参数processID：目标进程 ID。
// 参数access：进程访问权限。
// 参数inheritHandle：返回句柄是否可被子进程继承。
// 返回值：返回进程句柄和是否打开成功；true 表示句柄有效，使用后应 CloseHandle。
func OpenProcess(processID uint32, access uint32, inheritHandle bool) (uintptr, bool) {
	ret, _, _ := procOpenProcess.Call(uintptr(access), boolArg(inheritHandle), uintptr(processID))
	return ret, ret != 0
}

// TerminateProcess 终止指定进程。
// 参数process：进程句柄。
// 参数exitCode：进程退出码。
// 返回值：true 表示终止请求成功，false 表示请求失败。
func TerminateProcess(process uintptr, exitCode uint32) bool {
	ret, _, _ := procTerminateProcess.Call(process, uintptr(exitCode))
	return ret != 0
}

// GetExitCodeProcess 获取进程退出码。
// 参数process：进程句柄。
// 返回值：返回退出码和是否获取成功；退出码为 STILL_ACTIVE 表示进程仍在运行。
func GetExitCodeProcess(process uintptr) (uint32, bool) {
	var code uint32
	ret, _, _ := procGetExitCodeProcess.Call(process, uintptr(unsafe.Pointer(&code)))
	return code, ret != 0
}

// CreateProcess 创建新进程。
// 参数applicationName：可执行文件路径；为空时从 commandLine 解析。
// 参数commandLine：命令行字符串。
// 参数processAttributes：进程安全属性，nil 表示默认。
// 参数threadAttributes：主线程安全属性，nil 表示默认。
// 参数inheritHandles：是否继承句柄。
// 参数creationFlags：进程创建标志。
// 参数environment：环境块指针，0 表示继承父进程环境。
// 参数currentDirectory：工作目录；为空表示继承当前目录。
// 参数startupInfo：启动信息结构指针，nil 时使用默认结构。
// 返回值：返回进程信息和是否创建成功；true 表示进程与主线程句柄有效，使用后应 CloseHandle。
func CreateProcess(applicationName, commandLine string, processAttributes, threadAttributes *SECURITY_ATTRIBUTES, inheritHandles bool, creationFlags uint32, environment uintptr, currentDirectory string, startupInfo *STARTUPINFO) (PROCESS_INFORMATION, bool) {
	var pi PROCESS_INFORMATION
	var si STARTUPINFO
	if startupInfo == nil {
		si.Cb = uint32(unsafe.Sizeof(STARTUPINFO{}))
		startupInfo = &si
	} else if startupInfo.Cb == 0 {
		startupInfo.Cb = uint32(unsafe.Sizeof(STARTUPINFO{}))
	}
	var cmdPtr uintptr
	if commandLine != "" {
		cmd, err := syscall.UTF16FromString(commandLine)
		if err != nil {
			return pi, false
		}
		cmdPtr = uintptr(unsafe.Pointer(&cmd[0]))
	}
	ret, _, _ := procCreateProcessW.Call(utf16PtrOrNil(applicationName), cmdPtr, saPtr(processAttributes), saPtr(threadAttributes), boolArg(inheritHandles), uintptr(creationFlags), environment, utf16PtrOrNil(currentDirectory), uintptr(unsafe.Pointer(startupInfo)), uintptr(unsafe.Pointer(&pi)))
	return pi, ret != 0
}

// GetCurrentThread 获取当前线程伪句柄。
// 返回值：返回当前线程伪句柄；该句柄不需要 CloseHandle。
func GetCurrentThread() uintptr {
	ret, _, _ := procGetCurrentThread.Call()
	return ret
}

// GetCurrentThreadId 获取当前线程 ID。
// 返回值：返回当前线程 ID。
func GetCurrentThreadId() uint32 {
	ret, _, _ := procGetCurrentThreadId.Call()
	return uint32(ret)
}

// OpenThread 打开指定线程。
// 参数threadID：目标线程 ID。
// 参数access：线程访问权限。
// 参数inheritHandle：返回句柄是否可被子进程继承。
// 返回值：返回线程句柄和是否打开成功；true 表示句柄有效，使用后应 CloseHandle。
func OpenThread(threadID uint32, access uint32, inheritHandle bool) (uintptr, bool) {
	ret, _, _ := procOpenThread.Call(uintptr(access), boolArg(inheritHandle), uintptr(threadID))
	return ret, ret != 0
}

// CreateThread 创建新线程。
// 参数securityAttributes：线程安全属性，nil 表示默认。
// 参数stackSize：线程初始栈大小，0 表示默认。
// 参数startAddress：线程入口函数地址。
// 参数parameter：传给线程入口函数的参数。
// 参数creationFlags：线程创建标志。
// 返回值：返回线程句柄、线程 ID 和是否创建成功；true 表示句柄有效，使用后应 CloseHandle。
func CreateThread(securityAttributes *SECURITY_ATTRIBUTES, stackSize uintptr, startAddress uintptr, parameter uintptr, creationFlags uint32) (uintptr, uint32, bool) {
	var threadID uint32
	ret, _, _ := procCreateThread.Call(saPtr(securityAttributes), stackSize, startAddress, parameter, uintptr(creationFlags), uintptr(unsafe.Pointer(&threadID)))
	return ret, threadID, ret != 0
}

// ExitThread 退出当前线程。
// 参数exitCode：线程退出码。
// 返回值：无；调用后当前线程结束。
func ExitThread(exitCode uint32) {
	procExitThread.Call(uintptr(exitCode))
}

// GetExitCodeThread 获取线程退出码。
// 参数thread：线程句柄。
// 返回值：返回退出码和是否获取成功；退出码为 STILL_ACTIVE 表示线程仍在运行。
func GetExitCodeThread(thread uintptr) (uint32, bool) {
	var code uint32
	ret, _, _ := procGetExitCodeThread.Call(thread, uintptr(unsafe.Pointer(&code)))
	return code, ret != 0
}

// SuspendThread 挂起线程。
// 参数thread：线程句柄。
// 返回值：返回挂起前的挂起计数；返回 0xFFFFFFFF 表示调用失败。
func SuspendThread(thread uintptr) uint32 {
	ret, _, _ := procSuspendThread.Call(thread)
	return uint32(ret)
}

// ResumeThread 恢复线程。
// 参数thread：线程句柄。
// 返回值：返回恢复前的挂起计数；返回 0xFFFFFFFF 表示调用失败。
func ResumeThread(thread uintptr) uint32 {
	ret, _, _ := procResumeThread.Call(thread)
	return uint32(ret)
}

// WaitForSingleObject 等待单个内核对象变为有信号状态。
// 参数handle：等待对象句柄。
// 参数milliseconds：超时时间，单位毫秒；INFINITE 表示无限等待。
// 返回值：返回等待结果代码，例如 WAIT_OBJECT_0、WAIT_TIMEOUT 或 WAIT_FAILED。
func WaitForSingleObject(handle uintptr, milliseconds uint32) uint32 {
	ret, _, _ := procWaitForSingleObject.Call(handle, uintptr(milliseconds))
	return uint32(ret)
}

// WaitForMultipleObjects 等待多个内核对象。
// 参数handles：等待对象句柄数组。
// 参数waitAll：true 表示等待全部对象，false 表示任一对象满足即可。
// 参数milliseconds：超时时间，单位毫秒；INFINITE 表示无限等待。
// 返回值：返回等待结果代码；具体含义与 WaitForMultipleObjects 的 WAIT_* 返回值一致。
func WaitForMultipleObjects(handles []uintptr, waitAll bool, milliseconds uint32) uint32 {
	var first uintptr
	if len(handles) > 0 {
		first = uintptr(unsafe.Pointer(&handles[0]))
	}
	ret, _, _ := procWaitForMultipleObj.Call(uintptr(len(handles)), first, boolArg(waitAll), uintptr(milliseconds))
	return uint32(ret)
}

// Sleep 暂停当前线程一段时间。
// 参数milliseconds：暂停时间，单位毫秒。
// 返回值：无；函数会阻塞当前线程直到时间到达。
func Sleep(milliseconds uint32) {
	procSleep.Call(uintptr(milliseconds))
}
