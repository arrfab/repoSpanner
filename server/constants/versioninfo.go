package constants

import (
	"fmt"
	"runtime"
)

// These strings will be filled in by build.sh
var version string
var gitdescrip string
var vendordescrip string

func VersionBuiltIn() bool {
	return version != "" && gitdescrip != ""
}

func PublicVersionString() string {
	return fmt.Sprintf("%s+%s",
		version,
		gitdescrip,
	)
}

func VersionString() string {
	suffix := ""
	if hasProfiling {
		suffix = " WITH PROFILING"
	}

	return fmt.Sprintf("%s (deps %s) running on %s/%s (compiled with %s)%s",
		PublicVersionString(),
		vendordescrip,
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
		suffix,
	)
}
