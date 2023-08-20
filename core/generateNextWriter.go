package core

import (
	"io"
	"os"
)

func nextAlphabet(prev string) string {
	if prev == "" {
		return "a"
	}

	headRunes, trailingRune := prev[:len(prev)-1], prev[len(prev)-1]
	if trailingRune == 'z' {
		return nextAlphabet(headRunes) + "a"
	}
	return headRunes + string(trailingRune+1)
}

func fileGenerator(prefix string) func() (*os.File, error) {
	fileName := ""
	return func() (*os.File, error) {
		fileName = nextAlphabet(fileName)
		return os.Create(prefix + fileName)
	}
}

func GenerateNextWriter(prefix string) func() io.WriteCloser {
	if prefix == "" {
		prefix = "x"
	}
	createFile := fileGenerator(prefix)
	generator := func() io.WriteCloser {
		file, err := createFile()
		if err != nil {
			panic(err)
		}
		return file
	}
	return generator
}
