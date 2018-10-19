package encrypt

import (
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}
func TestCipher(t *testing.T) {
	testSuit := []struct {
		key       string
		plainText string
	}{
		{key: "password", plainText: "encodevalue"},
		{key: "password", plainText: "encodevalue"},
		{key: "password", plainText: "encodevalue"},
		{key: "password", plainText: "encodevalue"},
	}
	for i, test := range testSuit {
		hex, err := Encrypt(test.key, test.plainText)
		if err != nil {
			t.Error("error found encrypt function")
		}
		plainText, err := Decrypt(test.key, hex)
		if err != nil {
			t.Error("error found in decrypt function")
		}
		if test.plainText != plainText {
			t.Errorf("error in both value not match %d", i)
		}
	}
}
