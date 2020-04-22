package squl

import (
	"bytes"
	"encoding/json"
	"strings"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// OrderByDirection describes the ORDER BY direction ASC/DESC/USING.
type OrderByDirection int

const (
	// OrderByDirectionAsc indicates the sorting direction is ascending.
	OrderByDirectionAsc OrderByDirection = iota + 1

	// OrderByDirectionDesc indicates the sorting direction is descending.
	OrderByDirectionDesc

	// OrderByDirectionUsing indicates the usage of custom sorting direction operator.
	OrderByDirectionUsing
)

const (
	orderByDirectionAscStr   = "asc"
	orderByDirectionDescStr  = "desc"
	orderByDirectionUsingStr = "using"
)

var toStringOrderByDirection = map[OrderByDirection]string{
	OrderByDirection(Unknown): unknownStr,
	OrderByDirectionAsc:       orderByDirectionAscStr,
	OrderByDirectionDesc:      orderByDirectionDescStr,
	OrderByDirectionUsing:     orderByDirectionUsingStr,
}

// NewOrderByDirection creates a new instance of the enum from raw string.
func NewOrderByDirection(raw string) (OrderByDirection, error) {
	switch raw {
	case orderByDirectionAscStr:
		return OrderByDirectionAsc, nil
	case orderByDirectionDescStr:
		return OrderByDirectionDesc, nil
	case orderByDirectionUsingStr:
		return OrderByDirectionUsing, nil
	default:
		return OrderByDirection(Unknown), fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
}

// String returns the string representation of the enum type
func (r OrderByDirection) String() string {
	return toStringOrderByDirection[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *OrderByDirection) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case orderByDirectionAscStr:
		*r = OrderByDirectionAsc
	case orderByDirectionDescStr:
		*r = OrderByDirectionDesc
	case orderByDirectionUsingStr:
		*r = OrderByDirectionUsing
	default:
		*r = Unknown
		return fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r OrderByDirection) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringOrderByDirection[r]); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.Bytes(), nil
}
