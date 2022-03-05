package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "File csv dengan format 'soal,jawaban'")
	flag.Parse()
	file, error := os.Open(*csvFileName)
	if error != nil {
		log.Fatal(error)
	}

	r := csv.NewReader(file)

	f, error := r.ReadAll()

	if error != nil {
		log.Fatal(error)
	}

	problems := parseFiles(f)
	totalPoint := 0

	for i, item := range problems {
		fmt.Printf("#%d : %s = \n", i+1, item.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if item.a == ans {
			fmt.Println("CORRECT")
			totalPoint += 1
		} else {
			fmt.Println("INCORRECT")
		}
	}
	fmt.Printf("Your total score is %d", totalPoint)
}

func parseFiles(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
