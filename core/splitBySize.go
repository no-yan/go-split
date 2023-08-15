package core

import (
	"bufio"
	"fmt"
	"io"
)

func split(size int) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if len(data) >= size {
			return size, data[0:size], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func SplitBySize(r io.Reader, w func() io.Writer, size int) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(split((size)))

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		wInstance := w()
		bw := bufio.NewWriter(w())

		txt := scanner.Text()
		if len(txt) > 0 {
			fmt.Fprint(bw, txt)
		}

		bw.Flush()

		// Close the writer instance if it's a WriteCloser
		if closer, ok := wInstance.(io.WriteCloser); ok {
			closer.Close()
		}
	}

	return nil
}
