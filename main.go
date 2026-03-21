// Package main implements a minimal example CLI app.
package main

import (
	"fmt"
	"os"
)

// version info set at build time
var (
	version = "unknown version"
	commit  = "unknown commit"
	date    = "unknown date"
)

func main() {
	fmt.Printf("%s version %s @ %s (%s)\n", os.Args[0], version, commit, date)
}
