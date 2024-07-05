package questions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func LoadQuestions() ([]Question, error) {
	jsonData, err := os.ReadFile("quiz.json")
	if err != nil {
		return nil, fmt.Errorf("error reading %w", err)
	}

	var questions []Question

	err = json.Unmarshal(jsonData, &questions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshing JSON %w", err)
	}
	return questions, nil
}
