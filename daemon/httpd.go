package daemon

import (
	"github.com/ckeyer/spongebob/pkgs/httpsrv"
	pb "github.com/ckeyer/spongebob/protos"
	"golang.org/x/net/context"
)

func (d *Daemon) StartHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {
	d.chStopHTTP
	select {
	case <-chStopHTTP:
	default:

	}
	httpsrv.Run()
	return nil, nil
}

func (d *Daemon) StopHTTP(ctx context.Context, in *pb.HTTPOption) (*pb.HTTPOption, error) {
	return nil, nil
}
