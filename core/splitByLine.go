package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type NewWriterFunc func() io.Writer

func SplitByLine(r io.Reader, w NewWriterFunc, n int) {
	scanner := bufio.NewScanner(r)

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go

	for scanner.Scan() {
		bw := bufio.NewWriter(w())

		for i := 0; i < n; i++ {
			if e := scanner.Err(); e != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", e)
			}

			txt := scanner.Text()
			if len(txt) == 0 {
				break
			}
			fmt.Fprintln(bw, txt)
		}

		if scanner.Err() != nil {
			// TODO:
		}
		bw.Flush()
	}

}
