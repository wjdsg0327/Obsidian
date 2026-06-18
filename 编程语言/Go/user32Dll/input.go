//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procGetAsyncKeyState           = user32.NewProc("GetAsyncKeyState")
	procGetKeyState                = user32.NewProc("GetKeyState")
	procGetKeyboardState           = user32.NewProc("GetKeyboardState")
	procMapVirtualKeyW             = user32.NewProc("MapVirtualKeyW")
	procToUnicode                  = user32.NewProc("ToUnicode")
	procGetCursorPos               = user32.NewProc("GetCursorPos")
	procSetCursorPos               = user32.NewProc("SetCursorPos")
	procSendInput                  = user32.NewProc("SendInput")
	procKeybdEvent                 = user32.NewProc("keybd_event")
	procMouseEvent                 = user32.NewProc("mouse_event")
	procRegisterHotKey             = user32.NewProc("RegisterHotKey")
	procUnregisterHotKey           = user32.NewProc("UnregisterHotKey")
	procRegisterRawInputDevices    = user32.NewProc("RegisterRawInputDevices")
	procGetRawInputData            = user32.NewProc("GetRawInputData")
	procRegisterTouchWindow        = user32.NewProc("RegisterTouchWindow")
	procGetTouchInputInfo          = user32.NewProc("GetTouchInputInfo")
	procCloseTouchInputHandle      = user32.NewProc("CloseTouchInputHandle")
	procRegisterPointerInputTarget = user32.NewProc("RegisterPointerInputTarget")
	procGetPointerInfo             = user32.NewProc("GetPointerInfo")
	procGetKeyboardLayout          = user32.NewProc("GetKeyboardLayout")
	procActivateKeyboardLayout     = user32.NewProc("ActivateKeyboardLayout")
	procLoadKeyboardLayoutW        = user32.NewProc("LoadKeyboardLayoutW")
	procUnloadKeyboardLayout       = user32.NewProc("UnloadKeyboardLayout")
	procVkKeyScanW                 = user32.NewProc("VkKeyScanW")
	procDragDetect                 = user32.NewProc("DragDetect")
	procBlockInput                 = user32.NewProc("BlockInput")
)

const (
	VK_CONTROL = 0x11
	VK_CAPITAL = 0x14
	VK_A       = 0x41
	VK_RETURN  = 0x0D

	MAPVK_VK_TO_VSC = 0

	INPUT_MOUSE         = 0
	INPUT_KEYBOARD      = 1
	KEYEVENTF_KEYUP     = 0x0002
	KEYEVENTF_KEYUP_OLD = 0x0002

	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004

	MOD_ALT     = 0x0001
	MOD_CONTROL = 0x0002
	WM_HOTKEY   = 0x0312

	WM_INPUT        = 0x00FF
	RIDEV_INPUTSINK = 0x00000100
	RID_INPUT       = 0x10000003

	WM_TOUCH = 0x0240
	PT_TOUCH = 0x00000002

	KLF_ACTIVATE = 0x00000001
)

// KEYBDINPUT 表示 SendInput 使用的键盘输入结构。
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// INPUT 表示 SendInput 使用的输入结构。
type INPUT struct {
	Type    uint32
	Ki      KEYBDINPUT
	Padding [16]byte
}

// RAWINPUTDEVICE 表示 RegisterRawInputDevices 使用的原始输入设备。
type RAWINPUTDEVICE struct {
	UsUsagePage uint16
	UsUsage     uint16
	DwFlags     uint32
	HwndTarget  uintptr
}

