package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/martindrlik/rex"
)

var (
	intersect     = flag.Bool("intersect", false, "")
	naturaljoin   = flag.Bool("naturaljoin", false, "")
	setdifference = flag.Bool("setdifference", false, "")
	union         = flag.Bool("union", false, "")
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
		if *intersect {
			r = r.Intersect(s[i])
		}
		if *naturaljoin {
			r = r.NaturalJoin(s[i])
		}
		if *setdifference {
			r = r.SetDifference(s[i])
		}
		if *union {
			r = r.Union(s[i])
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
