package core

import (
	"bufio"
	"fmt"
	"io"
)

type NewWriterFunc func() io.WriteCloser

func SplitByLine(r io.Reader, newWriter NewWriterFunc, n int) error {
	scanner := bufio.NewScanner(r)

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する。
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	// buf := make([]byte, 0, 64*1024)
	// scanner.Buffer(buf, 10*64*1024) // 例: 10倍のサイズに設定

	var bw *bufio.Writer
	var count int

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}

		if count%n == 0 {
			if bw != nil {
				if err := bw.Flush(); err != nil {
					return err
				}
			}
			bw = bufio.NewWriter(newWriter())
		}

		txt := scanner.Text()
		fmt.Fprintln(bw, txt)
		count++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	if bw != nil {
		if err := bw.Flush(); err != nil {
			return err
		}
	}

	return nil
}
