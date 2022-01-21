package config

import (
	"log"
)

var (
	Version        = "0.0.1"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func ShowVersion() {
	log.Printf("%s-%s (%s)", Version, CommitHash, BuildTimestamp)
}

func GetVersion() string {
	return Version
}
