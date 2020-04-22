package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// IndexElem is used during used in CREATE INDEX, and in ON CONFLICT and
// represents an indexable column.
type IndexElem struct {
	Name interface{} `json:"name"` /* name of attribute to index, or NULL */
	Expr Node        `json:"expr"` /* expression to index, or NULL */
	// Indexcolname  *string     `json:"indexcolname"`   /* name for index column; NULL = default */
	// Collation List `json:"collation"` /* name of collation; NIL = default */
	// Opclass       List        `json:"opclass"`        /* name of desired opclass; NIL = default */
	// Ordering      SortByDir   `json:"ordering"`       /* ASC/DESC/default */
	// NullsOrdering SortByNulls `json:"nulls_ordering"` /* FIRST/LAST/default */
}

func (r *IndexElem) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	switch {
	case r.Name != nil && r.Expr == nil:
		if name, ok := r.Name.(string); ok {
			if _, err := buffer.WriteString(stdfmt.Sprintf("(%s)", name)); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
		}
	case r.Name == nil && r.Expr != nil:
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v or %#v", r.Name, r.Expr))
	}
	return buffer.String(), nil
}
