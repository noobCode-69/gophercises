package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}

func loadQuestions(filename string) ([]Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse csv: %w", err)
	}

	var questions []Question
	for i, row := range rows {
		if len(row) < 2 {
			return nil, fmt.Errorf("invalid row %d: %v", i+1, row)
		}

		q := Question{
			Text:   strings.TrimSpace(row[0]),
			Answer: strings.TrimSpace(row[1]),
		}
		questions = append(questions, q)
	}

	return questions, nil
}

func askQuestions(quiz []Question) int {
	score := 0
	timeLimit := 2

	for _, q := range quiz {
		fmt.Printf("\n%s: ", q.Text)


		answerCh := make(chan string);
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- strings.TrimSpace(input)
		}()

		select {
		case input := <-answerCh:
			if strings.EqualFold(input, q.Answer) {
				score++
			} else {
				fmt.Println("Wrong answer!")
				return score
			}
		case <- time.After(time.Duration(timeLimit)*time.Second):
			fmt.Println("\nTime up!")
    		return score
		}
	}

	return score
}

func main() {
	quiz, err := loadQuestions("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	score := askQuestions(quiz)

	fmt.Printf("\nYour score: %d/%d\n", score, len(quiz))
}
