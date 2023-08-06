package main

import (
	"fmt"

	"github.com/no-yan/go-split/core"
)

func main() {
	result := core.Split("hello")
	fmt.Println(result)
}
