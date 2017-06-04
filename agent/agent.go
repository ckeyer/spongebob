package agent

import (
	log "github.com/Sirupsen/logrus"
	pb "github.com/ckeyer/spongebob/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Agent struct {
	Name string
}

type TaskManager struct {
	cli pb.ControllerClient
}

func Start(endpoint string) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Infof("connected to %s", endpoint)

	cli := pb.NewControllerClient(conn)

	tm := newTaskManager(cli)
	tm.join()

	return nil
}

func newTaskManager(cli pb.ControllerClient) *TaskManager {
	return &TaskManager{
		cli: cli,
	}
}

func (tm *TaskManager) join() error {
	node := &pb.Node{
		Info: &pb.NodeInfo{
			Hostname: "test",
		},
	}

	log.Debugf("try to join cluster.")
	t, err := tm.cli.Join(context.Background(), node)
	if err != nil {
		return err
	}
	log.Debugf("get task %+v", t)

	return nil
}
