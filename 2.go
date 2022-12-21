package main

import (
	"fmt"
	"strings"
)

type RPS byte

const (
	_ RPS = iota
	Rock
	Paper
	Scissors
)

var beats = map[RPS]RPS{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}

func (r RPS) Score(other RPS) int {
	score := 0
	if beats[r] == other {
		score = 6
	} else if r == other {
		score = 3
	}
	return int(r) + score
}

func Day2Part1(input string) (string, error) {
	var score int
	for _, line := range strings.Split(input, "\n") {
		if len(line) != 3 {
			continue
		}
		opponent := RPS(line[0] - 'A' + 1)
		me := RPS(line[2] - 'X' + 1)

		score += me.Score(opponent)
	}
	return fmt.Sprintf("%d", score), nil
}

func Day2Part2(input string) (string, error) {
	var score int
	for _, line := range strings.Split(input, "\n") {
		if len(line) != 3 {
			continue
		}
		opponent := RPS(line[0] - 'A' + 1)
		var me RPS
		switch line[2] {
		case 'X': // loose
			me = beats[opponent]
		case 'Y': // draw
			me = opponent
		case 'Z': // win
			for a, b := range beats {
				if b == opponent {
					me = a
				}
			}
		}

		score += me.Score(opponent)
	}
	return fmt.Sprintf("%d", score), nil
}
