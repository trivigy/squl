package squl

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

// RangeVar describes a schema qualified target table.
type RangeVar struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
	// Inheritance bool        `json:"inheritance"`
	// Persistence Persistence `json:"persistence"`
	Alias string `json:"alias"`
}

func (r *RangeVar) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.Schema != "" {
		if _, err := buffer.WriteString(fmt.Sprintf("%s.", r.Schema)); err != nil {
			return "", errors.WithStack(err)
		}
	}
	if r.Name == "" {
		return "", errors.Errorf("required parameter %#v", r.Name)
	}
	if _, err := buffer.WriteString(r.Name); err != nil {
		return "", errors.WithStack(err)
	}

	if r.Alias != "" {
		if _, err := buffer.WriteString(fmt.Sprintf(" AS %s", r.Alias)); err != nil {
			return "", errors.WithStack(err)
		}
	}
	return buffer.String(), nil
}
