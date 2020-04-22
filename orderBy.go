package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// OrderBy defines the ORDER BY clause.
type OrderBy struct {
	// Value represents expresiion to order by.
	Value Node `json:"value"`

	// Direction represents sort direction ASC/DESC/USING/default.
	Direction OrderByDirection `json:"direction"`

	// UsingOp represents name of op to use with OrderByDirectionUsing.
	UsingOp string `json:"UsingOp"`

	// Nulls represents NULLS sort order FIRST/LAST.
	Nulls OrderByNulls `json:"nulls"`
}

func (r *OrderBy) dump(counter *ordinalMarker) (string, error) {
	if r.Value == nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Value))
	}

	buffer := bytes.NewBuffer(nil)
	dump, err := r.Value.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(dump); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}

	switch r.Direction {
	case OrderByDirectionAsc:
		if _, err := buffer.WriteString(" ASC"); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case OrderByDirectionDesc:
		if _, err := buffer.WriteString(" DESC"); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case OrderByDirectionUsing:
		if _, err := buffer.WriteString(" USING"); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		if r.UsingOp == "" {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.UsingOp))
		}
		if _, err := buffer.WriteString(stdfmt.Sprintf(" %s", r.UsingOp)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	return buffer.String(), nil
}
