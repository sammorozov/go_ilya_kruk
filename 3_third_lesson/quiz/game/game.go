package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/questions"
	"strings"
)

func Run(questions []questions.Question) (correctAnswers uint) {
	fmt.Println("Welcome")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}
	return correctAnswers

}

func askQuestion(question questions.Question) bool {
	fmt.Printf("\nEnter the capital of %s: ", question.Country)

	if getUserInput() == strings.ToLower(question.Capital) {
		fmt.Println("correct")
		return true
	} else {
		fmt.Println("incorrect")
		fmt.Println("correct ans is ", question.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("answer: ")
		result, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("error")
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\n"))
	}
}
