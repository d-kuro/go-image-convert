// Package convert is convert image extension.
package convert

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"github.com/d-kuro/go-image-convert/di"
	"github.com/d-kuro/go-image-convert/option"
)

const greenColor = "\x1b[32m%s\x1b[0m"

// Convert is interface that has Convert function.
type Convert interface {
	Convert(path string) error
}

type convert struct {
	option    *option.Option
	outStream io.Writer
}

// NewConvert is Convert interface constructor.
func NewConvert(option *option.Option, outStream io.Writer) Convert {
	return &convert{option: option, outStream: outStream}
}

// Convert image.
func (c *convert) Convert(path string) error {
	fromConverter, err := getConverter(c.option.FromExtension)
	if err != nil {
		return err
	}
	image, err := decode(path, fromConverter)
	if err != nil {
		return err
	}

	toConverter, err := getConverter(c.option.ToExtension)
	if err != nil {
		return err
	}
	if err := encode(image, toConverter, path, c.option.FromExtension, c.option.ToExtension); err != nil {
		return err
	}
	fmt.Fprintf(c.outStream, "%s -> to %s"+greenColor, path, c.option.ToExtension, "\tDONE\n")
	return nil
}

func decode(path string, converter di.Converter) (image.Image, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return converter.Decode(file)
}

func encode(image image.Image, converter di.Converter, path, fromExtension, toExtension string) error {
	output, err := os.Create(strings.TrimSuffix(path, fromExtension) + toExtension)
	if err != nil {
		return err
	}
	defer output.Close()
	return converter.Encode(output, image)
}

func getConverter(extension string) (di.Converter, error) {
	convert, ok := di.Converts[extension]
	if ok {
		return convert, nil
	} else {
		return nil, errors.New("unsupported extension")
	}
}
