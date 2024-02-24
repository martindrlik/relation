package main

type args []string

func (args args) first() string {
	return args[0]
}

func (args args) rest() args {
	return args[1:]
}
