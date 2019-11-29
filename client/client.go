package main

import (
    "context"
    "fmt"
	"log"

	"github.com/cnnrznn/docs/document"
	pb "github.com/cnnrznn/docs/editor"
	"google.golang.org/grpc"
)

type Client struct {
    doc document.Document
}

func main() {
    docClient := Client{}

    conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
    if err != nil {
        log.Panic(err)
    }
    defer conn.Close()

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
    docClient.doc.State = state.Buffer
    docClient.doc.Version = int(state.Version)

    log.Printf("Done.\n")

    fmt.Println(docClient.doc)
}

