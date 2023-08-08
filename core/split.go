package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Split(r io.Reader, w io.Writer) {
	lineLim := 5000

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	scanner := bufio.NewScanner(r)
	// data := make([]byte, 5000)
label:
	for next := scanner.Scan(); next; {
		w := bufio.NewWriter(w)
		defer w.Flush()

		for i := 0; i < lineLim; i++ {
			e := scanner.Err()
			if next == false {
				if e != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", e)
				}
				fmt.Fprint(w, scanner.Text())
				break label
			}

			fmt.Fprint(w, scanner.Text(), "\n")
			next = scanner.Scan()
		}

	}

}
