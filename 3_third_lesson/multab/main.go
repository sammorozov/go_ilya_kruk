package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("a\\b\t")

	for _, a := range numbers {
		fmt.Printf("%5d", a)
	}

	fmt.Print("\n\n")

	for _, a := range numbers {
		fmt.Printf("%d\t", a)
		for _, b := range numbers {
			result := a * b
			fmt.Printf("%5d", result)

		}
		fmt.Println()
	}

}
