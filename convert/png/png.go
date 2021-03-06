// Package png is encode and decode to image.
package png

import (
	"image"
	"io"

	"image/png"

	"github.com/d-kuro/go-image-convert/di"
)

// Png implements convert.Converter
type Png struct{}

func init() {
	di.Register("png", Png{})
}

// Decode returns image and error
func (Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode return error
func (Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
