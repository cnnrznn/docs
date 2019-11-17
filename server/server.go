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

	input chan document.Op
}

func (s *editorServer) run() {
	for {
		select {
		case op := <-s.input:
			s.doc.Operate(op)
		}
	}
}

func (s *editorServer) Update(stream pb.Editor_UpdateServer) error {
	// add client to clients

	for {
		op, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.input <- document.Op{Type: int(op.Type),
			Version: int(op.Version),
			Pos:     int(op.Pos),
			Char:    op.Char[0]}
	}
	return nil
}

func (s *editorServer) State(ctx context.Context, _ *pb.Nil) (*pb.DocState, error) {
	return &pb.DocState{Version: int64(s.doc.Version),
		Buffer: s.doc.State}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	es := editorServer{}
	es.input = make(chan document.Op)

	go es.run()

	grpcServer := grpc.NewServer()
	pb.RegisterEditorServer(grpcServer, &es)
	grpcServer.Serve(lis)
}
