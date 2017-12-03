package rpc

import (
	"log"
	"net"

	pb "github.com/ckeyer/spongebob/protos"
	"google.golang.org/grpc"
)

func Serve(rpcAddr string) error {
	lis, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterControllerServer(s, NewControllerServer())

	return s.Serve(lis)
}
