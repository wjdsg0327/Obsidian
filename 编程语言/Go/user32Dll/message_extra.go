//go:build windows

package user32Dll

import "unsafe"

var (
	procSendNotifyMessageW        = user32.NewProc("SendNotifyMessageW")
	procSendMessageCallbackW      = user32.NewProc("SendMessageCallbackW")
	procReplyMessage              = user32.NewProc("ReplyMessage")
	procInSendMessageEx           = user32.NewProc("InSendMessageEx")
	procWaitMessage               = user32.NewProc("WaitMessage")
	procMsgWaitForMultipleObjects = user32.NewProc("MsgWaitForMultipleObjects")
)

const (
	QS_ALLINPUT = 0x04FF
)

// SendNotifyMessage 异步或同步发送通知消息。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：true 表示消息发送或投递成功，false 表示失败。
func SendNotifyMessage(hwnd uintptr, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procSendNotifyMessageW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret != 0
}

// SendMessageCallback 发送消息并在处理完成后调用回调函数。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 参数callback：消息处理完成后调用的回调函数指针。
// 参数data：传递给回调函数的自定义数据。
// 返回值：true 表示消息发送请求成功，false 表示发送请求失败。
func SendMessageCallback(hwnd uintptr, msg uint32, wParam, lParam, callback, data uintptr) bool {
	ret, _, _ := procSendMessageCallbackW.Call(hwnd, uintptr(msg), wParam, lParam, callback, data)
	return ret != 0
}

// ReplyMessage 回复 SendMessage 调用方。
// 参数result：返回给发送方的消息处理结果。
// 返回值：true 表示回复成功，false 表示当前线程没有可回复的同步消息或回复失败。
func ReplyMessage(result uintptr) bool {
	ret, _, _ := procReplyMessage.Call(result)
	return ret != 0
}

// InSendMessageEx 查询当前窗口过程是否正在处理其他线程发送的同步消息。
// 参数reserved：保留参数，通常传 0。
// 返回值：返回消息发送状态标志位；0 表示当前不在 SendMessage 处理链中。
func InSendMessageEx(reserved uintptr) uint32 {
	ret, _, _ := procInSendMessageEx.Call(reserved)
	return uint32(ret)
}

// WaitMessage 等待当前线程消息队列出现新消息。
// 返回值：true 表示等待成功并有新消息可处理，false 表示等待失败。
func WaitMessage() bool {
	ret, _, _ := procWaitMessage.Call()
	return ret != 0
}

// MsgWaitForMultipleObjects 等待内核对象或消息队列事件。
// 参数handles：内核对象句柄数组。
// 参数waitAll：true 表示等待全部对象，false 表示任一对象满足即可。
// 参数milliseconds：超时时间，单位毫秒。
// 参数wakeMask：消息唤醒掩码，例如 QS_ALLINPUT。
// 返回值：返回等待结果代码；具体含义与 MsgWaitForMultipleObjects 的 WAIT_* 返回值一致。
func MsgWaitForMultipleObjects(handles []uintptr, waitAll bool, milliseconds uint32, wakeMask uint32) uint32 {
	var first uintptr
	if len(handles) > 0 {
		first = uintptr(unsafe.Pointer(&handles[0]))
	}
	ret, _, _ := procMsgWaitForMultipleObjects.Call(uintptr(len(handles)), first, boolArg(waitAll), uintptr(milliseconds), uintptr(wakeMask))
	return uint32(ret)
}
