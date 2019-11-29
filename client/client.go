package main

import (
    "context"
    "io"
	"log"
    "math/rand"
    "sync"

	"github.com/cnnrznn/docs/document"
	pb "github.com/cnnrznn/docs/editor"
	"google.golang.org/grpc"
)

type Client struct {
    Conn *grpc.ClientConn

    Doc document.Document
    Ec pb.EditorClient
    Id int64

    PushQ []document.Op
    Inflight *document.Op
}

func (c *Client) Tick() {
    if c.Inflight == nil && len(c.PushQ) > 0 {
        c.Inflight = &c.PushQ[0]
        c.Inflight.Version = c.Doc.Version
        c.PushQ = c.PushQ[1:]

        toSend := document.OpConvDoc(*c.Inflight)
        if _, err := c.Ec.Send(context.Background(), &toSend); err != nil {
            log.Panic("Error pushing operation to server")
            // TODO more sofisticated error handling
        }
    }

    // fetch updates since current document version
    stream, err := c.Ec.Recv(context.Background(), &pb.Version{Version: int64(c.Doc.Version)})
    if err != nil {
        log.Panic(err)
    }

    for {
        pbOp, err := stream.Recv()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Panic(err)
        }

        op := document.OpConvPB(*pbOp)

        if pbOp.Sender == c.Id {
            c.Inflight = nil
        } else {
            for i:=0; i<len(c.PushQ); i++ {
                copyOp := op
                op.Transform(c.PushQ[i])
                c.PushQ[i].Transform(copyOp)
            }
        }

        c.Doc.Apply(op)
    }
}

func randomBytes(size int) []byte {
    charset := "abcdefghijklmnopqrstuvwxyz "
    result := make([]byte, size)
    for i:=0; i<size; i++ {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return result
}

func (c *Client) Run(wg *sync.WaitGroup) {
    defer wg.Done()

    for i:=0; i<100; i++ {
        op := document.Op{Sender: int(c.Id),
                    Type: rand.Intn(2),
                    Version: c.Doc.Version,
                    Pos: 0, //int64(rand.Intn(10)),
                    Char: randomBytes(1)[0]}
        //if _, err := c.Ec.Send(context.Background(), &op); err != nil {
        //    log.Printf("Send failed: %v, %v\n", op, err)
        //}

        c.PushQ = append(c.PushQ, op)
        c.Tick()
    }

    // TODO wait a little

    // one final tick
    c.Tick()

    log.Println(&c.Doc)
}

func (c *Client) Close() {
    c.Ec.Leave(context.Background(), &pb.LeaveRequest{Id: c.Id})
    c.Conn.Close()
}

func main() {
    docClient := Client{}

    conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
    if err != nil {
        log.Panic(err)
    }
    defer conn.Close()
    docClient.Conn = conn

    client := pb.NewEditorClient(conn)

    resp, err := client.Join(context.Background(), &pb.JoinRequest{})
    if err != nil {
        log.Panic(err)
    }
    id := resp.Id
    defer client.Leave(context.Background(), &pb.LeaveRequest{Id: id})

    log.Printf("Joined the server as %v\n", id)

    log.Printf("Fetching current state...")

    state, err := client.State(context.Background(), &pb.Nil{})
    if err != nil {
        log.Panic(err)
    }
    docClient.Doc.State = state.Buffer
    docClient.Doc.Version = int(state.Version)

    log.Printf("Done.\n")

    docClient.Ec = client
    docClient.Id = id

    var wg sync.WaitGroup
    wg.Add(1)
    defer wg.Wait()

    docClient.Run(&wg)
    docClient.Close()
}

