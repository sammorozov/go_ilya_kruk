package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("enter hex or stop to out")
	var input string

	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "stop" {
			break
		}

		i := new(big.Int)

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("failed")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string { // like python typehint
	return strings.TrimPrefix(hexStr, "0x")
}
