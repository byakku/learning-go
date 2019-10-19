package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("c sv", "problems.csv", "csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse CSV")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
problemloop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == problem.question {
				correct++
			}
		}
	}
}

func parseLines(lines [][]string) []problem {
	problemSlice := make([]problem, len(lines))

	for i, line := range lines {
		problemSlice[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problemSlice
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
