package main

import (
    "context"
	"fmt"

    "google.golang.org/grpc"
	pb "github.com/cnnrznn/docs/editor"
)

func main() {
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())

    conn, err := grpc.Dial("localhost:8888", opts...)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    client := pb.NewEditorClient(conn)
    state, err := client.State(context.Background(), &pb.Nil{})
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(state)
}
