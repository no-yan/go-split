package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/no-yan/go-split/core"
)

var lineCount = flag.Int("l", 1000, "-l line_count\nCreate split files line_count lines in length.")
var chunkCount = flag.Int("n", 0, " -n chunk_count\nSplit file into chunk_count smaller files.  The first n - 1 files\nwill be of size (size of file / chunk_count ) and the last file\nwill contain the remaining bytes.")
var byteCount = flag.Int("b", 0, "-b byte_count[K|k|M|m|G|g]\nCreate split files byte_count bytes in length.  If k or K is\nappended to the number, the file is split into byte_count\nkilobyte pieces.  If m or M is appended to the number, the file\nis split into byte_count megabyte pieces.  If g or G is appended\nto the number, the file is split into byte_count gigabyte pieces.\n")

// TODO: コマンドの仕様を調べる
// ./go-split [options] [file [prefix]]

type Res struct {
	// TODO: サイズ考える
	line        *int
	chunk_count *int
	byteCount   *int
}

func main() {
	flag.Parse()
	path := flag.Arg(0)

	var file io.Reader

	// ファイルが存在しない場合、標準入力から受け取る
	switch path {
	case "", "-":

		file := bufio.NewReader(os.Stdin)
		defer os.Stdin.Close()

		// 標準入力が空であれば、ミスの可能性が高い
		_, err := file.Peek(1)
		if err != nil {
			fmt.Print("Stdin is empty. Are you specifying the command in the wrong way?")
			log.Fatal(err)
		}
	default:
		// TODO: Openで特殊なpathを指定された時の挙動を調べる(普通にやると開かないpathがあるか)
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		file = f
	}

	// fmt.Println(*chunkCount, *lineCount, *byteCount)

	// Res{lineCount, chunkCount, byteCount}

	core.SplitBySize(file, core.GenerateNextWriter(), 5)
}
