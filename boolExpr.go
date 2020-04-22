package squl

import (
	"bytes"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
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
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}

		if i < len(r.Args)-1 {
			switch r.Type {
			case BoolExprTypeAnd:
				if _, err := buffer.WriteString(" AND "); err != nil {
					return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
				}
			case BoolExprTypeOr:
				if _, err := buffer.WriteString(" OR "); err != nil {
					return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
				}
			default:
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", r.Type))
			}
		}
	}

	expr := buffer.String()
	if r.Wrap {
		expr = "(" + expr + ")"
	}
	return expr, nil
}
