package version

import "fmt"

var (
	version string
	os      string
	arch    string
)

func Version() string {
	return fmt.Sprintf("Version: %s\nPlatform: %s/%s", version, os, arch)
}
