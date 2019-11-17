package client

import (
	"context"

	pb "github.com/cnnrznn/docs/editor"
	"google.golang.org/grpc"
)

type Client struct {
	Stub   pb.EditorClient
	Stream pb.Editor_UpdateClient
}

func New() (c *Client, err error) {
	c = &Client{}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8888", opts...)
	if err != nil {
		return
	}

	client := pb.NewEditorClient(conn)
	if err != nil {
		return
	}

	c.Stub = client
	c.Stream, err = c.Stub.Update(context.Background())

	return
}

func (c *Client) Close() {
	c.Stream.CloseSend()
}
