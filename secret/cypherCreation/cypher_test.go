package cypher

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

// func TestEcrypt(t *testing.T) {
// 	testKey := "key1"
// 	text := "password"
// 	output, err := Encrypt(testKey, text)
// 	if err != nil {
// 		t.Error("error in encrypt")
// 	}
// 	plainText, err := Decrypt(testKey, output)
// 	if err != nil {
// 		t.Error("error in decrypt")
// 	}
// 	if output != plainText {
// 		t.Errorf("error in both value not match")
// 	}
// }

func TestCipher(t *testing.T) {
	testSuit := []struct {
		key       string
		plainText string
	}{
		// {key: "demo1", plainText: "123abc"},
		// {key: "demo2", plainText: "123456"},
		// {key: "demo3", plainText: "asdfg"},
		{key: "key1", plainText: "password"},
	}
	for i, test := range testSuit {
		hex, err := Encrypt(test.key, test.plainText)
		if err != nil {
			t.Error("error in encrypt")
		}
		plainText, err := Decrypt(test.key, hex)
		if err != nil {
			t.Error("error in decrypt")
		}
		if test.plainText != plainText {
			t.Errorf("error in both value not match %d", i)
		}
	}
}

//output, _ := Encrypt("abc", "password")
// fmt.Println(output)
// content := []byte(output)
// f, err := os.OpenFile("/home/neha/dev/src/gophercises/secret/cypherCreation/test.txt", os.O_CREATE|os.O_RDWR, 0666)
// if err != nil {
// 	fmt.Println("file not found")
// }
// if _, err := f.Write(content); err != nil {
// 	fmt.Println("write to file")
// 	log.Fatal(err)
// }
// if err := f.Close(); err != nil {
// 	log.Fatal(err)
// }
// f.Truncate(0)
// f.Close()
