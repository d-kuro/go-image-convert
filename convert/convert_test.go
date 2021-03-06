package convert_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/d-kuro/go-image-convert/convert"
	_ "github.com/d-kuro/go-image-convert/convert/gif"
	_ "github.com/d-kuro/go-image-convert/convert/jpg"
	_ "github.com/d-kuro/go-image-convert/convert/png"
	"github.com/d-kuro/go-image-convert/option"
)

var testFilePath = map[string]string{}

func init() {
	testFilePath["jpg"] = "./../testdata/gopher.jpg"
	testFilePath["png"] = "./../testdata/gopher.png"
	testFilePath["gif"] = "./../testdata/gopher.gif"
}

var convertTests = []struct {
	name string
	*option.Option
	outStream io.Writer
}{
	{
		"convert from jpg to png",
		&option.Option{FromExtension: "jpg", ToExtension: "png"},
		new(bytes.Buffer),
	},
	{
		"convert from jpg to gif",
		&option.Option{FromExtension: "jpg", ToExtension: "gif"},
		new(bytes.Buffer),
	},
	{
		"convert from png to jpg",
		&option.Option{FromExtension: "png", ToExtension: "jpg"},
		new(bytes.Buffer),
	},
	{
		"convert from png to gif",
		&option.Option{FromExtension: "png", ToExtension: "gif"},
		new(bytes.Buffer),
	},
	{
		"convert from gif to jpg",
		&option.Option{FromExtension: "gif", ToExtension: "jpg"},
		new(bytes.Buffer),
	},
	{
		"convert from gif to png",
		&option.Option{FromExtension: "gif", ToExtension: "png"},
		new(bytes.Buffer),
	},
}

func TestConvert_Convert(t *testing.T) {
	for _, tt := range convertTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			sut := convert.NewConvert(tt.Option, tt.outStream)
			path := testFilePath[tt.FromExtension]
			if err := sut.Convert(path); err != nil {
				t.Error("failed convert", err)
			}
		})
	}
}

func createTempFile(t *testing.T) string {
	t.Helper()
	tempFile, err := ioutil.TempFile("./", "temp_file")
	if err != nil {
		t.Fatal("failed create temp file", err)
	}
	defer tempFile.Close()
	return tempFile.Name()
}

func TestConvert_Convert_UnsupportedExtension(t *testing.T) {
	option := &option.Option{FromExtension: "foo", ToExtension: "bar"}
	outStream := new(bytes.Buffer)
	sut := convert.NewConvert(option, outStream)
	tempFileName := createTempFile(t)
	defer os.Remove(tempFileName)
	err := sut.Convert(tempFileName)
	if err == nil {
		t.Error("failed unsupported extension error is nothing")
	}
	errorMessage := "unsupported extension"
	if err.Error() != errorMessage {
		t.Errorf("failed different error message: %s", err.Error())
	}
}
