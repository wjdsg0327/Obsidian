//go:build windows

package kernel32Dll

import "unsafe"

var (
	procCreateMutexW         = kernel32.NewProc("CreateMutexW")
	procOpenMutexW           = kernel32.NewProc("OpenMutexW")
	procReleaseMutex         = kernel32.NewProc("ReleaseMutex")
	procCreateEventW         = kernel32.NewProc("CreateEventW")
	procOpenEventW           = kernel32.NewProc("OpenEventW")
	procSetEvent             = kernel32.NewProc("SetEvent")
	procResetEvent           = kernel32.NewProc("ResetEvent")
	procPulseEvent           = kernel32.NewProc("PulseEvent")
	procCreateSemaphoreW     = kernel32.NewProc("CreateSemaphoreW")
	procOpenSemaphoreW       = kernel32.NewProc("OpenSemaphoreW")
	procReleaseSemaphore     = kernel32.NewProc("ReleaseSemaphore")
	procCreateWaitableTimerW = kernel32.NewProc("CreateWaitableTimerW")
	procSetWaitableTimer     = kernel32.NewProc("SetWaitableTimer")
	procCancelWaitableTimer  = kernel32.NewProc("CancelWaitableTimer")
)

const (
	SYNCHRONIZE          = 0x00100000
	MUTEX_ALL_ACCESS     = 0x001F0001
	EVENT_ALL_ACCESS     = 0x001F0003
	SEMAPHORE_ALL_ACCESS = 0x001F0003
	TIMER_ALL_ACCESS     = 0x001F0003
)

// CreateMutex 创建或打开命名互斥体。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 参数initialOwner：true 表示调用线程立即拥有互斥体。
// 参数name：互斥体名称；为空表示匿名互斥体。
// 返回值：返回互斥体句柄和是否创建或打开成功；true 表示句柄有效，使用后应 CloseHandle。
func CreateMutex(securityAttributes *SECURITY_ATTRIBUTES, initialOwner bool, name string) (uintptr, bool) {
	ret, _, _ := procCreateMutexW.Call(saPtr(securityAttributes), boolArg(initialOwner), utf16PtrOrNil(name))
	return ret, ret != 0
}

// OpenMutex 打开已有命名互斥体。
// 参数access：访问权限，例如 MUTEX_ALL_ACCESS。
// 参数inheritHandle：返回句柄是否可被子进程继承。
// 参数name：互斥体名称。
// 返回值：返回互斥体句柄和是否打开成功；true 表示句柄有效。
func OpenMutex(access uint32, inheritHandle bool, name string) (uintptr, bool) {
	ret, _, _ := procOpenMutexW.Call(uintptr(access), boolArg(inheritHandle), utf16Ptr(name))
	return ret, ret != 0
}

// ReleaseMutex 释放互斥体所有权。
// 参数mutex：互斥体句柄。
// 返回值：true 表示释放成功，false 表示释放失败。
func ReleaseMutex(mutex uintptr) bool {
	ret, _, _ := procReleaseMutex.Call(mutex)
	return ret != 0
}

// CreateEvent 创建或打开事件对象。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 参数manualReset：true 表示手动重置事件，false 表示自动重置事件。
// 参数initialState：true 表示初始为有信号状态。
// 参数name：事件名称；为空表示匿名事件。
// 返回值：返回事件句柄和是否创建或打开成功；true 表示句柄有效。
func CreateEvent(securityAttributes *SECURITY_ATTRIBUTES, manualReset bool, initialState bool, name string) (uintptr, bool) {
	ret, _, _ := procCreateEventW.Call(saPtr(securityAttributes), boolArg(manualReset), boolArg(initialState), utf16PtrOrNil(name))
	return ret, ret != 0
}

// OpenEvent 打开已有命名事件对象。
// 参数access：访问权限，例如 EVENT_ALL_ACCESS。
// 参数inheritHandle：返回句柄是否可被子进程继承。
// 参数name：事件名称。
// 返回值：返回事件句柄和是否打开成功；true 表示句柄有效。
func OpenEvent(access uint32, inheritHandle bool, name string) (uintptr, bool) {
	ret, _, _ := procOpenEventW.Call(uintptr(access), boolArg(inheritHandle), utf16Ptr(name))
	return ret, ret != 0
}

