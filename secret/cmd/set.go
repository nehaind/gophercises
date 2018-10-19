package cmd

import (
	"fmt"

	"gophercises/secret/vault"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set command will put api key into secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsPath())
		err := v.Set(args[0], args[1])
		msg := fmt.Sprint("Not able to save in file\n")
		if err == nil {
			msg = fmt.Sprint("saved key success")
		}
		fmt.Println(msg)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
