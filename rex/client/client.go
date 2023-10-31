package client

import (
	"fmt"
	"net/rpc"

	"github.com/martindrlik/rex/rex/server"
)

func NaturalJoin(addr, left, right, target string) error {
	client, err := dial(addr)
	if err != nil {
		return err
	}
	args := server.NaturalJoinArgs{
		Left:   left,
		Right:  right,
		Target: target,
	}
	reply := struct{}{}
	return call(client, "Rex.NaturalJoin", &args, &reply)
}

func Insert(addr, name string, m map[string]any) error {
	client, err := dial(addr)
	if err != nil {
		return err
	}
	args := server.InsertArgs{Name: name, Tuple: m}
	reply := struct{}{}
	return call(client, "Rex.Insert", &args, &reply)
}

func List(addr, name string) ([]struct {
	Map        map[string]any
	IsPossible bool
}, error) {
	client, err := dial(addr)
	if err != nil {
		return nil, err
	}
	args := server.ListArgs{Name: name}
	reply := server.ListReply{}
	if err := call(client, "Rex.List", &args, &reply); err != nil {
		return nil, err
	}
	return reply.Tuples, nil
}

func dial(addr string) (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dialing: %w", err)
	}
	return client, nil
}

type client rpc.Client

func call(client *rpc.Client, serviceMethod string, args, reply any) error {
	if err := client.Call(serviceMethod, args, reply); err != nil {
		return fmt.Errorf("calling: %w", err)
	}
	return nil
}
