package cli_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	c "github.com/d-kuro/go-image-convert/cli"
	_ "github.com/d-kuro/go-image-convert/convert/gif"
	_ "github.com/d-kuro/go-image-convert/convert/jpg"
	_ "github.com/d-kuro/go-image-convert/convert/png"
)

func TestCLI_Run(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := c.NewCLI(outStream, errStream)
	args := strings.Split("convert ./../testdata", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeOK {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() > 0 {
		t.Errorf("failed cli run, output: %q", errStream.String())
	}
}

func TestCLI_Run_ParseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := c.NewCLI(outStream, errStream)
	args := strings.Split("convert -foo", " ") // undefined option
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeParseFlagError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}

func TestCLI_Run_InvalidArgsError_NotExistDirectory(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := c.NewCLI(outStream, errStream)
	args := strings.Split("convert ./foo", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeInvalidArgsError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}

func TestCLI_Run_InvalidArgsError_NotDirectory(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := c.NewCLI(outStream, errStream)
	args := strings.Split("convert ./../testdata/gopher.jpg", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeInvalidArgsError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}

func createTempFile(t *testing.T) string {
	t.Helper()
	tempFile, err := os.Create("unsupported_extension.foo")
	if err != nil {
		t.Fatal("failed create temp file", err)
	}
	defer tempFile.Close()
	return tempFile.Name()
}

func TestCLI_Run_ProcessError_Convert(t *testing.T) {
	tempFileName := createTempFile(t)
	defer os.Remove(tempFileName)

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := c.NewCLI(outStream, errStream)
	args := strings.Split("convert -from foo ./", " ")
	exitCode := cli.Run(args)

	if exitCode != c.ExitCodeProcessError {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() == 0 {
		t.Errorf("failed error message is not output")
	}
}
