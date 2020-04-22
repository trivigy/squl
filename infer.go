package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// Infer represents the ON CONFLICT unique index inference clause.
type Infer struct {
	IndexElems  *List   `json:"indexElems"`  /* IndexElems to infer unique index */
	WhereClause Node    `json:"whereClause"` /* qualification (partial-index predicate) */
	Conname     *string `json:"conname"`     /* Constraint name, or NULL if unnamed */
	Location    int     `json:"location"`    /* token location, or -1 if unknown */
}

func (r *Infer) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.IndexElems != nil {
		dump, err := r.IndexElems.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(stdfmt.Sprintf("(%s)", dump)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}
