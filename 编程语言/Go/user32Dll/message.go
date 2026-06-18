//go:build windows

package user32Dll

import "unsafe"

var (
	procSendMessageW            = user32.NewProc("SendMessageW")
	procSendMessageTimeoutW     = user32.NewProc("SendMessageTimeoutW")
	procPostMessageW            = user32.NewProc("PostMessageW")
	procPostThreadMessageW      = user32.NewProc("PostThreadMessageW")
	procRegisterWindowMessageW  = user32.NewProc("RegisterWindowMessageW")
	procGetMessageW             = user32.NewProc("GetMessageW")
	procTranslateMessage        = user32.NewProc("TranslateMessage")
	procDispatchMessageW        = user32.NewProc("DispatchMessageW")
	procPeekMessageW            = user32.NewProc("PeekMessageW")
	procPostQuitMessage         = user32.NewProc("PostQuitMessage")
	procSetTimer                = user32.NewProc("SetTimer")
	procKillTimer               = user32.NewProc("KillTimer")
	procCallMsgFilterW          = user32.NewProc("CallMsgFilterW")
	procBroadcastSystemMessageW = user32.NewProc("BroadcastSystemMessageW")
)

const (
	WM_CLOSE  = 0x0010
	WM_USER   = 0x0400
	PM_REMOVE = 0x0001

	SMTO_ABORTIFHUNG = 0x0002

	WM_TIMER = 0x0113

	BSF_POSTMESSAGE  = 0x00000010
	BSM_APPLICATIONS = 0x00000008
)

// MSG 表示 Win32 消息循环中的消息结构。
type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

// CloseBySendMessage 同步发送 WM_CLOSE 关闭窗口消息。
// 参数hwnd：窗口句柄。
// 返回值：返回目标窗口处理 WM_CLOSE 后的消息结果值。
func CloseBySendMessage(hwnd uintptr) uintptr {
	ret, _, _ := procSendMessageW.Call(hwnd, WM_CLOSE, 0, 0)
	return ret
}

// SendMessage 向窗口同步发送指定消息。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：返回目标窗口过程处理该消息后的结果值，具体含义由消息类型决定。
func SendMessage(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessageW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret
}

// SendCloseWithTimeout 带超时发送 WM_CLOSE，避免目标窗口无响应导致阻塞。
// 参数hwnd：窗口句柄。
// 参数timeoutMs：超时时间，单位毫秒。
// 返回值：true 表示 WM_CLOSE 在超时前发送并处理完成，false 表示发送失败或目标超时无响应。
func SendCloseWithTimeout(hwnd uintptr, timeoutMs uint32) bool {
	var result uintptr
	ret, _, _ := procSendMessageTimeoutW.Call(hwnd, WM_CLOSE, 0, 0, SMTO_ABORTIFHUNG, uintptr(timeoutMs), uintptr(unsafe.Pointer(&result)))
	return ret != 0
}

// SendMessageTimeout 带超时同步发送指定窗口消息。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 参数flags：调用标志位。
// 参数timeoutMs：超时时间，单位毫秒。
// 返回值：返回消息处理结果和是否在超时前成功完成；true 表示 result 有效，false 表示失败或超时。
func SendMessageTimeout(hwnd uintptr, msg uint32, wParam, lParam uintptr, flags, timeoutMs uint32) (uintptr, bool) {
	var result uintptr
	ret, _, _ := procSendMessageTimeoutW.Call(hwnd, uintptr(msg), wParam, lParam, uintptr(flags), uintptr(timeoutMs), uintptr(unsafe.Pointer(&result)))
	return result, ret != 0
}

// CloseByPostMessage 异步投递 WM_CLOSE 关闭窗口消息。
// 参数hwnd：窗口句柄。
// 返回值：true 表示 WM_CLOSE 已成功投递到消息队列，false 表示投递失败。
func CloseByPostMessage(hwnd uintptr) bool {
	ret, _, _ := procPostMessageW.Call(hwnd, WM_CLOSE, 0, 0)
	return ret != 0
}

// PostMessage 异步投递指定窗口消息。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：true 表示消息已成功投递到窗口消息队列，false 表示投递失败。
func PostMessage(hwnd uintptr, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessageW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret != 0
}

