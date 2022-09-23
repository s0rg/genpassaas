package config

import "fmt"

var (
	GitHash   string
	BuildDate string
)

func Version() (rv string) {
	return fmt.Sprintf("%s [build at %s]", GitHash, BuildDate)
}
