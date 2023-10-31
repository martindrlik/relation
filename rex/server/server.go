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

func readState[R any](f func(*serverState) (R, error)) (R, error) {
	state.mu.RLock()
	defer state.mu.RUnlock()
	return f(state)
}

func writeState[R any](f func(*serverState) (R, error)) (R, error) {
	state.mu.Lock()
	defer state.mu.Unlock()
	return f(state)
}

func (s *serverState) r(name string) *rex.Relation {
	r, ok := s.rs[name]
	if ok {
		return r
	}
	return rex.NewRelation()
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
