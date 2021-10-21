package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/s0rg/genpassaas/pkg/gen"
)

const appName = "GenPassaaS-cli"

var (
	GitHash   string
	BuildDate string
	version   = flag.Bool("version", false, "show version and exit")
	generator = flag.String("gen", "smart", "generator 'smart' or 'simple'")
	length    = flag.Int("len", 16, "length of each password")
	count     = flag.Int("count", 5, "count of passwords")
)

func myVersion() (v string) {
	return fmt.Sprintf("%s %s [build at %s]", appName, GitHash, BuildDate)
}

func generate(genkind string, length, count int) (rv []string) {
	genfn := gen.Simple

	if genkind == "smart" {
		genfn = gen.Smart
	}

	rv = make([]string, count)

	for i := 0; i < count; i++ {
		rv[i] = genfn(length)
	}

	return rv
}

func main() {
	flag.Parse()

	if *version {
		fmt.Println(myVersion())

		return
	}

	fmt.Println(strings.Join(generate(*generator, *length, *count), "\n"))
}
