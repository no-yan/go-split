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

// ファイル名生成の責務ををGenerateNextWriterから分離したが、ジェネレーター使いすぎているかもしれない
// ジェネレーターのネストは可読性が落ちそう
func fileGenerator() func() (*os.File, error) {
	fileName := ""
	prefix := "x"
	return func() (*os.File, error) {
		fileName = nextAlphabet(fileName)
		return os.Create(prefix + fileName)
	}
}

func GenerateNextWriter() func() io.Writer {
	createFile := fileGenerator()
	generator := func() io.Writer {
		file, err := createFile()
		if err != nil {
			panic(err)
		}
		return file
	}
	return generator
}
