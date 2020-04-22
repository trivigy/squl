package squl

import (
	"bytes"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
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
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}

	switch r.Type {
	case JoinTypeDefault:
		if _, err := buffer.WriteString(" JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeInner:
		if _, err := buffer.WriteString(" INNER JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeLeft:
		if _, err := buffer.WriteString(" LEFT JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeOuterLeft:
		if _, err := buffer.WriteString(" LEFT OUTER JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeRight:
		if _, err := buffer.WriteString(" RIGHT JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeOuterRight:
		if _, err := buffer.WriteString(" RIGHT OUTER JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeFull:
		if _, err := buffer.WriteString(" FULL JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeOuterFull:
		if _, err := buffer.WriteString(" FULL OUTER JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case JoinTypeCross:
		if _, err := buffer.WriteString(" CROSS JOIN "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", r.Type))
	}

	rhsDump, err := r.RHS.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(rhsDump); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}

	if _, err := buffer.WriteString(" ON "); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	qualsDump, err := r.Qualifiers.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(qualsDump); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.String(), nil
}
