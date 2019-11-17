package document

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDocument(t *testing.T) {
	doc := New()
	bs := [1]byte{}

	for i := 0; i < 1000; i++ {
		pos := rand.Intn(len(doc.State) + 1)
		typ := rand.Intn(2)
		rand.Read(bs[:])

		doc.Operate(Op{Type: typ,
			Version: i,
			Pos:     pos,
			Char:    bs[0]})
	}

	fmt.Println(doc)
	fmt.Println(doc.Log)
}
