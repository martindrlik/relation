package server

import (
	"github.com/martindrlik/rex"
)

type (
	InsertArgs struct {
		Name  string
		Tuple map[string]any
	}

	ListArgs struct {
		Name string
	}

	ListReply struct {
		Tuples []map[string]any
	}

	Rex struct{}
)

func (*Rex) Insert(args *InsertArgs, reply *struct{}) error {
	state.Write(func(s *serverState) {
		r, ok := s.rs[args.Name]
		if !ok {
			r = rex.NewRelation()
			s.rs[args.Name] = r
		}
		r.Insert(args.Tuple)
	})
	return nil
}

func (*Rex) List(args *ListArgs, reply *ListReply) error {
	state.Read(func(s *serverState) {
		r, ok := s.rs[args.Name]
		if !ok {
			return
		}
		reply.Tuples = []map[string]any{}
		r.Each(func(m map[string]any, b bool) bool {
			reply.Tuples = append(reply.Tuples, m)
			return true
		})
	})
	return nil
}
