package signatures

import (
	"github.com/spf13/cobra"
)

var signaturesCmd = &cobra.Command{
	Use:   "signatures",
	Short: "Create and verify signatures.",
	Long:  `Use subcommands to create signature (.sig) with private key and vrify`,
}

func Init(RootCmd *cobra.Command) {
	RootCmd.AddCommand(signaturesCmd)
}
