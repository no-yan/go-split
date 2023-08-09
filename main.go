package main

import (
	"fmt"
	"io"
	"os"

	"github.com/no-yan/go-split/core"
)

func main() {
	f, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println("cannot open the file")
	}
	defer f.Close()

	core.Split(f, func() io.Writer { return os.Stdout })
}
