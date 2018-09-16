// Package cli is image convert cli tool.
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"

	"github.com/d-kuro/go-image-convert/convert"
	"github.com/d-kuro/go-image-convert/option"
)

// Exit code.
const (
	// Exclude special meaningful exit code
	ExitCodeOK               = 0
	ExitCodeParseFlagError   = 64
	ExitCodeInvalidArgsError = 65
	ExitCodeProcessError     = 66
)

type CLI struct {
	outStream, errStream io.Writer
}

func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

// Run command.
func (c *CLI) Run(args []string) int {

	var from, to string
	flags := flag.NewFlagSet("convert", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&from, "from", "jpg",
		"input file extension (support: jpg/png/gif, default: jpg)")
	flags.StringVar(&to, "to", "png",
		"output file extension (support: jpg/png/gif, default: png)")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeParseFlagError
	}

	dirName := flags.Arg(0)

	info, err := os.Stat(dirName)
	if err != nil {
		fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeInvalidArgsError
	}
	if info.IsDir() == false {
		fmt.Fprintf(c.errStream, "%s is not directory\n", dirName)
		return ExitCodeInvalidArgsError
	}

	option := &option.Option{FromExtension: from, ToExtension: to}
	convert := convert.NewConvert(option, c.outStream)

	if err := walkDirectory(dirName, from, convert); err != nil {
		fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeProcessError
	}

	return ExitCodeOK
}

func walkDirectory(dirName, fromExtension string, convert convert.Convert) error {
	var eg errgroup.Group

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+fromExtension {
			eg.Go(func() error {
				return convert.Convert(path)
			})
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
