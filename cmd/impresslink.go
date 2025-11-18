//go:build impresslink

package main

import (
	"runtime"

	"github.com/codeation/itlog/impresslink"
)

func init() { runtime.LockOSThread() }
func main() { impresslink.Exec(root) }
