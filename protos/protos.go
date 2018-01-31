//go:generate sh -c "set -ex; protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. --go_out=plugins=grpc:. $GOPATH/src/github.com/ckeyer/spongebob/protos/*.proto"
// protoc --go_out=. *.proto
package protos
