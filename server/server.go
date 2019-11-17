package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"

    "github.com/cnnrznn/docs/document"
    pb "github.com/cnnrznn/docs/editor"
)

type editorServer struct {
    pb.UnimplementedEditorServer

    doc document.Document
}

func (s *editorServer) Update(stream pb.Editor_UpdateServer) error {
    return nil
}

func (s *editorServer) State(ctx context.Context, _ *pb.Nil) (*pb.DocState, error) {
    return &pb.DocState{Version: s.doc.Version,
                        Buffer: s.doc.State}, nil
}

func main() {
    lis, err := net.Listen("tcp", "localhost:8888")
    if err != nil {
        fmt.Println(err)
        return
    }

    grpcServer := grpc.NewServer()
    pb.RegisterEditorServer(grpcServer, &editorServer{})
    grpcServer.Serve(lis)
}
