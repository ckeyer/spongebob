package server

import (
	log "github.com/ckeyer/logrus"
	"github.com/ckeyer/spongebob/server/api"
	"github.com/ckeyer/spongebob/server/rpc"
)

func Run(rpcAddr, webAddr string) error {
	go startHTTP(rpcAddr, webAddr)
	return startRPC(rpcAddr, webAddr)
}

func startHTTP(rpcAddr, webAddr string) {
	err := api.Serve(rpcAddr, webAddr)
	if err != nil {
		log.Fatalf("http listenning failed, %s", err)
	}
}

func startRPC(rpcAddr, webAddr string) error {
	err := rpc.Serve(rpcAddr)
	if err != nil {
		log.Fatalf("http listenning failed, %s", err)
		return err
	}
	return nil
}
