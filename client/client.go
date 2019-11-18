package client

import (
	"context"
	"io"

	"github.com/cnnrznn/docs/document"
	pb "github.com/cnnrznn/docs/editor"
	"google.golang.org/grpc"
)

type Client struct {
	stub   pb.EditorClient
	stream pb.Editor_UpdateClient

	inflight *document.Op
	queue    []document.Op

	in, out    chan document.Op // interface to user
	fromServer chan document.Op // recv op from server
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

	c.stub = client
	c.stream, err = c.stub.Update(context.Background())
	if err != nil {
		return
	}

	c.in, c.out = make(chan document.Op), make(chan document.Op)
	c.fromServer = make(chan document.Op)

	go func() {
		for {
			op, err := c.stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
			c.fromServer <- document.Op{Version: int(op.Version),
				Type: int(op.Type),
				Char: op.Char[0],
				Pos:  int(op.Pos)}
		}
	}()

	return
}

func (c *Client) Close() {
	c.stream.CloseSend()
}

func (c *Client) push() {
	if c.inflight != nil || len(c.queue) == 0 {
		return
	}

	c.inflight = &c.queue[0]
	c.queue = c.queue[1:]

	op := *c.inflight

	c.stream.Send(&pb.Op{Version: int64(op.Version),
		Pos:  int64(op.Pos),
		Type: int32(op.Type),
		Char: []byte{op.Char}})
}

func (c *Client) Run() {
	for {
		select {
		case op := <-c.in:
			c.queue = append(c.queue, op)
			c.push()
		case op := <-c.fromServer:
			if c.inflight != nil && c.inflight.Sender == op.Sender {
				c.inflight = nil
				c.push()
			} else {
				// transform op against c.queue
				// transform c.queue against op
				for i, qop := range c.queue {
					copyOp := op
					copyOp.Transform(qop)
					c.queue[i].Transform(op)
					op = copyOp
				}

				c.out <- op
			}
		}
	}
}
