package cmd

import (
	"fmt"

	"gophercises/secret/vault"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get command will return api key from secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsPath())
		value, err := v.Get(args[0])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return
		}
		fmt.Println(value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
