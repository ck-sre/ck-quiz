package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	prbFileName := flag.String("csv", "questions.csv", "CSV for q n a")
	flag.Parse()

	f, err := os.Open(*prbFileName)
	if err != nil {
		quitFunc("Couldn't open file")
	}

	r := csv.NewReader(f)
	allLines, err := r.ReadAll()
	if err != nil {
		quitFunc("Failed to pase the file")
	}
	allQAs := getQAs(allLines)

	score := 0

	for i, qa := range allQAs {
		fmt.Printf("Question #%d: %s\n", i+1, qa.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == qa.a {
			fmt.Println("Correct")
			score++
		}
	}

	fmt.Printf("Score is %d out of %d\n", score, len(allQAs))

}

type questA struct {
	q string
	a string
}

func getQAs(records [][]string) []questA {
	allQAs := make([]questA, len(records))
	for i, record := range records {
		allQAs[i] = questA{
			q: strings.TrimSpace(record[0]),
			a: strings.TrimSpace(record[1]),
		}
	}
	return allQAs
}

func quitFunc(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
