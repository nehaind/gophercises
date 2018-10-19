package vault

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"gophercises/secret/encrypt"
)

// Vault is struct used to store API's keys
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// GetVault is used to get vault struct to perform get set operation
func GetVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	return dec.Decode(&v.keyValues)

}

func (v *Vault) writer(f io.Writer) error {
	enc := json.NewEncoder(f)
	return enc.Encode(v.keyValues)
}

// Set function is used to set the key value in  file
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.load()
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err == nil {
		defer f.Close()
		hex, err := encrypt.Encrypt(v.encodingKey, value)
		if err == nil {
			v.keyValues[key] = hex
			err = v.writer(f)
		}
	}
	return err

}

// Get is used to get the value of the encrypted value in the file
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", errors.New("File not found")
	}
	hex, ok := v.keyValues[key]
	if ok != true {
		return "", errors.New("Key not found")
	}
	return encrypt.Decrypt(v.encodingKey, hex)
}
