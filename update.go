package squl

import (
	"bytes"

	"github.com/pkg/errors"
)

// Update defines the struct for the UPDATE command.
type Update struct {
	// With represents WITH clause.
	// With *With `json:"with"`

	Relation *RangeVar `json:"relation"` /* relation to update */

	// Targets represents the target list to select.
	Targets Node `json:"targets"`

	// From represents the FROM clause.
	From Node `json:"from"`

	// Where represents WHERE qualification.
	Where Node `json:"where"`

	// Returning represents a list of expressions to return.
	Returning Node `json:"returning"`
}

func (r *Update) build() (string, []interface{}, error) {
	counter := &ordinalMarker{}
	query, err := r.dump(counter)
	return query, counter.args(), err
}

func (r *Update) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer([]byte("UPDATE"))
	if r.Relation != nil {
		if _, err := buffer.WriteString(" "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Relation.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.Targets != nil {
		if _, err := buffer.WriteString(" SET "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Targets.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.From != nil {
		if _, err := buffer.WriteString(" FROM "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.From.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.Where != nil {
		if _, err := buffer.WriteString(" WHERE "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Where.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.Returning != nil {
		if _, err := buffer.WriteString(" RETURNING "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Returning.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}
	return buffer.String(), nil
}
