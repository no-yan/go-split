package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Split(s string) []string {
	// chunkSize := 1000
	lineLim := 2
	var result []string
	// var fp *os.File
	r := strings.NewReader("some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n")
	scanner := bufio.NewScanner(r)
	// data := make([]byte, 5000)

	readMore := true
	for readMore {
		w := bufio.NewWriter(os.Stdout)
	exit_loop:
		for i := 0; i < lineLim; i++ {
			fmt.Println(i)
			next := scanner.Scan()
			fmt.Fprint(w, scanner.Text())
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
		fmt.Fprint(w, "\n")
		w.Flush()
	}

	return result
}
