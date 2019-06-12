package main

import (
	"github.com/schollz/progressbar"
	"time"
)

func main() {
	bar := progressbar.New(100)
	for i := 0; i < 100; i++ {
		_ = bar.Add(1)
		time.Sleep(50 * time.Millisecond)
	}
}
