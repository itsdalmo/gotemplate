package main

import (
	"fmt"
	"os"

	"github.com/itsdalmo/gotemplate/internal/cli"
)

func main() {
	app := cli.New(cli.Options{Writer: os.Stdout})

	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
