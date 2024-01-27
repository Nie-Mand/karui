package utils

import (
	"strings"

	"github.com/Nie-Mand/karui/core/internal"
)

type BasicFormatter struct {}

func (bf *BasicFormatter) Format(s string) string {

	lines := strings.Split(s, "\n")
	parsed := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == internal.COMMENT.String()[0] {
			continue
		}

		if line[len(line) - 1] != ';' && line[len(line) - 1] != '{' && line[len(line) - 1] != '}' {
			line += ";"
		}

		parsed = append(parsed, line)
	}

	s = strings.Join(parsed, "")

	return s
}


