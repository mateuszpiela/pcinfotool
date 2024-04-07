//go:build windows
// +build windows

package osinfo

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func GetOSInformation() *OsReport {
	name, _ := os.Hostname()
	majorVersion, minorVersion, buildNumber := windows.RtlGetNtVersionNumbers()

	version := fmt.Sprint(majorVersion) + "." + fmt.Sprint(minorVersion) + " Build: " + fmt.Sprint(buildNumber)
	
	osReport := OsReport{System: "Windows", Distribution: "Microsoft", Version: version, Hostname: name}

	return &osReport
}