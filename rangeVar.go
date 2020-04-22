package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// RangeVar describes a schema qualified target table.
type RangeVar struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
	// Inheritance bool        `json:"inheritance"`
	// Persistence Persistence `json:"persistence"`
	Alias string `json:"alias"`
}

func (r *RangeVar) dump(_ *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.Schema != "" {
		if _, err := buffer.WriteString(stdfmt.Sprintf("%s.", r.Schema)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	if r.Name == "" {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Name))
	}
	if _, err := buffer.WriteString(r.Name); err != nil {
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}

	if r.Alias != "" {
		if _, err := buffer.WriteString(stdfmt.Sprintf(" AS %s", r.Alias)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}