// IsKeyDown 检测指定虚拟键是否正在按下。
// 参数vk：虚拟键码。
// 返回值：true 表示指定虚拟键当前处于按下状态，false 表示未按下。
func IsKeyDown(vk int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

// IsCapsLockOn 判断 CapsLock 是否开启。
// 返回值：true 表示 CapsLock 当前开启，false 表示关闭。
func IsCapsLockOn() bool {
	ret, _, _ := procGetKeyState.Call(VK_CAPITAL)
	return ret&1 != 0
}

// GetKeyboardState 获取 256 个虚拟键状态。
// 返回值：返回 256 个虚拟键状态和是否获取成功；true 表示状态数组有效，false 表示获取失败。
func GetKeyboardState() ([256]byte, bool) {
	var state [256]byte
	ret, _, _ := procGetKeyboardState.Call(uintptr(unsafe.Pointer(&state[0])))
	return state, ret != 0
}

// VirtualKeyToScanCode 将虚拟键码转换为扫描码。
// 参数vk：虚拟键码。
// 返回值：返回虚拟键对应的扫描码；返回 0 通常表示无法转换。
func VirtualKeyToScanCode(vk uint32) uint32 {
	ret, _, _ := procMapVirtualKeyW.Call(uintptr(vk), MAPVK_VK_TO_VSC)
	return uint32(ret)
}

// KeyToUnicode 将键盘状态和键码转换为 Unicode 字符。
// 参数vk：虚拟键码。
// 参数scanCode：键盘扫描码。
// 参数state：键盘状态数组指针。
// 返回值：返回按键转换得到的 Unicode 字符串；无法转换、死键或失败时返回空字符串。
func KeyToUnicode(vk, scanCode uint32, state *[256]byte) string {
	buf := make([]uint16, 8)
	ret, _, _ := procToUnicode.Call(uintptr(vk), uintptr(scanCode), uintptr(unsafe.Pointer(&state[0])), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0)
	if int32(ret) <= 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetMousePos 获取鼠标屏幕坐标。
// 返回值：返回鼠标屏幕坐标和是否获取成功；true 表示坐标有效，false 表示获取失败。
func GetMousePos() (POINT, bool) {
	var pt POINT
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

// SetMousePos 设置鼠标屏幕坐标。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：true 表示鼠标位置设置成功，false 表示设置失败。
func SetMousePos(x, y int32) bool {
	ret, _, _ := procSetCursorPos.Call(uintptr(x), uintptr(y))
	return ret != 0
}

// PressA 使用 SendInput 模拟按下并松开 A 键。
// 返回值：返回 SendInput 成功插入的输入事件数量；通常返回 2 表示按下和松开都成功。
func PressA() uint32 {
	return SendKeyboardTap(VK_A)
}

// SendKeyboardTap 使用 SendInput 模拟指定虚拟键按下并松开。
// 参数vk：虚拟键码。
// 返回值：返回 SendInput 成功插入的输入事件数量；通常返回 2 表示按下和松开都成功。
func SendKeyboardTap(vk uint16) uint32 {
	inputs := []INPUT{
		{Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: vk}},
		{Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: vk, DwFlags: KEYEVENTF_KEYUP}},
	}
	ret, _, _ := procSendInput.Call(uintptr(len(inputs)), uintptr(unsafe.Pointer(&inputs[0])), unsafe.Sizeof(inputs[0]))
	return uint32(ret)
}

// PressEnterOld 使用旧式 keybd_event 模拟回车键。
// 返回值：无返回值；函数会用旧式 keybd_event 发送回车按下和松开事件。
func PressEnterOld() {
	procKeybdEvent.Call(VK_RETURN, 0, 0, 0)
	procKeybdEvent.Call(VK_RETURN, 0, KEYEVENTF_KEYUP_OLD, 0)
}

// LeftClickOld 使用旧式 mouse_event 模拟鼠标左键点击。
// 返回值：无返回值；函数会用旧式 mouse_event 发送鼠标左键按下和松开事件。
func LeftClickOld() {
	procMouseEvent.Call(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0)
	procMouseEvent.Call(MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)
}

// RegisterCtrlAltK 注册 Ctrl+Alt+K 全局快捷键。
// 返回值：true 表示 Ctrl+Alt+K 全局快捷键注册成功，false 表示注册失败或快捷键被占用。
func RegisterCtrlAltK() bool {
	id := uintptr(1)
	ret, _, _ := procRegisterHotKey.Call(0, id, MOD_CONTROL|MOD_ALT, 'K')
	return ret != 0
}

// RegisterHotKey 注册指定全局快捷键。
// 参数hwnd：窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数modifiers：快捷键修饰键标志。
// 参数vk：虚拟键码。
// 返回值：true 表示全局快捷键注册成功，false 表示注册失败或快捷键被占用。
func RegisterHotKey(hwnd uintptr, id int, modifiers, vk uint32) bool {
	ret, _, _ := procRegisterHotKey.Call(hwnd, uintptr(id), uintptr(modifiers), uintptr(vk))
	return ret != 0
}

// UnregisterCtrlAltK 注销 Ctrl+Alt+K 全局快捷键。
// 返回值：无返回值；函数会尝试注销 Ctrl+Alt+K 全局快捷键。
func UnregisterCtrlAltK() {
	procUnregisterHotKey.Call(0, 1)
}

// UnregisterHotKey 注销指定全局快捷键。
// 参数hwnd：窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：true 表示指定全局快捷键注销成功，false 表示注销失败。
func UnregisterHotKey(hwnd uintptr, id int) bool {
	ret, _, _ := procUnregisterHotKey.Call(hwnd, uintptr(id))
	return ret != 0
}

// RegisterRawKeyboard 注册窗口接收原始键盘输入。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口已成功注册接收原始键盘输入，false 表示注册失败。
func RegisterRawKeyboard(hwnd uintptr) bool {
	rid := RAWINPUTDEVICE{UsUsagePage: 0x01, UsUsage: 0x06, DwFlags: RIDEV_INPUTSINK, HwndTarget: hwnd}
	ret, _, _ := procRegisterRawInputDevices.Call(uintptr(unsafe.Pointer(&rid)), 1, unsafe.Sizeof(rid))
	return ret != 0
}

// RegisterRawInputDevices 注册一个或多个原始输入设备。
// 参数devices：原始输入设备数组首地址。
// 参数count：设备数量。
// 返回值：true 表示原始输入设备注册成功，false 表示注册失败。
func RegisterRawInputDevices(devices *RAWINPUTDEVICE, count uint32) bool {
	ret, _, _ := procRegisterRawInputDevices.Call(uintptr(unsafe.Pointer(devices)), uintptr(count), unsafe.Sizeof(RAWINPUTDEVICE{}))
	return ret != 0
}

// GetRawInputDataSize 获取原始输入数据大小。
// 参数rawInput：原始输入句柄。
// 返回值：返回原始输入数据大小和是否查询成功；true 表示大小有效，false 表示查询失败。
func GetRawInputDataSize(rawInput uintptr) (uint32, bool) {
	var size uint32
	ret, _, _ := procGetRawInputData.Call(rawInput, RID_INPUT, 0, uintptr(unsafe.Pointer(&size)), unsafe.Sizeof(uintptr(0))*2)
	return size, ret != ^uintptr(0)
}

// GetRawInputData 调用 GetRawInputData 读取原始输入数据。
// 参数rawInput：原始输入句柄。
// 参数command：GetRawInputData 命令。
// 参数data：接收数据的缓冲区指针。
// 参数size：输入输出参数，表示缓冲区大小或实际大小。
// 参数headerSize：RAWINPUTHEADER 结构大小。
// 返回值：返回读取到的字节数；返回 ^uintptr(0) 表示读取失败。
func GetRawInputData(rawInput uintptr, command uint32, data uintptr, size *uint32, headerSize uintptr) uintptr {
	ret, _, _ := procGetRawInputData.Call(rawInput, uintptr(command), data, uintptr(unsafe.Pointer(size)), headerSize)
	return ret
}

// RegisterTouch 注册窗口接收触摸输入。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口已成功注册接收触摸输入，false 表示注册失败。
func RegisterTouch(hwnd uintptr) bool {
	ret, _, _ := procRegisterTouchWindow.Call(hwnd, 0)
	return ret != 0
}

// GetTouchInputInfo 读取触摸输入信息到调用方提供的缓冲区。
// 参数touchInput：触摸输入句柄。
// 参数inputCount：触摸点数量。
// 参数inputs：接收 TOUCHINPUT 数组的缓冲区指针。
// 参数inputSize：单个 TOUCHINPUT 结构大小。
// 返回值：true 表示触摸输入信息读取成功，false 表示读取失败。
func GetTouchInputInfo(touchInput uintptr, inputCount uint32, inputs uintptr, inputSize int32) bool {
	ret, _, _ := procGetTouchInputInfo.Call(touchInput, uintptr(inputCount), inputs, uintptr(inputSize))
	return ret != 0
}

// CloseTouchInput 关闭触摸输入句柄。
// 参数handle：需要关闭的输入句柄。
// 返回值：true 表示触摸输入句柄关闭成功，false 表示关闭失败。
func CloseTouchInput(handle uintptr) bool {
	ret, _, _ := procCloseTouchInputHandle.Call(handle)
	return ret != 0
}

// RegisterPointerTouch 注册窗口接收 Pointer 触摸输入。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口已成功注册接收 Pointer 触摸输入，false 表示注册失败。
func RegisterPointerTouch(hwnd uintptr) bool {
	ret, _, _ := procRegisterPointerInputTarget.Call(hwnd, PT_TOUCH)
	return ret != 0
}

// GetPointerInfo 读取 Pointer 输入信息到调用方提供的缓冲区。
// 参数pointerID：Pointer 输入 ID。
// 参数pointerInfo：接收 Pointer 信息的缓冲区指针。
// 返回值：true 表示 Pointer 信息读取成功，false 表示读取失败。
func GetPointerInfo(pointerID uint32, pointerInfo uintptr) bool {
	ret, _, _ := procGetPointerInfo.Call(uintptr(pointerID), pointerInfo)
	return ret != 0
}

// CurrentKeyboardLayout 获取当前线程键盘布局。
// 返回值：返回当前线程键盘布局句柄；返回 0 表示获取失败。
func CurrentKeyboardLayout() uintptr {
	hkl, _, _ := procGetKeyboardLayout.Call(0)
	return hkl
}

// ActivateLayout 激活指定键盘布局并返回旧布局。
// 参数hkl：键盘布局句柄。
// 返回值：返回之前激活的键盘布局句柄；返回 0 表示切换失败或无旧布局。
func ActivateLayout(hkl uintptr) uintptr {
	old, _, _ := procActivateKeyboardLayout.Call(hkl, 0)
	return old
}

// LoadUSKeyboard 加载并激活 English US 键盘布局。
// 返回值：返回加载后的 English US 键盘布局句柄；返回 0 表示加载失败。
func LoadUSKeyboard() uintptr {
	hkl, _, _ := procLoadKeyboardLayoutW.Call(utf16Ptr("00000409"), KLF_ACTIVATE)
	return hkl
}

// UnloadLayout 卸载指定键盘布局。
// 参数hkl：键盘布局句柄。
// 返回值：true 表示键盘布局卸载成功，false 表示卸载失败。
func UnloadLayout(hkl uintptr) bool {
	ret, _, _ := procUnloadKeyboardLayout.Call(hkl)
	return ret != 0
}

// CharToVK 将字符转换为虚拟键码。
// 参数ch：要转换的字符。
// 返回值：返回字符对应的虚拟键和 Shift 状态组合；返回 -1 表示无法转换。
func CharToVK(ch rune) int16 {
	ret, _, _ := procVkKeyScanW.Call(uintptr(ch))
	return int16(ret)
}

// IsDrag 检测指定位置是否形成拖动意图。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：true 表示从指定位置开始的鼠标操作达到拖动阈值，false 表示未形成拖动。
func IsDrag(hwnd uintptr, x, y int32) bool {
	ret, _, _ := procDragDetect.Call(hwnd, packPoint(x, y))
	return ret != 0
}

// BlockUserInput 阻止或恢复键鼠输入。
// 参数block：是否阻止输入。
// 返回值：true 表示阻止或恢复输入的请求成功，false 表示请求失败。
func BlockUserInput(block bool) bool {
	ret, _, _ := procBlockInput.Call(boolArg(block))
	return ret != 0
}
