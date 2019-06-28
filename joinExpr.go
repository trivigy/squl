package squl

import (
	"bytes"

	"github.com/pkg/errors"
)

// JoinExpr defines the struct for generating JOIN clauses.
type JoinExpr struct {
	Type    JoinType `json:"type"`    /* type of join */
	Natural bool     `json:"natural"` /* Natural join? Will need to shape table */
	LHS     Node     `json:"lhs"`     /* left subtree */
	RHS     Node     `json:"rhs"`     /* right subtree */
	// Using     List     `json:"using"`     /* USING clause, if any (list of String) */
	Qualifiers Node `json:"qualifiers"` /* qualifiers on join, if any */
	// Alias     *Alias   `json:"alias"`     /* user-written alias clause, if any */
	// Rtindex   int      `json:"rtindex"`   /* RT index assigned for join, or 0 */
}

func (r *JoinExpr) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	lhsDump, err := r.LHS.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(lhsDump); err != nil {
		return "", errors.WithStack(err)
	}

	switch r.Type {
	case JoinTypeDefault:
		if _, err := buffer.WriteString(" JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeInner:
		if _, err := buffer.WriteString(" INNER JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeLeft:
		if _, err := buffer.WriteString(" LEFT JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeOuterLeft:
		if _, err := buffer.WriteString(" LEFT OUTER JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeRight:
		if _, err := buffer.WriteString(" RIGHT JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeOuterRight:
		if _, err := buffer.WriteString(" RIGHT OUTER JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeFull:
		if _, err := buffer.WriteString(" FULL JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeOuterFull:
		if _, err := buffer.WriteString(" FULL OUTER JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	case JoinTypeCross:
		if _, err := buffer.WriteString(" CROSS JOIN "); err != nil {
			return "", errors.WithStack(err)
		}
	default:
		return "", errors.Errorf("unknown type %q", r.Type)
	}

	rhsDump, err := r.RHS.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(rhsDump); err != nil {
		return "", errors.WithStack(err)
	}

	if _, err := buffer.WriteString(" ON "); err != nil {
		return "", errors.WithStack(err)
	}
	qualsDump, err := r.Qualifiers.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(qualsDump); err != nil {
		return "", errors.WithStack(err)
	}
	return buffer.String(), nil
}
