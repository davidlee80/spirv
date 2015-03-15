// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

// Application name and version constants.
const (
	AppName         = "spirv-dump"
	AppVersionMajor = 0
	AppVersionMinor = 1
)

// Version returns the application version as a string.
func Version() string {
	return fmt.Sprintf("%s %d.%d (Go runtime %s).\nCopyright (c) 2010-2015, Jim Teeuwen.",
		AppName, AppVersionMajor, AppVersionMinor, runtime.Version())
}
