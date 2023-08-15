package core

import (
	"bufio"
	"fmt"
	"io"
)

// いったんNのパターンのみ実装
// ラウンドロビン分割、標準出力、レコード分割はサポートしないなどはやらない
func SplitToChunk(r io.Reader, w func() io.Writer, chunk int, fileSize int) error {
	// FIXME: pathはここより上で取得する
	// fileInfo, err := os.Stat("/Users/noyan/tmp/sample.txt")
	// if err != nil {
	// 	log.Fatalf("failed getting information of file: %s", err)
	// }

	// 64-bit OSで int は 64 bits wide なので、int と int64 はお互い問題なくキャストできる
	// 32-bit OSは今回考慮しない
	chunkBytes := fileSize / chunk

	// MB単位で読み込む場合、エラーを出さずにSplitできるかわからないので試す
	scanner := bufio.NewScanner(r)
	scanner.Split(split(chunkBytes))

	// FIXME: bufioは行が65536文字を超えるとエラーが発生する
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
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

	// Write remaining bytes to a single file
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
