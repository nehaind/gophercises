package primitive

import (
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestWithMode(t *testing.T) {
	result := WithMode(ModeCombo)
	if result == nil {
		t.Error("Expected string but got not result")
	}
}

func TestTransform(t *testing.T) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/ghoper.jpg")
	f, _ := os.Open(imgPath)
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "jpg", 1, opts)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}

func TestTransformNegativePrimitive(t *testing.T) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/ghoper.jpg")
	f, _ := os.Open(imgPath)
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "jpg", -1, opts)
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

// func TestTransformNegativeImageReader(t *testing.T) {
// 	f, _ := os.Open("../img/ghoper.jpg")
// 	ioutil.ReadAll(f)
// 	opts := WithMode(ModeCombo)
// 	_, err := Transform(f, "jpg", 1, opts)
// 	if err == nil {
// 		t.Error("Expected error but got no error")
// 	}
// }

func TestRunPrimitive(t *testing.T) {
	args := WithMode(ModeCircle)
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/ghoper.jpg")
	outPath := filepath.Join(h, "img/out.jpg")
	_, err := runPrimitive(imgPath, outPath, 1, args...)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}

func TestCreateTempFile(t *testing.T) {
	_, err := createTempFile("", "txt")
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}
func TestCreateTempFileNegative(t *testing.T) {
	_, err := createTempFile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
