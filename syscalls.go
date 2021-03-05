package gowinkey

import "syscall"

var user32 = syscall.MustLoadDLL("user32")
var mapVirtualKey = user32.MustFindProc("MapVirtualKeyA")

// virtualKeyToString returns the string representation of vk.
//
// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeya
// for more details.
func virtualKeyToString(vk VirtualKey) string {
	char, _, _ := syscall.Syscall(mapVirtualKey.Addr(), 2, uintptr(vk), 2, 0)
	return string(char)
}
