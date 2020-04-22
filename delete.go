package squl

import (
	"bytes"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// Delete defines the struct for the DELETE command.
type Delete struct {
	// With represents WITH clause.
	// With *With `json:"with"`

	// Relation represents the relation to delete from.
	Relation *RangeVar `json:"relation"`

	// Using represents optional USING clause for more tables.
	Using Node `json:"using"`

	// Where represents WHERE qualification.
	Where Node `json:"where"`

	// Returning represents a list of expressions to return.
	Returning Node `json:"returning"`
}

func (r *Delete) build() (string, []interface{}, error) {
	counter := &ordinalMarker{}
	query, err := r.dump(counter)
	return query, counter.args(), err
}

func (r *Delete) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer([]byte("DELETE FROM"))
	if r.Relation != nil {
		if _, err := buffer.WriteString(" "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Relation.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Using != nil {
		if _, err := buffer.WriteString(" USING "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Using.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Where != nil {
		if _, err := buffer.WriteString(" WHERE "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Where.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Returning != nil {
		if _, err := buffer.WriteString(" RETURNING "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Returning.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}
