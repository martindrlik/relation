package server

import "github.com/martindrlik/rex"

type NaturalJoinArgs struct {
	Left, Right, Target string
}

func (*Rex) NaturalJoin(args *NaturalJoinArgs, reply *struct{}) error {
	r, _ := readState(func(s *serverState) (*rex.Relation, error) {
		l := s.r(args.Left)
		r := s.r(args.Right)
		return rex.NaturalJoin(l, r), nil
	})
	_, err := writeState(func(s *serverState) (struct{}, error) {
		s.rs[args.Target] = r
		return struct{}{}, nil
	})
	return err
}
