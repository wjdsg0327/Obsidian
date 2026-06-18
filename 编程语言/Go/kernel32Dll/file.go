//go:build windows

package kernel32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procCreateFileW        = kernel32.NewProc("CreateFileW")
	procReadFile           = kernel32.NewProc("ReadFile")
	procWriteFile          = kernel32.NewProc("WriteFile")
	procFlushFileBuffers   = kernel32.NewProc("FlushFileBuffers")
	procSetFilePointerEx   = kernel32.NewProc("SetFilePointerEx")
	procGetFileSizeEx      = kernel32.NewProc("GetFileSizeEx")
	procDeleteFileW        = kernel32.NewProc("DeleteFileW")
	procCopyFileW          = kernel32.NewProc("CopyFileW")
	procMoveFileW          = kernel32.NewProc("MoveFileW")
	procMoveFileExW        = kernel32.NewProc("MoveFileExW")
	procGetFileAttributesW = kernel32.NewProc("GetFileAttributesW")
	procSetFileAttributesW = kernel32.NewProc("SetFileAttributesW")
	procCreateDirectoryW   = kernel32.NewProc("CreateDirectoryW")
	procRemoveDirectoryW   = kernel32.NewProc("RemoveDirectoryW")
	procFindFirstFileW     = kernel32.NewProc("FindFirstFileW")
	procFindNextFileW      = kernel32.NewProc("FindNextFileW")
	procFindClose          = kernel32.NewProc("FindClose")
	procGetTempPathW       = kernel32.NewProc("GetTempPathW")
	procGetTempFileNameW   = kernel32.NewProc("GetTempFileNameW")
	procGetFullPathNameW   = kernel32.NewProc("GetFullPathNameW")
	procGetFileType        = kernel32.NewProc("GetFileType")
)

const (
	GENERIC_READ    = 0x80000000
	GENERIC_WRITE   = 0x40000000
	GENERIC_EXECUTE = 0x20000000
	GENERIC_ALL     = 0x10000000

	FILE_SHARE_READ   = 0x00000001
	FILE_SHARE_WRITE  = 0x00000002
	FILE_SHARE_DELETE = 0x00000004

	CREATE_NEW        = 1
	CREATE_ALWAYS     = 2
	OPEN_EXISTING     = 3
	OPEN_ALWAYS       = 4
	TRUNCATE_EXISTING = 5

	FILE_ATTRIBUTE_READONLY  = 0x00000001
	FILE_ATTRIBUTE_HIDDEN    = 0x00000002
	FILE_ATTRIBUTE_SYSTEM    = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE   = 0x00000020
	FILE_ATTRIBUTE_NORMAL    = 0x00000080

	FILE_BEGIN   = 0
	FILE_CURRENT = 1
	FILE_END     = 2

	INVALID_HANDLE_VALUE    = ^uintptr(0)
	INVALID_FILE_ATTRIBUTES = 0xFFFFFFFF

	MOVEFILE_REPLACE_EXISTING   = 0x00000001
	MOVEFILE_COPY_ALLOWED       = 0x00000002
	MOVEFILE_DELAY_UNTIL_REBOOT = 0x00000004
)

// WIN32_FIND_DATA 表示 FindFirstFileW 和 FindNextFileW 返回的文件信息。
type WIN32_FIND_DATA struct {
	DwFileAttributes   uint32
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	DwReserved0        uint32
	DwReserved1        uint32
	CFileName          [260]uint16
	CAlternateFileName [14]uint16
	DwFileType         uint32
	DwCreatorType      uint32
	WFinderFlags       uint16
}

// FileName 返回 WIN32_FIND_DATA 中的主文件名。
// 返回值：返回枚举到的文件名；名称为空时返回空字符串。
func (d WIN32_FIND_DATA) FileName() string {
	return syscall.UTF16ToString(d.CFileName[:])
}

// FileSize 返回 WIN32_FIND_DATA 中的文件大小。
// 返回值：返回文件字节数；目录项通常返回 0。
func (d WIN32_FIND_DATA) FileSize() int64 {
	return int64(uint64(d.NFileSizeHigh)<<32 | uint64(d.NFileSizeLow))
}

