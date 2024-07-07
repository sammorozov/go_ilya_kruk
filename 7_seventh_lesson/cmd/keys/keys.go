package keys

import (
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "Keys",
	Short: "Manage key pairs.",
	Long:  `Use subcommands to create public/private key pairs in PEM files`,
}

// init

func Init(RootCmd *cobra.Command) {
	RootCmd.AddCommand(keysCmd)
}
