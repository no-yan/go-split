package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type NewWriterFunc func() io.Writer

func Split(r io.Reader, w NewWriterFunc) {
	lineLim := 5000

	scanner := bufio.NewScanner(r)
	bw := bufio.NewWriter(w())
	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	next := scanner.Scan()
	defer bw.Flush()

	for i := 0; i < lineLim; i++ {
		e := scanner.Err()
		if next == false {
			if e != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", e)
			}

			// EOF
			txt := scanner.Text()
			fmt.Fprint(bw, txt)
			if len(txt) != 0 {
				fmt.Fprint(bw, "/n")
			}
			break
		}

		fmt.Fprintln(bw, scanner.Text())
		next = scanner.Scan()
	}

}
