package squl

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// ColumnRef describes a fully qualified column name, possibly with subscripts.
type ColumnRef struct {
	// Fields represents field names, `*`, or subscripts.
	Fields interface{} `json:"fields"`
}

func (r *ColumnRef) dump(counter *ordinalMarker) (string, error) {
	switch fields := r.Fields.(type) {
	case string:
		return fields, nil
	case []string:
		return strings.Join(fields, "."), nil
	case []interface{}:
		buffer := bytes.NewBuffer(nil)
		for i, field := range fields {
			switch field := field.(type) {
			case Var:
				mark, err := field.dump(counter)
				if err != nil {
					return "", err
				}
				switch field.Value.(type) {
				case string:
					if i != 0 {
						if _, err := buffer.WriteString("."); err != nil {
							return "", errors.WithStack(err)
						}
					}
					if _, err := buffer.WriteString(mark); err != nil {
						return "", errors.WithStack(err)
					}
				case int:
					if _, err := buffer.WriteString(fmt.Sprintf("[%s]", mark)); err != nil {
						return "", errors.WithStack(err)
					}
				default:
					return "", errors.Errorf("type error %q", field.Value)
				}
			case string:
				if i != 0 {
					if _, err := buffer.WriteString("."); err != nil {
						return "", errors.WithStack(err)
					}
				}
				if _, err := buffer.WriteString(field); err != nil {
					return "", errors.WithStack(err)
				}
			case int:
				if i == 0 {
					return "", errors.Errorf("syntax error %q", r.Fields)
				}
				if _, err := buffer.WriteString(fmt.Sprintf("[%d]", field)); err != nil {
					return "", errors.WithStack(err)
				}
			default:

			}
		}
		return buffer.String(), nil
	default:
		return "", errors.Errorf("unknown type (%T) for %q", r.Fields, r.Fields)
	}
}
