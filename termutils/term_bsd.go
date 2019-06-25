// +build darwin freebsd netbsd openbsd dragonfly
// +build !appengine

package termutils

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
const ioctlWriteTermios = syscall.TIOCSETA
