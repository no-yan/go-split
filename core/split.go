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

	readMore := true
	for readMore {
		w := bufio.NewWriter(w)
	exit_loop:
		for i := 0; i < lineLim; i++ {
			next := scanner.Scan()
			fmt.Fprint(w, scanner.Text(), "\n")
			// EOFでもresultが存在してる
			if next == false {
				// 複数のエラーケースはどう処理する？
				switch e := scanner.Err(); e {
				case nil:
					readMore = false
					break exit_loop
				default:
					fmt.Fprintln(os.Stderr, "reading standard input:", e)
				}
			}
		}
		w.Flush()
	}

	return result
}
