package server

import (
	"io"
	"net"

	log "github.com/Sirupsen/logrus"
	pb "github.com/ckeyer/spongebob/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	Name string
}

func (s *Server) Join(ctx context.Context, node *pb.Node) (*pb.Task, error) {
	return nil, nil
}

func (s *Server) ReportStatus(stream pb.Controller_ReportStatusServer) error {
	for {
		ns, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		log.Debugf("receive node status %+v", ns)

		task := &pb.Task{
			Name: "test task",
		}
		stream.Send(task)

	}
	return nil
}

func Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterControllerServer(s, &Server{})

	return s.Serve(lis)
}
