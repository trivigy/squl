package squl

import (
	stdfmt "fmt"
	"strings"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// Expr describes the expression clause.
type Expr struct {
	Type ExprType `json:"type"` /* see above */

	// Wrap indicates whether or not to wrap the expression in parentheses.
	Wrap bool

	// Name represents possibly-qualified name of operator.
	Name interface{} `json:"name"` /* possibly-qualified name of operator */

	LHS Node `json:"lhs"` /* left argument, or NULL if none */
	RHS Node `json:"rhs"` /* right argument, or NULL if none */
}

func (r *Expr) dump(counter *ordinalMarker) (string, error) {
	switch r.Type {
	case ExprTypeOp:
		lhsDump, err := r.dumpOperand(counter, r.LHS)
		if err != nil {
			return "", err
		}

		opDump, err := r.dumpOperator(counter, r.Name)
		if err != nil {
			return "", err
		}

		rhsDump, err := r.dumpOperand(counter, r.RHS)
		if err != nil {
			return "", err
		}

		expr := stdfmt.Sprintf("%s %s %s", lhsDump, opDump, rhsDump)
		if r.Wrap {
			expr = "(" + expr + ")"
		}

		return expr, nil
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", r.Type))
	}
}

func (r *Expr) dumpOperand(counter *ordinalMarker, op interface{}) (string, error) {
	switch op := op.(type) {
	case Node:
		return op.dump(counter)
	case string:
		return stdfmt.Sprintf("%q", op), nil
	case int:
		return stdfmt.Sprintf("%d", op), nil
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type (%T) for %q", op, op))
	}
}

func (r *Expr) dumpOperator(_ *ordinalMarker, _ interface{}) (string, error) {
	switch name := r.Name.(type) {
	case string:
		return name, nil
	case []string:
		return stdfmt.Sprintf("OPERATOR(%s)", strings.Join(name, ".")), nil
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type (%T) for %q", name, name))
	}
}
