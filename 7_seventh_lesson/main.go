package main

import (
	"7_seventh_lesson/cmd"
	"7_seventh_lesson/cmd/keys"
	"7_seventh_lesson/cmd/signatures"
	"7_seventh_lesson/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "Initial setup failed")
	}
}
