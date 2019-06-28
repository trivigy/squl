package squl

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

// ResTarget defines the struct for generating target clauses like column_name.
type ResTarget struct {
	// Name represents a column name or nil.
	Name Node `json:"name"`

	// Value represents the value expression to compute or assign
	Value Node `json:"value"`

	// Alias represents column name or NULL
	Alias string `json:"alias"`
}

func (r *ResTarget) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.Name != nil {
		dump, err := r.Name.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(fmt.Sprintf("%s = ", dump)); err != nil {
			return "", errors.WithStack(err)
		}
	}

	valDump, err := r.Value.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(valDump); err != nil {
		return "", errors.WithStack(err)
	}

	if r.Alias != "" {
		if _, err := buffer.WriteString(fmt.Sprintf(" AS %s", r.Alias)); err != nil {
			return "", errors.WithStack(err)
		}
	}
	return buffer.String(), nil
}
