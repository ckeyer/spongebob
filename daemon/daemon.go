package daemon

import (
	"github.com/ckeyer/api"
	"golang.org/x/net/context"
)

type Daemon struct{}

func NewDaemon() *Daemon {
	return &Daemon{}
}

// GetStatus get daemon status.
func (d *Daemon) GetStatus(ctx context.Context, in *api.Empty) (*api.ProcessStatus, error) {
	return &api.ProcessStatus{}, nil
}

// Reload get daemon status.
func (d *Daemon) Reload(ctx context.Context, in *api.Empty) (*api.ProcessStatus, error) {
	return &api.ProcessStatus{}, nil
}
