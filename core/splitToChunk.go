package core

import (
	"bufio"
	"fmt"
	"io"
)

func SplitToChunk(r io.Reader, w NewWriterFunc, chunk int, fileSize int) error {
	// 64-bit OSで int は 64 bits wide なので、int と int64 はお互い問題なくキャストできる
	// 32-bit OSは今回考慮しない
	chunkBytes := fileSize / chunk

	scanner := bufio.NewScanner(r)
	scanner.Split(split(chunkBytes))
	// MB, GB単位で読み込む場合、 バッファサイズを上げる
	if fileSize/chunk >= 65536 {
		buf := make([]byte, fileSize/chunk)
		scanner.Buffer(buf, fileSize/chunk)
	}

	count := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		bw := bufio.NewWriter(w())

		if txt := scanner.Text(); len(txt) > 0 {
			fmt.Fprint(bw, txt)
		}

		bw.Flush()
		count++
		if count >= chunk-1 {
			break
		}
	}

	// 入力の残りがあれば、1つのファイルに書き込む
	var lastBw *bufio.Writer
	for scanner.Scan() {
		if lastBw == nil {
			lastBw = bufio.NewWriter(w())

		}
		if err := scanner.Err(); err != nil {
			return err
		}

		if txt := scanner.Text(); len(txt) > 0 {
			fmt.Fprint(lastBw, txt)
		}
	}
	if lastBw != nil {
		if err := lastBw.Flush(); err != nil {
			return err
		}
	}

	return nil
}
