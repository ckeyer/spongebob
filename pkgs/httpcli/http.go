package httpcli

import (
	"net/url"
	"runtime"
	"strings"

	"github.com/ckeyer/attack/httpclient"

	"github.com/ckeyer/logrus"
)

const (
	defaultProcNum = runtime.NumCPU() * 2
)

var log = logrus.New(logrus.Fields{"module": "http"})

type Job struct {
	Method   string
	URL      *url.URL
	MaxClis  int
	MaxProcs int
	Headers  *url.Values
	Body     interface{}

	chStop chan struct{}
}

func New(method, Url string) *Job {
	u, err := url.Parse(Url)
	if err != nil {
		log.Fatalln(err)
	}
	switch strings.ToUpper(method) {
	case "GET", "POST", "PATCH", "PUT":
	default:
		log.Fatalln("not support http method %s", method)
	}
	return &Job{
		URL:    u,
		Method: strings.ToUpper(method),
		chStop: make(chan struct{}),
	}
}

func (j *Job) Start() {

}

func (j *Job) Stop() {
	select {
	case <-j.chStop:
	default:
		close(j.chStop)
		log.Infof("stop job")
	}
}

func (j *Job) httpCli() *httpclient.Client {
	httpclient.NewClient()
}
