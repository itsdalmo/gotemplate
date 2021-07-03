package main

import (
	"context"
	"fmt"
	"os"

	"github.com/itsdalmo/gotemplate/internal/cli"
)

func main() {
	app := cli.New(cli.Options{Writer: os.Stdout})

	if err := app.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}
}
