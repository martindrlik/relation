package client

import (
	"fmt"
	"net/rpc"

	"github.com/martindrlik/rex/rex/server"
)

func Insert(addr, name string, m map[string]any) error {
	client, err := dial(addr)
	if err != nil {
		return fmt.Errorf("dialing: %w", err)
	}
	args := server.InsertArgs{Name: name, Tuple: m}
	reply := struct{}{}
	err = client.Call("Rex.Insert", &args, &reply)
	if err != nil {
		return fmt.Errorf("calling: %w", err)
	}
	return nil
}

func List(addr, name string) ([]map[string]any, error) {
	client, err := dial(addr)
	if err != nil {
		return nil, fmt.Errorf("dialing: %w", err)
	}
	args := server.ListArgs{Name: name}
	reply := server.ListReply{}
	err = client.Call("Rex.List", &args, &reply)
	if err != nil {
		return nil, fmt.Errorf("calling: %w", err)
	}
	return reply.Tuples, nil
}

func dial(addr string) (*rpc.Client, error) {
	return rpc.DialHTTP("tcp", addr)
}
