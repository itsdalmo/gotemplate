package cli

import (
	"context"
	"flag"
	"io"

	"github.com/itsdalmo/gotemplate"

	"github.com/peterbourgon/ff/v3/ffcli"
)

type Options struct {
	Writer io.Writer
}

type Flags struct {
	Times int
}

func New(opts Options) *ffcli.Command {
	var (
		f  = Flags{}
		fs = flag.NewFlagSet("gotemplate", flag.ContinueOnError)
	)
	fs.SetOutput(opts.Writer)

	fs.IntVar(&f.Times, "times", 1, "Number of times to print the messages")

	return &ffcli.Command{
		Name:       "gotemplate",
		ShortHelp:  "Template for go CLIs",
		ShortUsage: "gotemplate [-n times] <arg>",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			for i := 0; i < f.Times; i++ {
				for _, m := range fs.Args() {
					gotemplate.Print(opts.Writer, m)
				}
			}
			return nil
		},
	}
}

//// New returns a new kingpin.Application.
//func New(w io.Writer) *kingpin.Application {
//	app := kingpin.New("gotemplate", "Template for go CLIs").DefaultEnvars()
//	var (
//		messages = app.Arg("message", "Message(s) to print").Strings()
//		times    = app.Flag("times", "Number of times to print the messages").Default("1").Int()
//	)
//	app.Action(func(_ *kingpin.ParseContext) error {
//		for i := 0; i < *times; i++ {
//			for _, m := range *messages {
//				gotemplate.Print(w, m)
//			}
//		}
//		return nil
//	})
//	return app
//}
