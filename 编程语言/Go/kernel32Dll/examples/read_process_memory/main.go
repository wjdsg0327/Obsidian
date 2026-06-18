//go:build windows

package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	kernel32Dll "../.."
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: go run . <pid> <address_hex> [size]")
		fmt.Println("示例: go run . 1234 0x7FF612340000 64")
		return
	}

	pid64, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("pid 格式错误:", err)
		return
	}

	addressText := strings.TrimPrefix(strings.ToLower(os.Args[2]), "0x")
	address, err := strconv.ParseUint(addressText, 16, 64)
	if err != nil {
		fmt.Println("地址格式错误:", err)
		return
	}

	size := uint64(32)
	if len(os.Args) >= 4 {
		size, err = strconv.ParseUint(os.Args[3], 10, 32)
		if err != nil {
			fmt.Println("size 格式错误:", err)
			return
		}
	}

	process, ok := kernel32Dll.OpenProcess(
		uint32(pid64),
		kernel32Dll.PROCESS_VM_READ|kernel32Dll.PROCESS_QUERY_INFORMATION,
		false,
	)
	if !ok {
		fmt.Println("OpenProcess 失败，错误码:", kernel32Dll.GetLastErrorCode())
		return
	}
	defer kernel32Dll.CloseHandle(process)

	buffer := make([]byte, size)
	n, ok := kernel32Dll.ReadProcessMemory(process, uintptr(address), buffer)
	if !ok {
		fmt.Println("ReadProcessMemory 失败，错误码:", kernel32Dll.GetLastErrorCode())
		return
	}

	fmt.Printf("成功读取 %d 字节\n", n)
	fmt.Println(hex.Dump(buffer[:n]))
}
