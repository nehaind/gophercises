package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//Encrypt function
func Encrypt(key, text string) (string, error) {

	cipherKey, _ := newCipherBlock(key)

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err == nil {
		stream := cipher.NewCFBEncrypter(cipherKey, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	}

	return fmt.Sprintf("%x", ciphertext), nil
}

//Decrypt will decrypt the key value
func Decrypt(key, cipherHex string) (string, error) {
	var ciphertext []byte
	block, err := newCipherBlock(key)
	if err == nil {
		ciphertext, err = hex.DecodeString(cipherHex)
		if err == nil {
			if len(ciphertext) >= aes.BlockSize {
				iv := ciphertext[:aes.BlockSize]
				ciphertext = ciphertext[aes.BlockSize:]

				stream := cipher.NewCFBDecrypter(block, iv)

				// XORKeyStream can work in-place if the two arguments are the same.
				stream.XORKeyStream(ciphertext, ciphertext)
			}
		}
	}
	return string(ciphertext), err
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
