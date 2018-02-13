package daemon

import (
	pb "github.com/ckeyer/api"
	"golang.org/x/net/context"
)

const (
	statusRun = "running"
)

func (d *Daemon) StartHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {

	return nil, nil
}

func (d *Daemon) StopHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {
	return nil, nil
}
