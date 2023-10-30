package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/martindrlik/rex/rex/client"
	"github.com/martindrlik/rex/rex/server"
)

var (
	serverFlagSet = flag.NewFlagSet("server", flag.ExitOnError)
	clientFlagSet = flag.NewFlagSet("client", flag.ExitOnError)

	addr string
	name string
	jstr string
)

func init() {
	for _, fs := range []*flag.FlagSet{serverFlagSet, clientFlagSet} {
		fs.StringVar(&addr, "addr", ":1234", "server address")
		fs.StringVar(&addr, "a", ":1234", "server address (shorthand)")
	}
	for _, fs := range []*flag.FlagSet{clientFlagSet} {
		fs.StringVar(&jstr, "json", "", "json representation of relation to insert")
		fs.StringVar(&jstr, "j", "", "json representation of relation to insert (shorthand)")
		fs.StringVar(&name, "name", "default", "name of relation")
	}
}

var command = map[string]func() error{
	"insert": func() error {
		m := map[string]any{}
		if err := json.Unmarshal([]byte(jstr), &m); err != nil {
			return fmt.Errorf("unmarshaling: %w: %q\n", err, jstr)
		}
		if err := client.Insert(addr, name, m); err != nil {
			return fmt.Errorf("calling: %w\n", err)
		}
		return nil
	},
	"list": func() error {
		ms, err := client.List(addr, name)
		if err != nil {
			return fmt.Errorf("calling: %w\n", err)
		}
		for _, m := range ms {
			fmt.Println(m)
		}
		return nil
	},
}

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "server":
		serverFlagSet.Parse(os.Args[2:])
		panic(server.Listen(addr))
	case "insert", "list":
		clientFlagSet.Parse(os.Args[2:])
	}
	h, ok := command[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		return
	}
	if err := h(); err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %v\n", cmd, err)
	}
}
