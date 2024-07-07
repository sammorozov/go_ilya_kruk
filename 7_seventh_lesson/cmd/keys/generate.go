package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"7_seventh_lesson/crypto_utils"
	"7_seventh_lesson/logger"
	"7_seventh_lesson/utils"

	"github.com/spf13/cobra"
)

type PrivateKeyGen struct {
	outputPath string
	keyBitSize int
	saltSize   int
}

func init() {
	keysCmd.AddCommand(keysGenerateCmd)

	keysGenerateCmd.Flags().String("pub-out", "pub_key.pem", "Path to save the public key")
	keysGenerateCmd.Flags().String("priv-out", "priv_key.pem", "Path to save the private key")
	keysGenerateCmd.Flags().Int("priv-size", 2048, "Private key size in bits")
	keysGenerateCmd.Flags().Int("salt-size", 16, "Salt size used in key derivation in bytes")
}

var keysGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates key pair.",
	Long:  `Generate an RSA key pair and store it in PEM files. The private key will be encrypted using a passphrase that you'll need to enter. AES encryption with Argon2 key derivation function is utilized.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkOut, _ := cmd.Flags().GetString("priv-out")
		pkSize, _ := cmd.Flags().GetInt("priv-size")
		saltSize, _ := cmd.Flags().GetInt("salt-size")

		pkGenConfig := PrivateKeyGen{
			outputPath: pkOut,
			keyBitSize: pkSize,
			saltSize:   saltSize,
		}

		privateKey, err := generatePrivKey(pkGenConfig)
		logger.HaltOnErr(err)

		pubOut, _ := cmd.Flags().GetString("pub-out")
		err = generatePubKey(pubOut, privateKey)
		logger.HaltOnErr(err)
	},
}

func generatePubKey(path string, privKey *rsa.PrivateKey) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %w", err)
	}

	file, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer file.Close()

	if err := pem.Encode(file, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubASN1}); err != nil {
		return fmt.Errorf("failed to encode public key to PEM: %w", err)
	}

	return nil
}

func generatePrivKey(pkGenConfig PrivateKeyGen) (*rsa.PrivateKey, error) {
	absPath, err := filepath.Abs(pkGenConfig.outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, pkGenConfig.keyBitSize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	passphrase, err := utils.GetPassphrase()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	salt, err := makeSalt(pkGenConfig.saltSize)
	if err != nil {
		return nil, err
	}

	key, err := crypto_utils.DeriveKey(crypto_utils.KeyDerivationConfig{
		Passphrase: passphrase,
		Salt:       salt,
	})
	if err != nil {
		return nil, err
	}

	crypter, err := crypto_utils.MakeCrypter(key)
	if err != nil {
		return nil, err
	}

	nonce, err := crypto_utils.MakeNonce(crypter)
	if err != nil {
		return nil, err
	}

	encryptedData := crypter.Seal(nil, nonce, privateKeyBytes, nil)

	// Create a PEM block with the encrypted data
	encryptedPEMBlock := &pem.Block{
		Type:  "ENCRYPTED PRIVATE KEY",
		Bytes: encryptedData,
		Headers: map[string]string{
			"Nonce":                   base64.StdEncoding.EncodeToString(nonce),
			"Salt":                    base64.StdEncoding.EncodeToString(salt),
			"Key-Derivation-Function": "Argon2",
		},
	}

	err = savePrivKeyToPEM(absPath, encryptedPEMBlock)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func savePrivKeyToPEM(absPath string, encryptedPEMBlock *pem.Block) error {
	privKeyFile, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %v", err)
	}
	defer privKeyFile.Close()

	if err := pem.Encode(privKeyFile, encryptedPEMBlock); err != nil {
		return fmt.Errorf("failed to encode private key to PEM: %w", err)
	}

	return nil
}

// makeSalt generates a cryptographic salt.
func makeSalt(saltSize int) ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %v", err)
	}

	return salt, nil
}
