package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	days := [][2]func(string) (string, error){
		{Day1Part1, Day1Part2},
	}

	for i, parts := range days {
		file := fmt.Sprintf("%d.txt", i+1)
		input, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		var answers []string
		for part := 0; part < len(parts); part++ {
			answer, err := parts[part](string(input))
			if err != nil {
				return err
			}
			answers = append(answers, answer)
		}
		fmt.Printf("day %d: %s\n", i+1, strings.Join(answers, ", "))
	}
	return nil
}
