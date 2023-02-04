package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var filename string
	flag.StringVar(&filename, "file", "questions.csv", "Path of csv file having format 'question,answer'")
	flag.Parse()

	f, err := os.Open(filename)

	if err != nil {
		exit(fmt.Sprintf("Error in reading file %s, %v", filename, err))
	}

	// https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer f.Close()

	fileReader := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)

	correct := 0

	content, err := fileReader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error in parsing file %s, %e", filename, err))
	}

	problems := parseProblems(content)

	for i, prob := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, prob.que)

		scanner.Scan()
		ans := scanner.Text()

		if ans == prob.ans {
			correct += 1
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseProblems(content [][]string) []problem {
	problems := make([]problem, len(content))
	for i, prob := range content {
		problems[i] = problem{
			que: prob[0],
			ans: prob[1],
		}
	}
	return problems
}

type problem struct {
	que string
	ans string
}

func exit(err string) {
	log.Fatal(err)
}
