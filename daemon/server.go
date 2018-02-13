package daemon

import (
	"log"
	"net"
	"strings"

	pb "github.com/ckeyer/spongebob/protos"
	"google.golang.org/grpc"
)

type DaemonOption struct {
	Name     string
	HTTPAddr string
	GRPCAddr string
	PromBin  string
}

func Serve(opt DaemonOption) (err error) {
	var (
		lis net.Listener
	)
	if strings.HasPrefix(opt.GRPCAddr, "unix") {
		lis, err = net.Listen("unix", strings.TrimPrefix(opt.GRPCAddr, "unix://"))
	} else {
		lis, err = net.Listen("tcp", strings.TrimPrefix(opt.GRPCAddr, "tcp://"))
	}
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHostDaemonServer(s, NewDaemon())

	return s.Serve(lis)
}
