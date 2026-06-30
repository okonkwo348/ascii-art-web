package logic

import (
	"errors"
	"os"
	"strings"
)

func Generate(text string, filename string) (string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("couldn't load file")
	}

	content := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(content) == 0 {
		return "", errors.New("file is empty")
	}

	if len(content) != 856 {
		return "", errors.New("file lines incomplete")
	}

	sliceData := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var b strings.Builder
	var result []string

	for i, word := range sliceData {
		if word == "" {
			if i == len(sliceData) {
				continue
			}
			b.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range word {
				if ch < 32 || ch > 126 {
					return "", errors.New("Input contains non-ascii characters")
				}
				lines := int(ch-32)*9 + 1 + i
				b.WriteString(content[lines])
			}
			result = append(result, b.String())
			b.Reset()
		}

	}

	return strings.Join(result, "\n"), nil

}
