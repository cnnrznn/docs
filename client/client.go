package main

import (
    "context"
    //"fmt"
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

    PushQ []pb.Op
    Inflight *pb.Op
}

func (c *Client) Tick() {
    if c.Inflight == nil && len(c.PushQ) > 0 {
        c.Inflight = &c.PushQ[0]
        c.PushQ = c.PushQ[1:]

        if _, err := c.Ec.Send(context.Background(), c.Inflight); err != nil {
            log.Panic("Error pushing operation to server")
            // TODO more sofisticated error handling
        }
    }

    // fetch updates since current document version
    // resolve updates agains pushq
    // if updates match inflight, make inflight nil
    // apply updates to local document
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
        op := pb.Op{Sender: c.Id,
                    Type: int32(rand.Intn(2)),
                    Version: int64(i),
                    Pos: int64(rand.Intn(100)),
                    Char: []byte{randomBytes(1)[0]}}
        //if _, err := c.Ec.Send(context.Background(), &op); err != nil {
        //    log.Printf("Send failed: %v, %v\n", op, err)
        //}

        c.PushQ = append(c.PushQ, op)
        c.Tick()
    }

    // TODO wait a little

    // one final tick
    c.Tick()
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

