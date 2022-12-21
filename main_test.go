package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type AnswerFn func(string) (string, error)

var days = [][2]AnswerFn{
	{Day1Part1, Day1Part2},
}

func TestDays(t *testing.T) {
	for i, parts := range days {
		day := i + 1
		t.Run(fmt.Sprintf("day%d", day), func(t *testing.T) {
			pattern := filepath.Join("testdata", fmt.Sprintf("%d.*.txt", day))
			files, err := filepath.Glob(pattern)
			require.NoError(t, err)

			var testFile TestFile
			for _, file := range files {
				require.NoError(t, testFile.Load(file))
				t.Run(testFile.Name, func(t *testing.T) {
					for j, partFn := range parts[:] {
						part := j + 1
						answer, err := partFn(testFile.Input)
						require.NoError(t, err)
						if answer != testFile.Answers[j] {
							t.Errorf("part=%d got=%s want=%s", part, answer, testFile.Answers[j])
						}
					}
				})
			}
		})
	}
}

type TestFile struct {
	Name    string
	Input   string
	Answers [2]string
}

func (t *TestFile) Load(name string) error {
	t.Name = filepath.Base(name)
	t.Name = strings.TrimSuffix(t.Name, filepath.Ext(t.Name))
	t.Answers = [2]string{"", ""}
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	input := strings.SplitN(string(data), "---\n", 2)
	switch len(input) {
	case 1:
		t.Input = input[0]
	case 2:
		copy(t.Answers[:], strings.Split(strings.TrimSpace(input[0]), " "))
		t.Input = input[1]
	default:
		return fmt.Errorf("bad input: %q", input)
	}

	return nil
}
