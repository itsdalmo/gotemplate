package cli

import (
	"github.com/spf13/cobra"
	"io"

	"github.com/itsdalmo/gotemplate"
)

type Options struct {
	Writer io.Writer
}

type Flags struct {
	Times int
}

func New(opts Options) *cobra.Command {
	f := Flags{}
	c := &cobra.Command{
		Use:   "gotemplate [flags] <arg...>",
		Short: "Template for go CLIs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for i := 0; i < f.Times; i++ {
				for _, m := range args {
					gotemplate.Print(opts.Writer, m)
				}
			}
			return nil
		},
	}

	c.SetOut(opts.Writer)
	c.Flags().IntVar(&f.Times, "times", 1, "Number of times to print the messages")

	return c
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
