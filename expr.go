package squl

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Expr describes the expression clause.
type Expr struct {
	Type ExprType `json:"type"` /* see above */

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

		return fmt.Sprintf("%s %s %s", lhsDump, opDump, rhsDump), nil
	default:
		return "", errors.Errorf("unknown type %q", r.Type)
	}
}

func (r *Expr) dumpOperand(counter *ordinalMarker, op interface{}) (string, error) {
	switch op := op.(type) {
	case *Expr:
		opDump, err := op.dump(counter)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(%s)", opDump), nil
	case Node:
		return op.dump(counter)
	case string:
		return fmt.Sprintf("%q", op), nil
	case int:
		return fmt.Sprintf("%d", op), nil
	default:
		return "", errors.Errorf("unknown type (%T) for %q", op, op)
	}
}

func (r *Expr) dumpOperator(counter *ordinalMarker, op interface{}) (string, error) {
	switch name := r.Name.(type) {
	case string:
		return name, nil
	case []string:
		return fmt.Sprintf("OPERATOR(%s)", strings.Join(name, ".")), nil
	default:
		return "", errors.Errorf("unknown type (%T) for %q", name, name)
	}
}
