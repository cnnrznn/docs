package main

import (
    "fmt"

    "github.com/cnnrznn/docs/document"
)

func main() {
    doc := document.New()

    doc.Insert('\n', 0, 0);
    doc.Insert('C', 1, 1);
    doc.Insert('o', 2, 2);
    doc.Insert('n', 3, 3);
    doc.Insert('n', 4, 4);

    fmt.Println(doc)
}
