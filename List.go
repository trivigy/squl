package squl

import (
	"bytes"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
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
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		if i < len(*r)-1 {
			if _, err := buffer.WriteString(","); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
		}
	}
	return buffer.String(), nil
}
