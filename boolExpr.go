package squl

import (
	"bytes"

	"github.com/pkg/errors"
)

// BoolExpr describes the AND/OR expression combinations.
type BoolExpr struct {
	Xpr  Node         `json:"xpr"`
	Type BoolExprType `json:"type"`
	Wrap bool         `json:"wrap"` /* indicate if expr should be wrapped with parentheses */
	Args []Node       `json:"args"` /* arguments to this expression */
}

func (r *BoolExpr) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	for i, each := range r.Args {
		eachDump, err := each.dump(counter)
		if err != nil {
			return "", err
		}

		if _, err := buffer.WriteString(eachDump); err != nil {
			return "", errors.WithStack(err)
		}

		if i < len(r.Args)-1 {
			switch r.Type {
			case BoolExprTypeAnd:
				if _, err := buffer.WriteString(" AND "); err != nil {
					return "", errors.WithStack(err)
				}
			case BoolExprTypeOr:
				if _, err := buffer.WriteString(" OR "); err != nil {
					return "", errors.WithStack(err)
				}
			default:
				return "", errors.Errorf("unknown type %q", r.Type)
			}
		}
	}

	expr := buffer.String()
	if r.Wrap {
		expr = "(" + expr + ")"
	}
	return expr, nil
}
