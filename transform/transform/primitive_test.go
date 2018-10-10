package transform

import (
	"os"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

func TestTransform(t *testing.T) {
	inFile, _ := os.Open("/home/neha/dev/src/gophercises/out.png")
	_, err := Transform(inFile, 20)
	if err != nil {
		t.Error("error found", err)
	}

	// _, err = Transform("/home/neha/dev/src/gophercises/out.png", 20)
	// if err != nil {
	// 	t.Error("error found", err)
	// }

}
