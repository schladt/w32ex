/* Contains selected Win32 APIs not found in other repos*/

package w32ex

import (
	"errors"
	"syscall"
	"unsafe"
)

type FILETIME syscall.Filetime
type SYSTEMTIME syscall.Systemtime

// BOOL WINAPI SystemTimeToFileTime
//   _In_  const SYSTEMTIME *lpSystemTime,
//   _Out_       LPFILETIME lpFileTime
func SystemTimeToFileTime(inputSystemTime *SYSTEMTIME) (FILETIME, error) {

	var outputFileTime FILETIME
	var err error

	libkernel32, _ := syscall.LoadLibrary("kernel32.dll")
	sytemTimeToFileTime, _ := syscall.GetProcAddress(libkernel32, "SystemTimeToFileTime")
	ret, _, _ := syscall.Syscall(sytemTimeToFileTime, 2,
		uintptr(unsafe.Pointer(inputSystemTime)),
		uintptr(unsafe.Pointer(&outputFileTime)),
		0)

	if ret == 0 {
		err = errors.New("Unable to call SystemTimeToFileTime")
	}
	return outputFileTime, err
}

// BOOL WINAPI FileTimeToSystemTime
//   _In_  const FILETIME     *lpFileTime,
//   _Out_       LPSYSTEMTIME lpSystemTime
func FileTimeToSystemTime(inputFileTime *FILETIME) (SYSTEMTIME, error) {

	var outputSystemTime SYSTEMTIME
	var err error

	libkernel32, _ := syscall.LoadLibrary("kernel32.dll")
	fileTimeToSystemTime, _ := syscall.GetProcAddress(libkernel32, "FileTimeToSystemTime")
	ret, _, _ := syscall.Syscall(fileTimeToSystemTime, 2,
		uintptr(unsafe.Pointer(inputFileTime)),
		uintptr(unsafe.Pointer(&outputSystemTime)),
		0)

	if ret == 0 {
		err = errors.New("Unable to call FileTimeToSystemTime")
	}
	return outputSystemTime, err
}

// LONG WINAPI CompareFileTime
//   _In_ const FILETIME *lpFileTime1,
//   _In_ const FILETIME *lpFileTime2
func CompareFileTime(t1 *FILETIME, t2 *FILETIME) int {
	libkernel32, _ := syscall.LoadLibrary("kernel32.dll")
	fileTimeToSystemTime, _ := syscall.GetProcAddress(libkernel32, "CompareFileTime")
	ret, _, _ := syscall.Syscall(fileTimeToSystemTime, 2,
		uintptr(unsafe.Pointer(t1)),
		uintptr(unsafe.Pointer(t2)),
		0)

	return int(int16(ret))
}

// void WINAPI GetSystemTime
//   _Out_ LPSYSTEMTIME lpSystemTime
func GetSystemTime() SYSTEMTIME {
	var systemTime SYSTEMTIME
	libkernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getSystemTime, _ := syscall.GetProcAddress(libkernel32, "GetSystemTime")
	_, _, _ = syscall.Syscall(getSystemTime, 1,
		uintptr(unsafe.Pointer(&systemTime)),
		0,
		0)

	return systemTime
}
