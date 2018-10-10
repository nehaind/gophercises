package transform

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Mode defines the shapes used when transforming images.
type Mode int

// Modes supported by the primitive package.
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

//Transform function takes an input file, takes the no of shape and create the image from it.
func Transform(image io.Reader, ext string, numShapes int, ModeCom Mode) (io.Reader, error) {
	b := bytes.NewBuffer(nil)
	var err error
	in, _ := os.Create(fmt.Sprintf("%s.%s", "in", ext))

	defer os.Remove(in.Name())
	out, _ := os.Create(fmt.Sprintf("%s.%s", "out", ext))

	//defer os.Remove(out.Name())
	// Read image into in file
	_, err = io.Copy(in, image)
	if err == nil {
		var OutputString string

		OutputString, err = primitive(in.Name(), out.Name(), numShapes, ModeCom)
		fmt.Println(OutputString)
		// _, err = io.Copy(b, out)
		//my code
		// b := []byte(OutputString)

		// ioutil.WriteFile(out.Name(), b, 0644)
		//changes made

		b := bytes.NewBuffer(nil)
		_, err = io.Copy(b, out)
		if err != nil {
			return nil, errors.New("primitive: Failed to copy output file into byte buffer")
		}
		return b, nil

	}
	return b, nil
}

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argStr)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
