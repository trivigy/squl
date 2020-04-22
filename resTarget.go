package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// ResTarget defines the struct for generating target clauses like column_name.
type ResTarget struct {
	// Name represents a column name or nil.
	Name interface{} `json:"name"`

	// Value represents the value expression to compute or assign
	Value Node `json:"value"`

	// Alias represents column name or NULL
	Alias string `json:"alias"`
}

func (r *ResTarget) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.Name != nil {
		switch name := r.Name.(type) {
		case string:
			if _, err := buffer.WriteString(stdfmt.Sprintf("%s = ", name)); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
		default:
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("type error %q", r.Name))
		}
	}

	valDump, err := r.Value.dump(counter)
	if err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(valDump); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}

	if r.Alias != "" {
		if _, err := buffer.WriteString(stdfmt.Sprintf(" AS %s", r.Alias)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}
