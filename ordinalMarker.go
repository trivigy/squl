package squl

import (
	"fmt"
)

type ordinalMarker struct {
	entries []interface{}
}

func (r *ordinalMarker) mark(value interface{}) string {
	r.entries = append(r.entries, value)
	return fmt.Sprintf("$%d", len(r.entries))
}

func (r *ordinalMarker) args() []interface{} {
	if r.entries == nil {
		r.entries = []interface{}{}
	}
	return r.entries
}
