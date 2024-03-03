package table

import (
	"fmt"
)

func ErrDomainTable(attribute string) error {
	return fmt.Errorf("domain table need relation for attribute %s", attribute)
}
