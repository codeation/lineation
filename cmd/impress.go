//go:build !impresslink

package main

import (
	"fmt"

	"github.com/codeation/impress/duo/duodriver"
)

func main() {
	d, err := duodriver.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	root(d)
}
