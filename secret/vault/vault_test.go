package vault

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}

func getSecretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
func TestSet(t *testing.T) {
	testStruct := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "123", filepath: getSecretpath(), key: "password", plainText: "value"},
	}
	for _, test := range testStruct {
		v := GetVault(test.encodingKey, test.filepath)
		err := v.Set(test.key, test.plainText)
		if err != nil {
			t.Error("found error")
		}
	}
}

func TestGet(t *testing.T) {
	testStruct := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "123", filepath: getSecretpath(), key: "password", plainText: "value"},
		{encodingKey: "123", filepath: getSecretpath() + "ds", key: "password", plainText: ""},
		{encodingKey: "123", filepath: getSecretpath(), key: "whichisnotpresent", plainText: ""},
	}
	for _, test := range testStruct {
		v := GetVault(test.encodingKey, test.filepath)
		plainText, _ := v.Get(test.key)
		if plainText != test.plainText {
			t.Error("error in Get")
		}
	}
}
