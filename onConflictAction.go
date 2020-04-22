package squl

import (
	"bytes"
	"encoding/json"
	"strings"

	errors "golang.org/x/xerrors"
	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// OnConflictAction represents the action an ON CONFLICT clause should take.
type OnConflictAction int

const (
	// OnConflictNothing indicates that the conflict resolution policy should
	// do nothing to resolve conflict.
	OnConflictNothing OnConflictAction = iota + 1

	// OnConflictUpdate indicates that the conflict resolution policy should
	// update the existing value.
	OnConflictUpdate
)

const (
	onConflictNothingStr = "nothing"
	onConflictUpdateStr  = "update"
)

var toStringOnConflictAction = map[OnConflictAction]string{
	OnConflictAction(Unknown): unknownStr,
	OnConflictNothing:         onConflictNothingStr,
	OnConflictUpdate:          onConflictUpdateStr,
}

// NewOnConflictAction creates a new instance of the enum from raw string.
func NewOnConflictAction(raw string) (OnConflictAction, error) {
	switch raw {
	case onConflictNothingStr:
		return OnConflictNothing, nil
	case onConflictUpdateStr:
		return OnConflictUpdate, nil
	default:
		return OnConflictAction(Unknown), fmt.Errorf(global.ErrFmt, pkg.Name(), errors.Errorf("unknown type %q", raw))
	}
}

// String returns the string representation of the enum type
func (r OnConflictAction) String() string {
	return toStringOnConflictAction[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *OnConflictAction) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case onConflictNothingStr:
		*r = OnConflictNothing
	case onConflictUpdateStr:
		*r = OnConflictUpdate
	default:
		*r = Unknown
		return fmt.Errorf(global.ErrFmt, pkg.Name(), errors.Errorf("unknown type %q", raw))
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r OnConflictAction) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringOnConflictAction[r]); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.Bytes(), nil
}
