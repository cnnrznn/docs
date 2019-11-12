package client

import (
	"fmt"

	"github.com/cnnrznn/docs/document"
)

type Client struct {
	Doc document.Document
	Ops []document.Op
}

func New() *Client {
	return &Client{}
}

func (c *Client) Run() (input, output chan document.Op) {
	input, output = make(chan document.Op, 1024), make(chan document.Op, 1024)

	// TODO some initialization here with the server (current doc state, etc.)

	go c.serve(input, output)

	return input, output
}

func (c *Client) serve(input, output chan document.Op) {
	for {
		select {
		case op := <-input:
			fmt.Println("Received Op", op)
			c.Ops = append(c.Ops, op)
			// do another case here for listening to the server
		}
	}
}
