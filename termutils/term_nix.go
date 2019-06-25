// +build linux darwin freebsd netbsd openbsd dragonfly
// +build !appengine

package termutils

import "syscall"

const sysIoctl = syscall.SYS_IOCTL
