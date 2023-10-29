package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martindrlik/flag"
	"github.com/martindrlik/rex/rex/client"
	"github.com/martindrlik/rex/rex/server"
)

var (
	addr     = flag.String("addr", ":1234", "server address")
	isServer = flag.Bool("server", false, "run as server")
	name     = flag.String("name", "test", "name of relation")
	jsonStr  = flag.String("json", "", "json representation of relation to insert")
)

var command = map[string]func() error{
	"insert": func() error {
		m := map[string]any{}
		if err := json.Unmarshal([]byte(*jsonStr), &m); err != nil {
			return fmt.Errorf("unmarshaling: %w: %q\n", err, *jsonStr)
		}
		if err := client.Insert(*addr, *name, m); err != nil {
			return fmt.Errorf("calling: %w\n", err)
		}
		return nil
	},
	"list": func() error {
		ms, err := client.List(*addr, *name)
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
	flag.Parse()
	if *isServer {
		panic(server.Listen(*addr))
	}
	cmdName := flag.Arg(0)
	cmd, ok := command[cmdName]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmdName)
		return
	}
	if err := cmd(); err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %v\n", cmdName, err)
	}
}
