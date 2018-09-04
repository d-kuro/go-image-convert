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
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeInvalidArgsError
	ExitCodeProcessError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run command.
func (c *CLI) Run(args []string) int {

	var from, to string
	flags := flag.NewFlagSet("convert", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&from, "from", "jpg",
		"input file extension (support: jpg/png/gif, default: jpg)")
	flags.StringVar(&to, "to", "png",
		"output file extension (support: jpg/png/gif, default: png)")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeParseFlagError
	}

	dirName := flags.Arg(0)

	info, err := os.Stat(dirName)
	if err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeInvalidArgsError
	}
	if info.IsDir() == false {
		fmt.Fprintf(c.ErrStream, "%s is not directory\n", dirName)
		return ExitCodeInvalidArgsError
	}

	option := &option.Option{FromExtension: from, ToExtension: to}
	convert := convert.NewConvert(option)

	if err := walkDirectory(dirName, from, convert); err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
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