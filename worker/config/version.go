package config

import "fmt"

var (
	Version        = "0.0.1"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func ShowVersion() string {
	return fmt.Sprintf("%s-%s (%s)", Version, CommitHash, BuildTimestamp)
}

func GetVersion() string {
	return Version
}
