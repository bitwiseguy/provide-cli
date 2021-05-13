package contracts

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const contractTypeRegistry = "registry"

var contract map[string]interface{}
var contracts []interface{}
var contractType string

var ContractsCmd = &cobra.Command{
	Use:   "contracts",
	Short: "Manage smart contracts",
	Long:  `Compile and deploy smart contracts locally from source or execute previously-deployed contracts`,
	Run: func(cmd *cobra.Command, args []string) {
		generalPrompt(cmd, args, "")

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Prompt Exit\n")
				os.Exit(1)
			}
		}()
	},
}

func init() {
	ContractsCmd.AddCommand(contractsListCmd)
	ContractsCmd.AddCommand(contractsExecuteCmd)
}
