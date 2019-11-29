package main

import (
	"context"
	"fmt"
    "log"
    "math/rand"
	"net"
    "sync"

	"google.golang.org/grpc"

	"github.com/cnnrznn/docs/document"
	pb "github.com/cnnrznn/docs/editor"
)

type editorServer struct {
	pb.UnimplementedEditorServer

    ids map[int64]bool

	// In the real world this would loaded from cold storage,
	// authenticated, etc.
	doc document.Document
    mux sync.Mutex
}

func (s *editorServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
    s.mux.Lock()
    defer s.mux.Unlock()

    var id int64

    for {
        id = rand.Int63()
        if _, ok := s.ids[id]; !ok {
            s.ids[id] = true
            break
        }
    }

    log.Printf("%v joined the server\n", id)

    return &pb.JoinResponse{Id: id}, nil
}

func (s *editorServer) Leave(ctx context.Context, req *pb.LeaveRequest) (*pb.LeaveResponse, error) {
    s.mux.Lock()
    defer s.mux.Unlock()

    if _, ok := s.ids[req.Id]; ok {
        s.ids[req.Id] = false
        log.Printf("%v left the server\n", req.Id)
    } else {
        log.Printf("Invalid leave request for %v\n", req.Id)
    }

    return &pb.LeaveResponse{}, nil
}

func (s *editorServer) State(ctx context.Context, _ *pb.Nil) (*pb.DocState, error) {
    s.mux.Lock()
    defer s.mux.Unlock()

	return &pb.DocState{Version: int64(s.doc.Version),
		Buffer: s.doc.State}, nil
}

func (s *editorServer) Send(ctx context.Context, req *pb.Op) (*pb.Nil, error) {
    s.mux.Lock()
    s.doc.Operate(document.OpConvPB(*req))
    s.mux.Unlock()

    log.Println(len(s.doc.State), &s.doc)

    return &pb.Nil{}, nil
}

func (s *editorServer) Recv(version *pb.Version, stream pb.Editor_RecvServer) error {
    s.mux.Lock()
    defer s.mux.Unlock()

    for i:=int(version.Version); i<len(s.doc.Log); i++ {
        pbOp := document.OpConvDoc(s.doc.Log[i])
        if err := stream.Send(&pbOp); err != nil {
            return err
        }
    }

    return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	es := editorServer{}
    es.ids = make(map[int64]bool)

	grpcServer := grpc.NewServer()
	pb.RegisterEditorServer(grpcServer, &es)
	grpcServer.Serve(lis)
}
