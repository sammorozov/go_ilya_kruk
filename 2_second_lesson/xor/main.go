package main

// xor --cipher --decipher --secret "secret"

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "default cipher mode")
var secretKey = flag.String("secret", "", "your key")

func main() {

	flag.Parse()

	if len(*secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "no secret key! Exiting...") // here is not placeholders
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plainText := getUserInput("Enter your text to cipher: ")
		cipherdText, err := cipherer.Cipher(plainText, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error encrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipherdText)

	case "decipher":
		cipheredText := getUserInput("Enter your ciphered data to decipher: ")
		decipherText, err := cipherer.Decipher(cipheredText, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error encrypting text: %v\n", err) // here is placeholder
			os.Exit(1)
		}

		fmt.Println(decipherText)
	default:
		fmt.Println("Invalid mode")
		os.Exit(1)

	}

}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {

		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error of reading")
			continue
		}
		return strings.TrimRight(result, "\n")
	}

}
