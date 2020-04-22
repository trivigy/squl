package squl

import (
	"bytes"
	"encoding/json"
	"strings"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// OrderByNulls describes NULLS sorting order.
type OrderByNulls int

const (
	// OrderByNullsFirst indicates that sorting should list nulls first.
	OrderByNullsFirst OrderByNulls = iota + 1

	// OrderByNullsLast indicates that sorting should list nulls last.
	OrderByNullsLast
)

const (
	orderByNullsFirstStr = "first"
	orderByNullsLastStr  = "last"
)

var toStringOrderByNulls = map[OrderByNulls]string{
	OrderByNulls(Unknown): unknownStr,
	OrderByNullsFirst:     orderByNullsFirstStr,
	OrderByNullsLast:      orderByNullsLastStr,
}

// NewOrderByNulls creates a new instance of the enum from raw string.
func NewOrderByNulls(raw string) (OrderByNulls, error) {
	switch raw {
	case orderByNullsFirstStr:
		return OrderByNullsFirst, nil
	case orderByNullsLastStr:
		return OrderByNullsLast, nil
	default:
		return OrderByNulls(Unknown), fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
}

// String returns the string representation of the enum type
func (r OrderByNulls) String() string {
	return toStringOrderByNulls[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *OrderByNulls) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case orderByNullsFirstStr:
		*r = OrderByNullsFirst
	case orderByNullsLastStr:
		*r = OrderByNullsLast
	default:
		*r = Unknown
		return fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r OrderByNulls) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringOrderByNulls[r]); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.Bytes(), nil
}
