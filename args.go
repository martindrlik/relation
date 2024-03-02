package main

type args []string

func (a args) Head() string {
	return a[0]
}

func (a args) Tail() args {
	return a[1:]
}
