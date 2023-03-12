package rex

type column struct {
	name string
	data []any
}

func (c *column) dataAt(i int) any { return c.data[i] }
func (c *column) dataLen() int     { return len(c.data) }
func (c *column) insertData(v any) { c.data = append(c.data, v) }

func (c *column) removeDataAt(i int) {
	last := c.dataLen() - 1
	if i < last {
		c.data[i] = c.data[last]
	}
	c.data = c.data[:last]
}

type columns []column

func (cs columns) Len() int           { return len(cs) }
func (cs columns) Less(i, j int) bool { return cs[i].name < cs[j].name }
func (cs columns) Swap(i, j int) {
	x := cs[i]
	cs[i] = cs[j]
	cs[j] = x
}
