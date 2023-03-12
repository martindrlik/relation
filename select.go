package rex

import "encoding/json"

type Select struct {
	whereMap map[string]interface{}
}

func (sel *Select) setWhere(s string) {
	w := map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &w)
	if err != nil {
		panic(err)
	}
	sel.whereMap = w
}

func Where(s string) func(*Select) {
	return func(sel *Select) {
		sel.setWhere(s)
	}
}
