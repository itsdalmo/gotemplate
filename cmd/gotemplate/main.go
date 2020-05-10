package main

import (
	"os"

	"github.com/itsdalmo/gotemplate/internal/cli"

	"gopkg.in/alecthomas/kingpin.v2"
)

var version string

func main() {
	app := cli.New(os.Stdout).Version(version)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
