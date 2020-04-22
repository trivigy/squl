package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// ArrayExpr - an ARRAY[] construct
type ArrayExpr List

var _ Node = &ArrayExpr{}

func (r *ArrayExpr) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r != nil && len(*r) > 0 {
		super := List(*r)
		dump, err := super.dump(counter)
		if err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}

		if _, err := buffer.WriteString(stdfmt.Sprintf("ARRAY[%s]", dump)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}
