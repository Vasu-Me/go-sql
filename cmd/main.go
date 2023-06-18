package main

import (
	pck "github.com/Vasu-Me/gosql/pck"
)

func main() {
	mb := pck.NewMemoryBackend()

	pck.RunRepl(mb)
}