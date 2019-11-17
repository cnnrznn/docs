package main

import (
	"context"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"

	"github.com/cnnrznn/docs/document"
	pb "github.com/cnnrznn/docs/editor"
)

type editorServer struct {
	pb.UnimplementedEditorServer

	// In the real world this would loaded from cold storage,
	// authenticated, etc.
	doc document.Document
}

func (s *editorServer) Update(stream pb.Editor_UpdateServer) error {
	for {
		op, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.doc.Log = append(s.doc.Log, document.Op{Pos: op.Pos,
			Version: op.Version,
			Char:    byte(op.Char),
			Type:    op.Type})
	}
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
