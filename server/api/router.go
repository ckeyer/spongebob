package api

import (
	"net/http"

	pb "github.com/ckeyer/spongebob/protos"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func newGateway(endpoint string) (http.Handler, error) {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterControllerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
