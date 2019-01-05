// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall2

import "syscall"
import "unsafe"

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := syscall.RawSyscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func readlen(fd int, p *byte, np int) (n int, err error) {
	r0, _, e1 := syscall.RawSyscall(syscall.SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np))
	n = int(r0)
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}

func rawVforkSyscall(trap, a1 uintptr) (r1 uintptr, err syscall.Errno)
