package rpc

import (
	"io"

	log "github.com/ckeyer/logrus"
	pb "github.com/ckeyer/spongebob/protos"
	"golang.org/x/net/context"
)

type ControllerServer struct {
}

func NewControllerServer() *ControllerServer {
	return &ControllerServer{}
}

func (s *ControllerServer) Join(ctx context.Context, node *pb.Node) (*pb.Task, error) {
	return nil, nil
}

func (s *ControllerServer) ReportStatus(stream pb.Controller_ReportStatusServer) error {
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
