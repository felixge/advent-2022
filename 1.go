package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1Part1(input string) (string, error) {
	var max int64
	var current int64
	for _, line := range strings.Split(input+"\n", "\n") {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return "", err
			}
			current += num
		}
	}
	return fmt.Sprintf("%d", max), nil
}

func Day1Part2(input string) (string, error) {
	var max = make([]int64, 3)
	var current int64
	for _, line := range strings.Split(input+"\n", "\n") {
		if line == "" {
			for i, v := range max {
				if current > v {
					max[i] = current
					sort.Slice(max, func(i, j int) bool { return max[i] < max[j] })
					break
				}
			}
			current = 0
		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return "", err
			}
			current += num
		}
	}
	var sum int64
	for _, v := range max {
		sum += v
	}
	return fmt.Sprintf("%d", sum), nil
}
