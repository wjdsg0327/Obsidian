//go:build windows

package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	kernel32Dll "../.."
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run . <pid> [preview_bytes]")
		fmt.Println("示例: go run . 1234 32")
		return
	}

	pid64, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("pid 格式错误:", err)
		return
	}

	previewBytes := uint64(32)
	if len(os.Args) >= 3 {
		previewBytes, err = strconv.ParseUint(os.Args[2], 10, 32)
		if err != nil {
			fmt.Println("preview_bytes 格式错误:", err)
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

	sysInfo := kernel32Dll.GetNativeSystemInfo()
	address := sysInfo.LpMinimumApplicationAddress
	maxAddress := sysInfo.LpMaximumApplicationAddress

	for address < maxAddress {
		mbi, ok := kernel32Dll.VirtualQueryEx(process, address)
		if !ok || mbi.RegionSize == 0 {
			break
		}

		next := mbi.BaseAddress + mbi.RegionSize
		if mbi.State == kernel32Dll.MEM_COMMIT && isReadable(mbi.Protect) {
			fmt.Printf(
				"0x%016X - 0x%016X  size=%10d  protect=%s  type=%s\n",
				mbi.BaseAddress,
				next,
				mbi.RegionSize,
				protectName(mbi.Protect),
				typeName(mbi.Type),
			)

			if previewBytes > 0 {
				size := previewBytes
				if uintptr(size) > mbi.RegionSize {
					size = uint64(mbi.RegionSize)
				}
				buf := make([]byte, size)
				n, ok := kernel32Dll.ReadProcessMemory(process, mbi.BaseAddress, buf)
				if ok && n > 0 {
					fmt.Print(hex.Dump(buf[:n]))
				} else {
					fmt.Println("  预览读取失败，错误码:", kernel32Dll.GetLastErrorCode())
				}
			}
		}

		if next <= address {
			break
		}
		address = next
	}
}

func isReadable(protect uint32) bool {
	if protect&kernel32Dll.PAGE_GUARD != 0 || protect&kernel32Dll.PAGE_NOACCESS != 0 {
		return false
	}
	base := protect & 0xFF
	switch base {
	case kernel32Dll.PAGE_READONLY,
		kernel32Dll.PAGE_READWRITE,
		kernel32Dll.PAGE_WRITECOPY,
		kernel32Dll.PAGE_EXECUTE_READ,
		kernel32Dll.PAGE_EXECUTE_READWRITE,
		kernel32Dll.PAGE_EXECUTE_WRITECOPY:
		return true
	default:
		return false
	}
}

func protectName(protect uint32) string {
	flags := ""
	if protect&kernel32Dll.PAGE_GUARD != 0 {
		flags += "|GUARD"
	}
	base := protect & 0xFF
	switch base {
	case kernel32Dll.PAGE_READONLY:
		return "READONLY" + flags
	case kernel32Dll.PAGE_READWRITE:
		return "READWRITE" + flags
	case kernel32Dll.PAGE_WRITECOPY:
		return "WRITECOPY" + flags
	case kernel32Dll.PAGE_EXECUTE:
		return "EXECUTE" + flags
	case kernel32Dll.PAGE_EXECUTE_READ:
		return "EXECUTE_READ" + flags
	case kernel32Dll.PAGE_EXECUTE_READWRITE:
		return "EXECUTE_READWRITE" + flags
	case kernel32Dll.PAGE_EXECUTE_WRITECOPY:
		return "EXECUTE_WRITECOPY" + flags
	case kernel32Dll.PAGE_NOACCESS:
		return "NOACCESS" + flags
	default:
		return fmt.Sprintf("0x%X", protect)
	}
}

func typeName(t uint32) string {
	switch t {
	case kernel32Dll.MEM_PRIVATE:
		return "PRIVATE"
	case kernel32Dll.MEM_MAPPED:
		return "MAPPED"
	case kernel32Dll.MEM_IMAGE:
		return "IMAGE"
	default:
		return fmt.Sprintf("0x%X", t)
	}
}
