package squl

import (
	"bytes"
	"encoding/json"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// JoinType describes the types of joins available.
type JoinType int

const (
	// JoinTypeDefault indicates the usage of JOIN.
	JoinTypeDefault JoinType = iota + 1

	// JoinTypeInner indicates the usage of INNER JOIN.
	JoinTypeInner

	// JoinTypeLeft indicates the usage of LEFT JOIN.
	JoinTypeLeft

	// JoinTypeOuterLeft indicates the usage of LEFT OUTER JOIN.
	JoinTypeOuterLeft

	// JoinTypeRight indicates the usage of RIGHT JOIN.
	JoinTypeRight

	// JoinTypeOuterRight indicates the usage of RIGHT OUTER JOIN.
	JoinTypeOuterRight

	// JoinTypeFull indicates the usage of FULL JOIN.
	JoinTypeFull

	// JoinTypeOuterFull indicates the usage of FULL OUTER JOIN.
	JoinTypeOuterFull

	// JoinTypeCross indicates the usage of CROSS JOIN.
	JoinTypeCross
)

const (
	joinTypeDefaultStr    = "default"
	joinTypeInnerStr      = "inner"
	joinTypeLeftStr       = "left"
	joinTypeOuterLeftStr  = "outerLeft"
	joinTypeRightStr      = "right"
	joinTypeOuterRightStr = "outerRight"
	joinTypeFullStr       = "full"
	joinTypeOuterFullStr  = "outerFull"
	joinTypeCrossStr      = "cross"
)

var toStringJoinType = map[JoinType]string{
	JoinType(Unknown):  unknownStr,
	JoinTypeDefault:    joinTypeDefaultStr,
	JoinTypeInner:      joinTypeInnerStr,
	JoinTypeLeft:       joinTypeLeftStr,
	JoinTypeOuterLeft:  joinTypeOuterLeftStr,
	JoinTypeRight:      joinTypeRightStr,
	JoinTypeOuterRight: joinTypeOuterRightStr,
	JoinTypeFull:       joinTypeFullStr,
	JoinTypeOuterFull:  joinTypeOuterFullStr,
	JoinTypeCross:      joinTypeCrossStr,
}

// NewJoinType creates a new instance of the enum from raw string.
func NewJoinType(raw string) (JoinType, error) {
	switch raw {
	case joinTypeDefaultStr:
		return JoinTypeDefault, nil
	case joinTypeInnerStr:
		return JoinTypeInner, nil
	case joinTypeLeftStr:
		return JoinTypeLeft, nil
	case joinTypeOuterLeftStr:
		return JoinTypeOuterLeft, nil
	case joinTypeRightStr:
		return JoinTypeRight, nil
	case joinTypeOuterRightStr:
		return JoinTypeOuterRight, nil
	case joinTypeFullStr:
		return JoinTypeFull, nil
	case joinTypeOuterFullStr:
		return JoinTypeOuterFull, nil
	case joinTypeCrossStr:
		return JoinTypeCross, nil
	default:
		return JoinType(Unknown), fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
}

// String returns the string representation of the enum type
func (r JoinType) String() string {
	return toStringJoinType[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *JoinType) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	switch s {
	case joinTypeDefaultStr:
		*r = JoinTypeDefault
	case joinTypeInnerStr:
		*r = JoinTypeInner
	case joinTypeLeftStr:
		*r = JoinTypeLeft
	case joinTypeOuterLeftStr:
		*r = JoinTypeOuterLeft
	case joinTypeRightStr:
		*r = JoinTypeRight
	case joinTypeOuterRightStr:
		*r = JoinTypeOuterRight
	case joinTypeFullStr:
		*r = JoinTypeFull
	case joinTypeOuterFullStr:
		*r = JoinTypeOuterFull
	case joinTypeCrossStr:
		*r = JoinTypeCross
	default:
		*r = Unknown
		return fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", s))
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r JoinType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringJoinType[r]); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.Bytes(), nil
}
