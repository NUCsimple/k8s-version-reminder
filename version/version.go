package version

import "fmt"

var (
	Version   = "1.0.0"
	GitCommit = ""
)

func VersionInfo() string {
	return fmt.Sprintf("k8s-version-reminder version is %s", Version)
}
