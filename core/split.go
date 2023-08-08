package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Split(r io.Reader, w io.Writer) []string {
	// chunkSize := 1000
	lineLim := 2
	var result []string
	// var fp *os.File

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	scanner := bufio.NewScanner(r)
	// data := make([]byte, 5000)

	for next := scanner.Scan(); next; {
		w := bufio.NewWriter(w)
		for i := 0; i < lineLim; i++ {

			fmt.Fprint(w, scanner.Text(), "\n")
			// EOFでもresultが存在してる
			if next == false {
				if e := scanner.Err(); e != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", e)
					panic(e)
				}
				break
			}
			next = scanner.Scan()
		}
		w.Flush()
	}

	return result
}
