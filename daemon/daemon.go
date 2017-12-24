package daemon

import (
	"fmt"
	"net"
	"strings"

	log "github.com/ckeyer/logrus"
	pb "github.com/ckeyer/spongebob/protos"
	"google.golang.org/grpc"
)

type Daemon struct {
	httpStatus string
	chStopHTTP chan struct{}
}

func NewDaemon() *Daemon {
	return &Daemon{}
}

func Serve(addr string) (err error) {
	var (
		lis net.Listener
	)
	if strings.HasPrefix(addr, "unix") {
		lis, err = net.Listen("unix", strings.TrimPrefix(addr, "unix://"))
	} else if strings.HasPrefix(addr, "tcp") {
		lis, err = net.Listen("tcp", strings.TrimPrefix(addr, "tcp://"))
	} else {
		return fmt.Errorf("invalid listenning address, %s", addr)
	}
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHostDaemonServer(s, NewDaemon())

	return s.Serve(lis)
}
