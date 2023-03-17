package rex

type DumpOptions struct {
	padding map[string]int
}

func buildDumpOptions(options ...func(*DumpOptions)) *DumpOptions {
	o := &DumpOptions{
		padding: map[string]int{},
	}
	for _, option := range options {
		option(o)
	}
	return o
}

func Pad(name string, pad int) func(*DumpOptions) {
	return func(do *DumpOptions) {
		do.padding[name] = pad
	}
}
