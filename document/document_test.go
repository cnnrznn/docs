package document

import (
	"fmt"
	"math/rand"
	"testing"
    //"time"
)

func randomByte() byte {
    charset := "abcdefghijklmnopqrstuvwxyz "
    return charset[rand.Intn(len(charset))]
}

func TestDocument(t *testing.T) {
	doc := New()

	for i := 0; i < 20; i++ {
		pos := rand.Intn(len(doc.State) + 1)
		typ := rand.Intn(2)

		doc.Operate(Op{Type: typ,
			Version: i,
			Pos:     pos,
			Char:    randomByte()})

        fmt.Println(doc.Log)
        fmt.Println(doc)
        //time.Sleep(1 * time.Second)
	}
}
