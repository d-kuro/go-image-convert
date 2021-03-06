package main

import (
	"os"

	"github.com/d-kuro/go-image-convert/cli"
	_ "github.com/d-kuro/go-image-convert/convert/gif"
	_ "github.com/d-kuro/go-image-convert/convert/jpg"
	_ "github.com/d-kuro/go-image-convert/convert/png"
)

func main() {
	cli := cli.NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
