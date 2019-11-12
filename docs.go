package main

import (
	"math/rand"
	"time"

	"github.com/cnnrznn/docs/client"
	"github.com/cnnrznn/docs/document"
)

func main() {
	c := client.New()
	in, out := c.Run()

	for i := 0; i < 100; i++ {
		in <- document.Op{Pos: rand.Intn(i + 1),
			Char: byte(rand.Intn(256)),
			Type: 0}
	}

	time.Sleep(5 * time.Second)

	for i := 0; i < 100; i++ {
		op := <-out
	}
}
