package config

import "fmt"

var (
	GitTag    string
	GitHash   string
	BuildDate string
)

func Version() (rv string) {
	return fmt.Sprintf("%s-%s [build at %s]", GitTag, GitHash, BuildDate)
}
