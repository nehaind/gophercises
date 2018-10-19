package cmd

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var encodingKey string

//RootCmd is the first command to start the execution.
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret is secrets manager CLI Application",
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
