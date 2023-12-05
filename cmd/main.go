package main

import (
	"flag"
	"flashslothmor3/quiz"
	"fmt"
	"os"
)

func main() {
	//use flag packgae
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	quizzes := &quiz.List{}

	if err := quizzes.Load(*csvFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	quizzes.Play(*limit)
}
