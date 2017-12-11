// http pressure test server
package httpsrv

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	static2k  io.ReadCloser
	static64k io.ReadCloser
)

func init() {
	buf2k := &bytes.Buffer{}
	for i := 0; i < 128*2; i++ {
		buf2k.WriteString("aaaaaaaa")
	}
	static2k = ioutil.NopCloser(buf2k)

	buf64k := &bytes.Buffer{}
	for i := 0; i < 128*64; i++ {
		buf64k.WriteString("aaaaaaaa")
	}
	static64k = ioutil.NopCloser(buf64k)

	prometheus.MustRegister()
}

func Run(addr string, chStop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/empty", handleEmpty)
	mux.HandleFunc("/static2k", handleStatic2k)
	mux.HandleFunc("/static64k", handleStatic64k)
	mux.HandleFunc("/static2M", handleStatic2M)

	mux.Handle("/metrics", promhttp.Handler())

	lis, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}

	go func(l net.Listener, ch <-chan struct{}) {
		select {
		case <-ch:
			l.Close()
		}
	}(lis, chStop)

	return http.Serve(lis, mux)
}

func handleEmpty(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(204)
}

func handleStatic2k(rw http.ResponseWriter, req *http.Request) {
	io.Copy(rw, static2k)
}

func handleStatic64k(rw http.ResponseWriter, req *http.Request) {
	io.Copy(rw, static64k)
}

func handleStatic2M(rw http.ResponseWriter, req *http.Request) {
	for i := 0; i < 32; i++ {
		io.Copy(rw, static64k)
	}
}
