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
		Tuples []struct {
			Map        map[string]any
			IsPossible bool
		}
	}

	Rex struct{}
)

func (*Rex) Insert(args *InsertArgs, reply *struct{}) error {
	_, err := writeState(func(s *serverState) (struct{}, error) {
		r, ok := s.rs[args.Name]
		if !ok {
			r = rex.NewRelation()
			s.rs[args.Name] = r
		}
		r.Insert(args.Tuple)
		return struct{}{}, nil
	})
	return err
}

func (*Rex) List(args *ListArgs, reply *ListReply) error {
	_, err := readState(func(s *serverState) (struct{}, error) {
		r, ok := s.rs[args.Name]
		if !ok {
			return struct{}{}, nil
		}
		reply.Tuples = []struct {
			Map        map[string]any
			IsPossible bool
		}{}
		r.Each(func(m map[string]any, b bool) bool {
			reply.Tuples = append(reply.Tuples, struct {
				Map        map[string]any
				IsPossible bool
			}{m, b})
			return true
		})
		return struct{}{}, nil
	})
	return err
}
