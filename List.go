package squl

import (
	"bytes"

	"github.com/pkg/errors"
)

// List describes a collection of nodes aggregated together.
type List []Node

func (r *List) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	for i, each := range *r {
		eachDump, err := each.dump(counter)
		if err != nil {
			return "", err
		}

		if _, err := buffer.WriteString(eachDump); err != nil {
			return "", errors.WithStack(err)
		}
		if i < len(*r)-1 {
			if _, err := buffer.WriteString(","); err != nil {
				return "", errors.WithStack(err)
			}
		}
	}
	return buffer.String(), nil
}
