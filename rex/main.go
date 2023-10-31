package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/martindrlik/rex/rex/client"
	"github.com/martindrlik/rex/rex/server"
)

var (
	addr string

	binary struct {
		left, right string
	}

	target  string
	jsonStr string
)

func main() {
	must(command(os.Args[1])(os.Args[2:]))
}

func command(name string) func([]string) error {
	switch name {
	case "server":
		return listen
	case "insert":
		return insert
	case "list":
		return list
	case "natural-join", "nj":
		return naturalJoin
	}
	return func([]string) error {
		return fmt.Errorf("unknown command: %q", name)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func listen(args []string) error {
	flagParse("server", args, addrfs)
	return server.Listen(addr)
}

func insert(args []string) error {
	flagParse("insert", args, addrfs, jsonStrfs, targetfs)
	m := map[string]any{}
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		return fmt.Errorf("unmarshaling: %w: %q\n", err, jsonStr)
	}
	if err := client.Insert(addr, target, m); err != nil {
		return fmt.Errorf("calling: %w\n", err)
	}
	return nil
}

func list(args []string) error {
	flagParse("list", args, addrfs, targetfs)
	ts, err := client.List(addr, target)
	if err != nil {
		return fmt.Errorf("calling: %w\n", err)
	}
	for _, t := range ts {
		certainty := func() string {
			if t.IsPossible {
				return "*"
			}
			return ""
		}
		fmt.Printf("%v%v\n", t.Map, certainty())
	}
	return nil
}

func naturalJoin(args []string) error {
	flagParse("natural-join", args, addrfs, targetfs, binaryfs)
	return client.NaturalJoin(addr, binary.left, binary.right, target)
}

func flagParse(name string, args []string, fns ...func(*flag.FlagSet)) error {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	for _, f := range fns {
		f(fs)
	}
	return fs.Parse(args)
}

func addrfs(fs *flag.FlagSet) {
	fs.StringVar(&addr, "addr", ":1234", "address of the server")
}

func jsonStrfs(fs *flag.FlagSet) {
	fs.StringVar(&jsonStr, "json", "", "relation as a json string")
}

func targetfs(fs *flag.FlagSet) {
	fs.StringVar(&target, "target", "", "name of the relation")
}

func binaryfs(fs *flag.FlagSet) {
	fs.StringVar(&binary.left, "left", "", "name of the left relation")
	fs.StringVar(&binary.right, "right", "", "name of the right relation")
}
