package utils

import (
	"fmt"
	"os"
)

func GetPassphrase() ([]byte, error) {
	fmt.Println("Enter passphrase")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}
	defer safeRestore(int(os.Stdin.Fd()), oldState)

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read passphrase: %w", err)
	}

	return passphrase, nil
}
