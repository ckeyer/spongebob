package api

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func newGateway(endpoint string) (http.Handler, error) {
	mux := runtime.NewServeMux()

	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := pb.RegisterControllerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	// if err != nil {
	// 	return nil, err
	// }

	return mux, nil
}
