package squl

import (
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// TypeCast describes the type casting expression.
type TypeCast struct {
	Arg  Node   `json:"arg"`  /* the expression being casted */
	Type string `json:"type"` /* the target type */
}

func (r *TypeCast) dump(counter *ordinalMarker) (string, error) {
	if r.Arg == nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Arg))
	}

	if r.Type == "" {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Type))
	}

	dump, err := r.Arg.dump(counter)
	if err != nil {
		return "", err
	}
	return stdfmt.Sprintf("%s::%s", dump, r.Type), nil
}
