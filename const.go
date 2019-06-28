package squl

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// Const describes a constant primitive value.
type Const struct {
	Value interface{}
}

func (r *Const) dump(counter *ordinalMarker) (string, error) {
	switch val := r.Value.(type) {
	case string:
		return fmt.Sprintf("'%s'", val), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	default:
		return "", errors.Errorf("unknown type (%T) for %q", val, val)
	}
}
