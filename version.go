package libvirtwrap

import (
	"fmt"
)

// Version information
const (
	Major   = 0
	Minor   = 2
	Patch   = 0
	Release = "alpha"
)

func FullVersion() string {
	return fmt.Sprintf("%d.%d.%d-%s", Major, Minor, Patch, Release)
}
