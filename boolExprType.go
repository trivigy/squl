package squl

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// BoolExprType describes the types of boolean expressions available.
type BoolExprType int

const (
	// BoolExprTypeAnd describes the AND expression.
	BoolExprTypeAnd BoolExprType = iota + 1

	// BoolExprTypeOr describes the OR expression.
	BoolExprTypeOr
)

const (
	boolExprTypeAndStr = "and"
	boolExprTypeOrStr  = "or"
)

var toStringBoolExprType = map[BoolExprType]string{
	BoolExprType(Unknown): unknownStr,
	BoolExprTypeAnd:       boolExprTypeAndStr,
	BoolExprTypeOr:        boolExprTypeOrStr,
}

// NewBoolExprType creates a new instance of the enum from raw string.
func NewBoolExprType(raw string) (BoolExprType, error) {
	switch raw {
	case boolExprTypeAndStr:
		return BoolExprTypeAnd, nil
	case boolExprTypeOrStr:
		return BoolExprTypeOr, nil
	default:
		return BoolExprType(Unknown), errors.Errorf("unknown type %q", raw)
	}
}

// String returns the string representation of the enum type
func (r BoolExprType) String() string {
	return toStringBoolExprType[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *BoolExprType) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case boolExprTypeAndStr:
		*r = BoolExprTypeAnd
	case boolExprTypeOrStr:
		*r = BoolExprTypeOr
	default:
		*r = Unknown
		return errors.Errorf("unknown type %q", raw)
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r BoolExprType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringBoolExprType[r]); err != nil {
		return nil, errors.WithStack(err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, errors.WithStack(err)
	}
	return buffer.Bytes(), nil
}
