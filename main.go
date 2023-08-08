package main

import (
	"os"
	"strings"

	"github.com/no-yan/go-split/core"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n")
	core.Split(r, os.Stdout)
}
