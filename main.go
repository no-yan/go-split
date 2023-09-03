package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/no-yan/go-split/core"
)

var lineCount = flag.Int("l", 1000, "-l line_count\nCreate split files line_count lines in length.")
var chunkCount = flag.Int("n", 0, " -n chunk_count\nSplit file into chunk_count smaller files.  The first n - 1 files\nwill be of size (size of file / chunk_count ) and the last file\nwill contain the remaining bytes.")
var byteCount = flag.Int("b", 0, "-b byte_count\nCreate split files byte_count bytes in length.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file [prefix]]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if err := validateFlags(*lineCount, *chunkCount, *byteCount); err != nil {
		log.Fatalf("CLI error: %s", err)
	}
	path := flag.Arg(0)
	file, fileSize, err := getReader(path)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	prefix := flag.Arg(1)

	switch {
	case *chunkCount != 0:
		err := core.SplitToChunk(file, core.GenerateNextWriter(prefix), *chunkCount, int(fileSize))
		if err != nil {
			log.Fatalln(err)
		}
	case *byteCount != 0:
		err := core.SplitByByte(file, core.GenerateNextWriter(prefix), *byteCount)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		if file == nil {
			fmt.Println("file is nil")
		}
		err := core.SplitByLine(file, core.GenerateNextWriter(prefix), *lineCount)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func validateFlags(line, chunk, byte int) error {
	if *lineCount != 1000 && *chunkCount != 0 || *chunkCount != 0 && *byteCount != 0 || *byteCount != 0 && *lineCount != 1000 {
		return errors.New("Options 'l', 'n', and 'b' cannot be specified simultaneously. Please choose only one.")
	}
	if *lineCount < 0 || *chunkCount < 0 || *byteCount < 0 {
		return errors.New("A negative value was entered. Please input a positive integer only.")
	}
	return nil
}

func getReader(path string) (io.Reader, int64, error) {

	var file io.Reader
	// io.Readerからファイルサイズを取得できないため、この関数で同時に返却する
	var fileSize int64

	switch path {
	// If file is a single dash (‘-’) or absent, split reads from the standard input.
	case "", "-":
		f := os.Stdin
		defer os.Stdin.Close()

		fi, err := f.Stat()
		if err != nil {
			return nil, 0, err
		}

		file = f
		fileSize = fi.Size()

		if fileSize == 0 {
			// 標準入力のサイズが0であれば、コマンドの入力ミスの可能性が高い
			// https://stackoverflow.com/a/22564526
			err := errors.New("Stdin is empty. Are you specifying the command in the wrong way?")
			return nil, 0, err
		}

	default:
		f, err := os.Open(path)
		if err != nil {
			return nil, 0, err
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			return nil, 0, err
		}

		file = f
		fileSize = fi.Size()
	}

	return file, fileSize, nil
}
