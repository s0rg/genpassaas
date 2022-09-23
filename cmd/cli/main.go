package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/s0rg/genpassaas/pkg/config"
	"github.com/s0rg/genpassaas/pkg/gen"
)

const appName = config.Name + "-cli"

var (
	version   = flag.Bool("version", false, "show version and exit")
	generator = flag.String("gen", "smart", "generator 'smart' or 'simple'")
	length    = flag.Int("len", config.DefaultLength, "length of each password")
	count     = flag.Int("count", config.DefaultCount, "count of passwords")
)

func generate(kind string, length, count int) (rv []string) {
	var fn gen.Fn

	switch kind {
	case "smart":
		fn = gen.Smart
	default:
		fn = gen.Simple
	}

	rv = make([]string, count)

	for i := 0; i < count; i++ {
		rv[i] = fn(length)
	}

	return rv
}

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", appName, config.Version())

		return
	}

	fmt.Println(strings.Join(
		generate(
			*generator,
			config.ClampLength(*length),
			config.ClampCount(*count),
		), "\n"))
}