// CreateFile 打开或创建文件、设备、管道等内核对象。
// 参数path：文件、设备或命名对象路径。
// 参数access：访问权限，例如 GENERIC_READ。
// 参数shareMode：共享模式，例如 FILE_SHARE_READ。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 参数creationDisposition：创建方式，例如 OPEN_EXISTING。
// 参数flagsAndAttributes：文件属性和标志。
// 参数templateFile：模板文件句柄，没有模板时传 0。
// 返回值：返回文件句柄和是否打开成功；true 表示句柄有效，false 表示打开或创建失败。
func CreateFile(path string, access, shareMode uint32, securityAttributes *SECURITY_ATTRIBUTES, creationDisposition, flagsAndAttributes uint32, templateFile uintptr) (uintptr, bool) {
	ret, _, _ := procCreateFileW.Call(utf16Ptr(path), uintptr(access), uintptr(shareMode), saPtr(securityAttributes), uintptr(creationDisposition), uintptr(flagsAndAttributes), templateFile)
	return ret, ret != INVALID_HANDLE_VALUE
}

// ReadFile 从文件句柄读取数据。
// 参数handle：文件、管道或设备句柄。
// 参数buffer：接收数据的字节切片。
// 返回值：返回实际读取的字节数和是否读取成功；true 表示读取调用成功。
func ReadFile(handle uintptr, buffer []byte) (uint32, bool) {
	var read uint32
	var ptr uintptr
	if len(buffer) > 0 {
		ptr = uintptr(unsafe.Pointer(&buffer[0]))
	}
	ret, _, _ := procReadFile.Call(handle, ptr, uintptr(len(buffer)), uintptr(unsafe.Pointer(&read)), 0)
	return read, ret != 0
}

// WriteFile 向文件句柄写入数据。
// 参数handle：文件、管道或设备句柄。
// 参数buffer：要写入的字节切片。
// 返回值：返回实际写入的字节数和是否写入成功；true 表示写入调用成功。
func WriteFile(handle uintptr, buffer []byte) (uint32, bool) {
	var written uint32
	var ptr uintptr
	if len(buffer) > 0 {
		ptr = uintptr(unsafe.Pointer(&buffer[0]))
	}
	ret, _, _ := procWriteFile.Call(handle, ptr, uintptr(len(buffer)), uintptr(unsafe.Pointer(&written)), 0)
	return written, ret != 0
}

// FlushFileBuffers 刷新文件缓冲区到磁盘或设备。
// 参数handle：文件、管道或设备句柄。
// 返回值：true 表示刷新成功，false 表示刷新失败。
func FlushFileBuffers(handle uintptr) bool {
	ret, _, _ := procFlushFileBuffers.Call(handle)
	return ret != 0
}

// SetFilePointer 设置文件指针位置。
// 参数handle：文件句柄。
// 参数distance：移动距离或目标偏移。
// 参数moveMethod：移动基准，例如 FILE_BEGIN。
// 返回值：返回新的文件指针位置和是否设置成功；true 表示 newPosition 有效。
func SetFilePointer(handle uintptr, distance int64, moveMethod uint32) (int64, bool) {
	var newPosition int64
	ret, _, _ := procSetFilePointerEx.Call(handle, uintptr(distance), uintptr(unsafe.Pointer(&newPosition)), uintptr(moveMethod))
	return newPosition, ret != 0
}

// GetFileSize 获取文件大小。
// 参数handle：文件句柄。
// 返回值：返回文件字节数和是否获取成功；true 表示 size 有效。
func GetFileSize(handle uintptr) (int64, bool) {
	var size int64
	ret, _, _ := procGetFileSizeEx.Call(handle, uintptr(unsafe.Pointer(&size)))
	return size, ret != 0
}

// DeleteFile 删除指定文件。
// 参数path：要删除的文件路径。
// 返回值：true 表示文件删除成功，false 表示删除失败。
func DeleteFile(path string) bool {
	ret, _, _ := procDeleteFileW.Call(utf16Ptr(path))
	return ret != 0
}

// CopyFile 复制文件。
// 参数src：源文件路径。
// 参数dst：目标文件路径。
// 参数failIfExists：true 表示目标已存在时失败，false 表示允许覆盖。
// 返回值：true 表示文件复制成功，false 表示复制失败。
func CopyFile(src, dst string, failIfExists bool) bool {
	ret, _, _ := procCopyFileW.Call(utf16Ptr(src), utf16Ptr(dst), boolArg(failIfExists))
	return ret != 0
}

// MoveFile 移动或重命名文件。
// 参数src：源文件路径。
// 参数dst：目标文件路径。
// 返回值：true 表示移动或重命名成功，false 表示失败。
func MoveFile(src, dst string) bool {
	ret, _, _ := procMoveFileW.Call(utf16Ptr(src), utf16Ptr(dst))
	return ret != 0
}

// MoveFileEx 按标志移动、替换或延迟移动文件。
// 参数src：源文件路径。
// 参数dst：目标文件路径；某些标志下可为空。
// 参数flags：移动标志，例如 MOVEFILE_REPLACE_EXISTING。
// 返回值：true 表示移动请求成功，false 表示失败。
func MoveFileEx(src, dst string, flags uint32) bool {
	ret, _, _ := procMoveFileExW.Call(utf16Ptr(src), utf16PtrOrNil(dst), uintptr(flags))
	return ret != 0
}

