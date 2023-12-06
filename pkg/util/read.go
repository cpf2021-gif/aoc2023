package util

import (
	"os"
	"strings"
)

func ReadByLine(filename string) []string {
	context, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(context), "\n")

	return lines
}
