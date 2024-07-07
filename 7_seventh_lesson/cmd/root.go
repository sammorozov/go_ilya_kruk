package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "brave_signer",
		Short: "Bravely generate key pairs, sign files, and check signatures.",
		Long:  `A collection of tools to generate key pairs in PEM files, sign files, and verify signatures.`,
	}
}