// GetFileAttributes 获取文件或目录属性。
// 参数path：文件或目录路径。
// 返回值：返回属性位掩码和是否获取成功；true 表示 attrs 有效。
func GetFileAttributes(path string) (attrs uint32, ok bool) {
	ret, _, _ := procGetFileAttributesW.Call(utf16Ptr(path))
	return uint32(ret), uint32(ret) != INVALID_FILE_ATTRIBUTES
}

// SetFileAttributes 设置文件或目录属性。
// 参数path：文件或目录路径。
// 参数attrs：属性位掩码。
// 返回值：true 表示属性设置成功，false 表示设置失败。
func SetFileAttributes(path string, attrs uint32) bool {
	ret, _, _ := procSetFileAttributesW.Call(utf16Ptr(path), uintptr(attrs))
	return ret != 0
}

// CreateDirectory 创建目录。
// 参数path：要创建的目录路径。
// 参数securityAttributes：安全属性结构指针，nil 表示默认安全属性。
// 返回值：true 表示目录创建成功，false 表示创建失败。
func CreateDirectory(path string, securityAttributes *SECURITY_ATTRIBUTES) bool {
	ret, _, _ := procCreateDirectoryW.Call(utf16Ptr(path), saPtr(securityAttributes))
	return ret != 0
}

// RemoveDirectory 删除空目录。
// 参数path：要删除的目录路径。
// 返回值：true 表示目录删除成功，false 表示删除失败。
func RemoveDirectory(path string) bool {
	ret, _, _ := procRemoveDirectoryW.Call(utf16Ptr(path))
	return ret != 0
}

// FindFirstFile 开始枚举匹配模式的文件。
// 参数pattern：文件匹配模式，例如 C:\\Temp\\*。
// 返回值：返回查找句柄、首个文件信息和是否成功；true 表示句柄和 data 有效。
func FindFirstFile(pattern string) (uintptr, WIN32_FIND_DATA, bool) {
	var data WIN32_FIND_DATA
	handle, _, _ := procFindFirstFileW.Call(utf16Ptr(pattern), uintptr(unsafe.Pointer(&data)))
	return handle, data, handle != INVALID_HANDLE_VALUE
}

// FindNextFile 继续枚举下一个文件。
// 参数findHandle：FindFirstFile 返回的查找句柄。
// 返回值：返回下一个文件信息和是否枚举成功；false 表示没有更多文件或枚举失败。
func FindNextFile(findHandle uintptr) (WIN32_FIND_DATA, bool) {
	var data WIN32_FIND_DATA
	ret, _, _ := procFindNextFileW.Call(findHandle, uintptr(unsafe.Pointer(&data)))
	return data, ret != 0
}

// FindClose 关闭文件查找句柄。
// 参数findHandle：FindFirstFile 返回的查找句柄。
// 返回值：true 表示查找句柄关闭成功，false 表示关闭失败。
func FindClose(findHandle uintptr) bool {
	ret, _, _ := procFindClose.Call(findHandle)
	return ret != 0
}

// GetTempPath 获取当前临时目录路径。
// 返回值：返回临时目录路径；获取失败时返回空字符串。
func GetTempPath() string {
	buf := make([]uint16, 260)
	ret, _, _ := procGetTempPathW.Call(uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetTempFileName 创建临时文件名。
// 参数path：临时目录路径。
// 参数prefix：临时文件名前缀，最多使用前三个字符。
// 参数unique：唯一编号；传 0 时系统生成唯一文件名并创建空文件。
// 返回值：返回生成的临时文件路径和是否成功；true 表示路径有效。
func GetTempFileName(path, prefix string, unique uint32) (string, bool) {
	buf := make([]uint16, 260)
	ret, _, _ := procGetTempFileNameW.Call(utf16Ptr(path), utf16Ptr(prefix), uintptr(unique), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf), ret != 0
}

// GetFullPathName 获取路径的完整绝对路径。
// 参数path：待解析路径。
// 返回值：返回完整路径；解析失败时返回空字符串。
func GetFullPathName(path string) string {
	buf := make([]uint16, 1024)
	ret, _, _ := procGetFullPathNameW.Call(utf16Ptr(path), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])), 0)
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetFileType 获取文件句柄类型。
// 参数handle：文件、管道、控制台或设备句柄。
// 返回值：返回句柄类型代码；0 表示未知类型或调用失败。
func GetFileType(handle uintptr) uint32 {
	ret, _, _ := procGetFileType.Call(handle)
	return uint32(ret)
}
