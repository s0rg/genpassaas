package config

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	ver := Version()

	if !strings.Contains(ver, GitTag) {
		t.Fatal("no tag")
	}

	if !strings.Contains(ver, GitHash) {
		t.Fatal("no hash")
	}
}
