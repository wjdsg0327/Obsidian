# uiautomation

```go
package main

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/uandersonricardo/uiautomation"
)

// 1. 定义 InvokePattern 虚表结构 📐
type IUIAutomationInvokePatternVtbl struct {
	ole.IUnknownVtbl
	Invoke uintptr // Invoke() 方法索引
}

// 2. 调用 COM Invoke 接口触发点击 🖱️
func invokePattern(patternUnknown *ole.IUnknown) error {
	vTable := (*IUIAutomationInvokePatternVtbl)(unsafe.Pointer(patternUnknown.RawVTable))
	hr, _, _ := syscall.SyscallN(vTable.Invoke, uintptr(unsafe.Pointer(patternUnknown)))
	if hr != 0 {
		return fmt.Errorf("Invoke 调用失败, HRESULT: 0x%X", hr)
	}
	return nil
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitialize(0); err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer ole.CoUninitialize()

	auto, err := uiautomation.NewUIAutomation()
	if err != nil {
		log.Fatalf("UIAutomation 初始化失败: %v", err)
	}

	// 1. 查找主窗口 🪟
	winBstr := ole.SysAllocStringLen("ok-nte v1.2.26 China  - ok-nte")
	defer ole.SysFreeString(winBstr)
	winVal := ole.NewVariant(ole.VT_BSTR, int64(uintptr(unsafe.Pointer(winBstr))))

	winCondition, err := auto.CreatePropertyCondition(30005, &winVal)
	root, _ := auto.GetRootElement()
	window, err := root.FindFirst(uiautomation.TreeScopeChildren, winCondition)

	if err != nil || window == nil {
		log.Fatalf("未找到 BetterGI 窗口 ⚠️")
	}

	// 2. 构造组合条件：ControlType=50000 (Button) 且 Name="启动" 🔘
	typeVal := ole.NewVariant(ole.VT_I4, 50000)
	condType, _ := auto.CreatePropertyCondition(30003, &typeVal)

	btnName := "截图方式"
	btnBstr := ole.SysAllocStringLen(btnName)
	defer ole.SysFreeString(btnBstr)
	nameVal := ole.NewVariant(ole.VT_BSTR, int64(uintptr(unsafe.Pointer(btnBstr))))
	condName, _ := auto.CreatePropertyCondition(30005, &nameVal)

	btnCondition, err := auto.CreateAndCondition(condType, condName)
	if err != nil {
		log.Fatalf("创建组合条件失败: %v", err)
	}

	// 3. 查找符合条件的 Button 控件 🔍
	button, err := window.FindFirst(uiautomation.TreeScopeSubtree, btnCondition)
	if err != nil || button == nil {
		log.Fatalf("未找到类型为 Button 且名称为 [%s] 的控件 ⚠️", btnName)
	}
	fmt.Println("成功定位到真正的 Button 控件！🎉")

	// 4. 获取 Invoke Pattern (ID: 10000) 并执行点击 ⚙️
	patternUnknown, err := button.GetCurrentPattern(10000)
	if err != nil || patternUnknown == nil {
		log.Fatalf("找到按钮，但该控件仍不支持 Invoke 操作 ⚠️")
	}
	defer patternUnknown.Release()

	err = invokePattern(patternUnknown)
	if err != nil {
		log.Fatalf("点击执行失败: %v", err)
	}

	fmt.Println("使用 Invoke 模式点击成功！🚀")
}

```

