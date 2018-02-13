package daemon

import (
	"net"
	"strings"

	log "github.com/ckeyer/logrus"
	pb "github.com/ckeyer/spongebob/protos"
	"google.golang.org/grpc"
)

type DaemonOption struct {
	Name     string
	HTTPAddr string
	GRPCAddr string
	PromBin  string
}

type Daemon struct {
	httpStatus string
	chStopHTTP chan struct{}
}

func NewDaemon() *Daemon {
	return &Daemon{}
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
