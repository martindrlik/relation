package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/martindrlik/rex"
)

var (
	union       = flag.Bool("union", false, "")
	naturaljoin = flag.Bool("naturaljoin", false, "")
)

func main() {
	flag.Parse()
	s := []*rex.Relation{}
	for _, name := range flag.Args() {
		s = append(s, tryReadRelation(name))
	}
	if len(s) == 0 {
		os.Exit(1)
	}
	r := s[0]
	for i := 1; i < len(s); i++ {
		if *union {
			r = r.Union(s[i])
		}
		if *naturaljoin {
			r = r.NaturalJoin(s[i])
		}
	}
	r.Serialize(os.Stdout)
	fmt.Println()
}

func tryReadRelation(name string) *rex.Relation {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return rex.NewRelation().InsertManyJson(f)
}
