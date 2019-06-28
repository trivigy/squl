package squl

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// Persistence describes the SELECT INTO persistence type.
type Persistence int

const (
	// PersistencePermanent indicates permanent table persistence.
	PersistencePermanent Persistence = iota + 1

	// PersistenceUnlogged indicates unlogged table persistence
	PersistenceUnlogged

	// PersistenceTemporary indicates temporary table persistence.
	PersistenceTemporary
)

const (
	persistencePermanentStr = "permanent"
	persistenceUnloggedStr  = "unlogged"
	persistenceTemporaryStr = "temporary"
)

var toStringPersistence = map[Persistence]string{
	Persistence(Unknown): unknownStr,
	PersistencePermanent: persistencePermanentStr,
	PersistenceUnlogged:  persistenceUnloggedStr,
	PersistenceTemporary: persistenceTemporaryStr,
}

// NewPersistence creates a new instance of the enum from raw string.
func NewPersistence(raw string) (Persistence, error) {
	switch raw {
	case persistencePermanentStr:
		return PersistencePermanent, nil
	case persistenceUnloggedStr:
		return PersistenceUnlogged, nil
	case persistenceTemporaryStr:
		return PersistenceTemporary, nil
	default:
		return Persistence(Unknown), errors.Errorf("unknown type %q", raw)
	}
}

// String returns the string representation of the enum type
func (r Persistence) String() string {
	return toStringPersistence[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *Persistence) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case persistencePermanentStr:
		*r = PersistencePermanent
	case persistenceUnloggedStr:
		*r = PersistenceUnlogged
	case persistenceTemporaryStr:
		*r = PersistenceTemporary
	default:
		*r = Unknown
		return errors.Errorf("unknown type %q", raw)
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r Persistence) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringPersistence[r]); err != nil {
		return nil, errors.WithStack(err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, errors.WithStack(err)
	}
	return buffer.Bytes(), nil
}
