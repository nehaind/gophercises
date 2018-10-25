package primitive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//Mode means what kind of image the user want
type Mode int

//Mode defines the all modes supported by the primitive command
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

//WithMode is used to transform image with specific mode
// func WithMode(mode Mode) func() []string {
// 	return func() []string {
// 		return []string{"-m", fmt.Sprintf("%d", mode)}
// 	}
// }

//WithMode is used to transform image with specific mode
func WithMode(mode Mode) []string {
	return []string{"-m", fmt.Sprintf("%d", mode)}
}

//Transform is used to transform the image
func Transform(image io.Reader, e string, numShapes int, opts []string) (io.Reader, error) {
	var args []string
	var err error
	var inFile *os.File
	var outFile *os.File
	for _, opt := range opts {
		args = append(args, opt)
	}
	inFile, err = createTempFile("in_", e)
	if err == nil {
		defer os.Remove(inFile.Name())
		outFile, err = createTempFile("in_", e)
		if err == nil {
			defer os.Remove(outFile.Name())
			_, err = io.Copy(inFile, image)
			if err == nil {
				_, err = runPrimitive(inFile.Name(), outFile.Name(), numShapes, args...)
				if err == nil {
					b := bytes.NewBuffer(nil)
					_, err = io.Copy(b, outFile)
					return b, err
				}
			}
		}
	}
	return nil, err
}

func runPrimitive(inFile, outFile string, numShape int, args ...string) (string, error) {
	arg := (fmt.Sprintf("-i %s -o %s -n %d", inFile, outFile, numShape))
	args = append(strings.Fields(arg), args...)
	cmd := exec.Command("primitive", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func createTempFile(name, e string) (*os.File, error) {
	f, err := ioutil.TempFile("", name)
	if err != nil {
		return nil, errors.New("failed to create temporary file")
	}
	defer os.Remove(f.Name())
	return os.Create(fmt.Sprintf("%s.%s", f.Name(), e))
}
