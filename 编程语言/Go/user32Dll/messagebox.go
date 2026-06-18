//go:build windows

package user32Dll

var (
	procMessageBoxW = user32.NewProc("MessageBoxW")
	procMessageBeep = user32.NewProc("MessageBeep")
)

const (
	MB_OK              = 0x00000000
	MB_ICONINFORMATION = 0x00000040
	MB_ICONASTERISK    = 0x00000040
)

// ShowInfoBox 显示 Windows 原生信息提示框。
// 参数title：窗口标题或提示框标题。
// 参数text：要传入的文本内容。
// 返回值：返回用户点击的按钮 ID，例如 IDOK；返回 0 表示调用失败。
func ShowInfoBox(title, text string) int {
	ret, _, _ := procMessageBoxW.Call(0, utf16Ptr(text), utf16Ptr(title), MB_OK|MB_ICONINFORMATION)
	return int(ret)
}

// PlayNotifySound 播放系统通知提示音。
// 返回值：true 表示系统提示音播放请求成功，false 表示播放失败。
func PlayNotifySound() bool {
	ret, _, _ := procMessageBeep.Call(MB_ICONASTERISK)
	return ret != 0
}
