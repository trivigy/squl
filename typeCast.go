package squl

import (
	"fmt"

	"github.com/pkg/errors"
)

// TypeCast describes the type casting expression.
type TypeCast struct {
	Arg  Node   `json:"arg"`  /* the expression being casted */
	Type string `json:"type"` /* the target type */
}

func (r *TypeCast) dump(counter *ordinalMarker) (string, error) {
	if r.Arg == nil {
		return "", errors.Errorf("required parameter %#v", r.Arg)
	}

	if r.Type == "" {
		return "", errors.Errorf("required parameter %#v", r.Type)
	}

	dump, err := r.Arg.dump(counter)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s::%s", dump, r.Type), nil
}
