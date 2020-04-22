package squl

import (
	"bytes"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// OnConflict representation of ON CONFLICT clause.
type OnConflict struct {
	Action      OnConflictAction `json:"action"`      /* DO NOTHING or UPDATE? */
	Infer       *Infer           `json:"infer"`       /* Optional index inference clause */
	TargetList  *List            `json:"targetList"`  /* the target list (of ResTarget) */
	WhereClause Node             `json:"whereClause"` /* qualifications */
	Location    int              `json:"location"`    /* token location, or -1 if unknown */
}

func (r *OnConflict) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if r.Infer != nil {
		dump, err := r.Infer.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	switch r.Action {
	case OnConflictNothing:
		if _, err := buffer.WriteString(" DO NOTHING"); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	case OnConflictUpdate:
		if _, err := buffer.WriteString(" DO UPDATE SET "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	default:
		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", r.Action))
	}

	if r.TargetList != nil {
		dump, err := r.TargetList.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	// if r.Schema != "" {
	// 	if _, err := buffer.WriteString(stdfmt.Sprintf("%s.", r.Schema)); err != nil {
	// 		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	// 	}
	// }
	// if r.Name == "" {
	// 	return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Name))
	// }
	// if _, err := buffer.WriteString(r.Name); err != nil {
	// 	return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	// }
	//
	// if r.Alias != "" {
	// 	if _, err := buffer.WriteString(stdfmt.Sprintf(" AS %s", r.Alias)); err != nil {
	// 		return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	// 	}
	// }
	return buffer.String(), nil
}
