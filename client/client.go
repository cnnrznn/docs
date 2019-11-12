package client

import (
	"github.com/cnnrznn/docs/document"
)

type Client struct {
	Doc Document
	// probably some other stuff like outgoing/incoming queue
}

func New() *Client {
	return &Client{}
}
