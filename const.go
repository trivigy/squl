package squl

import (
	stdfmt "fmt"
	"strconv"

	"github.com/google/uuid"
	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// Const describes a constant primitive value.
type Const struct {
	Value interface{}
}

func (r *Const) dump(_ *ordinalMarker) (string, error) {
	switch val := r.Value.(type) {
	case uuid.UUID:
		return stdfmt.Sprintf("'%s'", val.String()), nil
	case string:
		return stdfmt.Sprintf("'%s'", val), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return stdfmt.Sprintf("%d", val), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type (%T) for %q", val, val))
	}
}
