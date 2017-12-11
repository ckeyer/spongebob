package daemon

import (
	"github.com/ckeyer/spongebob/pkgs/httpsrv"
	pb "github.com/ckeyer/spongebob/protos"
	"golang.org/x/net/context"
)

const (
	statusRun = "running"
)

func (d *Daemon) StartHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {
	if d.httpStatus != statusRun {
		httpsrv.Run("", d.chStopHTTP)
	}
	return nil, nil
}

func (d *Daemon) StopHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {
	return nil, nil
}