// SetEvent 将事件设置为有信号状态。
// 参数event：事件句柄。
// 返回值：true 表示设置成功，false 表示设置失败。
func SetEvent(event uintptr) bool {
	ret, _, _ := procSetEvent.Call(event)
	return ret != 0
}

// ResetEvent 将事件设置为无信号状态。
// 参数event：事件句柄。
// 返回值：true 表示重置成功，false 表示重置失败。
func ResetEvent(event uintptr) bool {
	ret, _, _ := procResetEvent.Call(event)
	return ret != 0
}

// PulseEvent 瞬间触发事件对象。
// 参数event：事件句柄。
// 返回值：true 表示触发调用成功，false 表示调用失败；该 API 语义不可靠，新代码不推荐使用。
func PulseEvent(event uintptr) bool {
	ret, _, _ := procPulseEvent.Call(event)
	return ret != 0
}

// CreateSemaphore 创建或打开信号量。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 参数initialCount：初始计数。
// 参数maximumCount：最大计数。
// 参数name：信号量名称；为空表示匿名信号量。
// 返回值：返回信号量句柄和是否创建或打开成功；true 表示句柄有效。
func CreateSemaphore(securityAttributes *SECURITY_ATTRIBUTES, initialCount int32, maximumCount int32, name string) (uintptr, bool) {
	ret, _, _ := procCreateSemaphoreW.Call(saPtr(securityAttributes), uintptr(initialCount), uintptr(maximumCount), utf16PtrOrNil(name))
	return ret, ret != 0
}

// OpenSemaphore 打开已有命名信号量。
// 参数access：访问权限，例如 SEMAPHORE_ALL_ACCESS。
// 参数inheritHandle：返回句柄是否可被子进程继承。
// 参数name：信号量名称。
// 返回值：返回信号量句柄和是否打开成功；true 表示句柄有效。
func OpenSemaphore(access uint32, inheritHandle bool, name string) (uintptr, bool) {
	ret, _, _ := procOpenSemaphoreW.Call(uintptr(access), boolArg(inheritHandle), utf16Ptr(name))
	return ret, ret != 0
}

// ReleaseSemaphore 增加信号量计数。
// 参数semaphore：信号量句柄。
// 参数releaseCount：释放数量。
// 返回值：返回释放前的计数和是否释放成功；true 表示 previousCount 有效。
func ReleaseSemaphore(semaphore uintptr, releaseCount int32) (previousCount int32, ok bool) {
	ret, _, _ := procReleaseSemaphore.Call(semaphore, uintptr(releaseCount), uintptr(unsafe.Pointer(&previousCount)))
	return previousCount, ret != 0
}

// CreateWaitableTimer 创建或打开等待计时器。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 参数manualReset：true 表示手动重置计时器。
// 参数name：计时器名称；为空表示匿名计时器。
// 返回值：返回计时器句柄和是否创建或打开成功；true 表示句柄有效。
func CreateWaitableTimer(securityAttributes *SECURITY_ATTRIBUTES, manualReset bool, name string) (uintptr, bool) {
	ret, _, _ := procCreateWaitableTimerW.Call(saPtr(securityAttributes), boolArg(manualReset), utf16PtrOrNil(name))
	return ret, ret != 0
}

// SetWaitableTimer 设置等待计时器。
// 参数timer：计时器句柄。
// 参数dueTime：触发时间指针，负值表示相对时间，单位 100 纳秒。
// 参数period：周期毫秒数，0 表示只触发一次。
// 参数resume：true 表示尝试唤醒休眠系统。
// 返回值：true 表示计时器设置成功，false 表示设置失败。
func SetWaitableTimer(timer uintptr, dueTime *int64, period int32, resume bool) bool {
	ret, _, _ := procSetWaitableTimer.Call(timer, uintptr(unsafe.Pointer(dueTime)), uintptr(period), 0, 0, boolArg(resume))
	return ret != 0
}

// CancelWaitableTimer 取消等待计时器。
// 参数timer：计时器句柄。
// 返回值：true 表示计时器取消成功，false 表示取消失败。
func CancelWaitableTimer(timer uintptr) bool {
	ret, _, _ := procCancelWaitableTimer.Call(timer)
	return ret != 0
}
