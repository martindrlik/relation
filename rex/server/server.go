package server

import (
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/martindrlik/rex"
)

type (
	serverState struct {
		rs map[string]*rex.Relation
		mu sync.RWMutex
	}
)

var state = &serverState{
	rs: make(map[string]*rex.Relation),
}

func (s *serverState) Read(read func(*serverState)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	read(s)
}

func (s *serverState) Write(write func(*serverState)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	write(s)
}

func Listen(addr string) error {
	rpc.Register(new(Rex))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", addr)
	if err == nil {
		err = http.Serve(l, nil)
	}
	return err
}
