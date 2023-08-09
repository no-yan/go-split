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

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go

	for next := scanner.Scan(); next; {
		bw := bufio.NewWriter(w())
		defer bw.Flush()

		for i := 0; i < lineLim; i++ {
			if next == false {
				if e := scanner.Err(); e != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", e)
				}

				// EOF
				txt := scanner.Text()
				if len(txt) > 0 {
					// In POSIX, EOF should be "\n".
					// If not, add it.
					fmt.Println(bw, txt)
				}
				break
			}

			fmt.Fprintln(bw, scanner.Text())
			next = scanner.Scan()
		}

		if next == false && scanner.Err() == nil {
			break
		}
	}

}