// NotifyThread 向线程消息队列投递一个示例 WM_USER 消息。
// 参数threadID：目标线程 ID。
// 返回值：true 表示示例 WM_USER 消息已成功投递到线程消息队列，false 表示投递失败。
func NotifyThread(threadID uint32) bool {
	ret, _, _ := procPostThreadMessageW.Call(uintptr(threadID), WM_USER+1, 123, 456)
	return ret != 0
}

// PostThreadMessage 向线程消息队列投递指定消息。
// 参数threadID：目标线程 ID。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：true 表示消息已成功投递到线程消息队列，false 表示投递失败。
func PostThreadMessage(threadID uint32, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostThreadMessageW.Call(uintptr(threadID), uintptr(msg), wParam, lParam)
	return ret != 0
}

// RegisterMyMessage 注册唯一窗口消息并返回消息 ID。
// 参数name：类型为 string 的调用参数。
// 返回值：返回注册得到的全局唯一消息 ID；返回 0 表示注册失败。
func RegisterMyMessage(name string) uint32 {
	msg, _, _ := procRegisterWindowMessageW.Call(utf16Ptr(name))
	return uint32(msg)
}

// MessageLoop 运行标准阻塞消息循环。
// 返回值：返回收到 WM_QUIT 时的退出码；返回 -1 表示 GetMessageW 出错。
func MessageLoop() int {
	var msg MSG
	for {
		ret, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if int32(ret) == -1 {
			return -1
		}
		if ret == 0 {
			return int(msg.WParam)
		}
		procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	}
}

// PumpOnce 非阻塞读取并分发一条消息。
// 返回值：true 表示读取并分发了一条消息，false 表示当前没有可处理消息。
func PumpOnce() bool {
	var msg MSG
	ret, _, _ := procPeekMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0, PM_REMOVE)
	if ret == 0 {
		return false
	}
	procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
	procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	return true
}

// QuitApp 向当前线程投递退出消息。
// 参数exitCode：退出码。
// 返回值：无返回值；函数会向当前线程消息队列投递 WM_QUIT。
func QuitApp(exitCode int) {
	procPostQuitMessage.Call(uintptr(exitCode))
}

// StartTimer 为窗口启动一个 1000 毫秒计时器。
// 参数hwnd：窗口句柄。
// 返回值：返回新建计时器 ID；返回 0 表示启动失败。
func StartTimer(hwnd uintptr) uintptr {
	id, _, _ := procSetTimer.Call(hwnd, 1, 1000, 0)
	return id
}

// SetWindowTimer 为窗口启动指定 ID 和间隔的计时器。
// 参数hwnd：窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数intervalMs：计时器间隔，单位毫秒。
// 返回值：返回计时器 ID；返回 0 表示启动失败。
func SetWindowTimer(hwnd uintptr, id, intervalMs uintptr) uintptr {
	ret, _, _ := procSetTimer.Call(hwnd, id, intervalMs, 0)
	return ret
}

// StopTimer 停止窗口计时器。
// 参数hwnd：窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：true 表示计时器停止成功，false 表示停止失败。
func StopTimer(hwnd uintptr, id uintptr) bool {
	ret, _, _ := procKillTimer.Call(hwnd, id)
	return ret != 0
}

// CallMessageFilter 调用消息过滤 Hook。
// 参数msg：窗口消息编号。
// 参数code：消息过滤代码。
// 返回值：true 表示消息被过滤器处理，false 表示未处理或调用失败。
func CallMessageFilter(msg *MSG, code int) bool {
	ret, _, _ := procCallMsgFilterW.Call(uintptr(unsafe.Pointer(msg)), uintptr(code))
	return ret != 0
}

// BroadcastToApplications 向应用程序广播指定系统消息。
// 参数message：要发送或广播的消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：返回广播结果；正值通常表示成功，负值表示失败。
func BroadcastToApplications(message uint32, wParam, lParam uintptr) int32 {
	recipients := uint32(BSM_APPLICATIONS)
	ret, _, _ := procBroadcastSystemMessageW.Call(BSF_POSTMESSAGE, uintptr(unsafe.Pointer(&recipients)), uintptr(message), wParam, lParam)
	return int32(ret)
}
