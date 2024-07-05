package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/questions"
	"quiz/shuffler"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't load", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)

	correctAnswers := game.Run(questions)

	fmt.Printf("correctAnswers: %d", correctAnswers)
}
