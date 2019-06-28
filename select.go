package squl

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

// Select defines the struct for the SELECT/SELECT INTO commands.
type Select struct {
	// With represents WITH clause.
	// With *With `json:"with"`

	// Distinct represents NULL, list of DISTINCT ON exprs, or * lcons(NIL,NIL) for all (SELECT DISTINCT)
	Distinct Node `json:"distinct"`

	// Into represents target for SELECT INTO command.
	Into *Into `json:"into"`

	// Targets represents the target list to select.
	Targets Node `json:"targets"`

	// From represents the FROM clause.
	From Node `json:"from"`

	// Where represents WHERE qualification.
	Where Node `json:"where"`

	// Group represents GROUP BY clauses.
	Group Node `json:"group"`

	// Having represents HAVING conditional-expression.
	Having Node `json:"having"`

	// Window represents WINDOW window_name AS (...), ...
	Window Node `json:"window"`

	// Values represents a VALUES list. If a "leaf" node representing a VALUES
	// list, the above fields are all null, and instead this field is set. Note
	// that the elements of the sublists are just expressions, without ResTarget
	// decoration. Also note that a list element can be DEFAULT (represented as
	// a SetToDefault node), regardless of the context of the VALUES list. It's
	// up to parse analysis to reject that where not valid.
	Values []Node `json:"values"`

	// Sort represents a list of ORDER BY.
	OrderBy Node `json:"orderBy"`

	// Offset represents the number of rows to skip before returning rows.
	Offset Node `json:"offset"`

	// Limit represents the maximum number of rows to returnl
	Limit Node `json:"limit"`

	// LockingClause List        `json:"lockingClause"` /* FOR UPDATE (list of LockingClause's) */

	/*
	 * These fields are used only in upper-level SelectStmts.
	 */
	// Name   SetOperation `json:"op"`   /* type of set op */
	// All  bool         `json:"all"`  /* ALL specified? */
	// Larg *SelectStmt  `json:"larg"` /* left child */
	// Rarg *SelectStmt  `json:"rarg"` /* right child */
}

func (r *Select) build() (string, []interface{}, error) {
	counter := &ordinalMarker{}
	query, err := r.dump(counter)
	return query, counter.args(), err
}

func (r *Select) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer([]byte("SELECT"))
	if r.Distinct != nil {
		if _, err := buffer.WriteString(" DISTINCT"); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Distinct.dump(counter)
		if err != nil {
			return "", err
		}

		if len(dump) > 0 {
			if _, err := buffer.WriteString(fmt.Sprintf(" ON (%s)", dump)); err != nil {
				return "", errors.WithStack(err)
			}
		}
	}

	if r.Targets != nil {
		if _, err := buffer.WriteString(" "); err != nil {
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

	if r.OrderBy != nil {
		if _, err := buffer.WriteString(" ORDER BY "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.OrderBy.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.Offset != nil {
		if _, err := buffer.WriteString(" OFFSET "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Offset.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if r.Limit != nil {
		if _, err := buffer.WriteString(" LIMIT "); err != nil {
			return "", errors.WithStack(err)
		}
		dump, err := r.Limit.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", errors.WithStack(err)
		}
	}

	return buffer.String(), nil
}

func (r *Select) isSelectStmt() bool {
	return r.Distinct != nil ||
		r.Into != nil ||
		r.Targets != nil ||
		r.Where != nil ||
		r.Group != nil ||
		r.Having != nil ||
		r.Window != nil
}

func (r *Select) isValuesList() bool {
	return r.Values != nil
}
