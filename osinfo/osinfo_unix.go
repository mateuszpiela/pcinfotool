//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

package osinfo

import (
	"golang.org/x/sys/unix"
)

func GetOSInformation() *OsReport {
	utsname := &unix.Utsname{}
    unix.Uname(utsname)

	osReport := OsReport{
		System: string(utsname.Sysname[:]),
		Distribution: string(utsname.Sysname[:]), 
		Version: string(utsname.Version[:]),
		Hostname: string(utsname.Nodename[:]),
	}

	return &osReport
}