package squl

import (
	"bytes"
	stdfmt "fmt"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// Insert defines the struct for the INSERT command.
type Insert struct {
	// With represents WITH clause.
	// With *With `json:"with"`

	// Relation represents the relation to insert into.
	Relation *RangeVar `json:"relation"`

	// Columns represents names of target columns (optional).
	Columns *List `json:"columns"`

	// Values represents the SELECT, VALUES, or NULL.
	Select *Select `json:"select"`

	// OnConflict represents the conflict resolution policy.
	OnConflict *OnConflict `json:"onConflict"`

	// Returning represents a list of expressions to return.
	Returning Node `json:"returning"`
}

func (r *Insert) build() (string, []interface{}, error) {
	counter := &ordinalMarker{}
	query, err := r.dump(counter)
	return query, counter.args(), err
}

func (r *Insert) dump(counter *ordinalMarker) (string, error) {
	buffer := bytes.NewBuffer([]byte("INSERT INTO"))
	if r.Relation != nil {
		if _, err := buffer.WriteString(" "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Relation.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Columns != nil && len(*r.Columns) > 0 {
		dump, err := r.Columns.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(stdfmt.Sprintf(" (%s)", dump)); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Select == nil {
		if r.Columns != nil && len(*r.Columns) > 0 {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Select))
		}
		if _, err := buffer.WriteString(" DEFAULT VALUES"); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	} else {
		if !r.Select.isSelectStmt() && !r.Select.isValuesList() {
			if r.Columns != nil && len(*r.Columns) > 0 {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("required parameter %#v", r.Select))
			}
			if _, err := buffer.WriteString(" DEFAULT VALUES"); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
		} else if r.Select.isSelectStmt() && !r.Select.isValuesList() {
			if _, err := buffer.WriteString(" "); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
			dump, err := r.Select.dump(counter)
			if err != nil {
				return "", err
			}
			if _, err := buffer.WriteString(dump); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
		} else if !r.Select.isSelectStmt() && r.Select.isValuesList() {
			if _, err := buffer.WriteString(" VALUES "); err != nil {
				return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
			}
			for i, list := range r.Select.Values {
				dump, err := list.dump(counter)
				if err != nil {
					return "", err
				}
				if _, err := buffer.WriteString(stdfmt.Sprintf("(%s)", dump)); err != nil {
					return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
				}
				if i < len(r.Select.Values)-1 {
					if _, err := buffer.WriteString(","); err != nil {
						return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
					}
				}
			}
		} else {
			return "", fmt.Errorf("")
		}
	}

	if r.OnConflict != nil {
		if _, err := buffer.WriteString(" ON CONFLICT "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.OnConflict.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}

	if r.Returning != nil {
		if _, err := buffer.WriteString(" RETURNING "); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
		dump, err := r.Returning.dump(counter)
		if err != nil {
			return "", err
		}
		if _, err := buffer.WriteString(dump); err != nil {
			return "", fmt.Errorf(global.ErrFmt, pkg.Name(), err)
		}
	}
	return buffer.String(), nil
}

// func (r *Insert) build() (string, []interface{}, error) {
// 	buffer := bytes.NewBuffer([]byte("INSERT INTO"))
//
// 	if r.Relation.Schema != "" {
// 		buffer.WriteString(fmt.Sprintf(" %s.", r.Relation.Schema))
// 	}
//
// 	if r.Relation.Name == "" {
// 		return "", nil, fmt.Errorf("required parameter `Relation.Alias`")
// 	}
//
// 	if r.Relation.Schema == "" {
// 		buffer.WriteString(" ")
// 	}
// 	buffer.WriteString(fmt.Sprintf("%s", r.Relation.Name))
//
// 	if r.Relation.Alias != "" {
// 		buffer.WriteString(fmt.Sprintf(" AS %s", r.Relation.Alias))
// 	}
//
// 	if len(r.Columns) > 0 {
// 		buffer.WriteString(" (")
// 		for i, col := range r.Columns {
// 			buffer.WriteString(col.Alias)
// 			if i < len(r.Columns)-1 {
// 				buffer.WriteString(",")
// 			}
// 		}
// 		buffer.WriteString(")")
// 	}
//
// 	if r.Values.isEmpty() || (!r.Values.isValuesList() && !r.Values.isSelectStmt()) {
// 		buffer.WriteString(" DEFAULT VALUES")
// 	} else if r.Values.isValuesList() {
// 		// r.Values.command = r.command
// 		buffer.WriteString(" VALUES")
// 		for i := range r.Values.Values {
// 			buffer.WriteString(fmt.Sprintf(" (%s)", r.Values.buildValuesRow(i)))
// 			if i < len(r.Values.Values)-1 {
// 				buffer.WriteString(",")
// 			}
// 		}
// 	} else {
// 		panic("not implemented")
// 	}
//
// 	if len(r.Returning) > 0 {
// 		buffer.WriteString(" RETURNING")
// 		for i, exp := range r.Returning {
// 			// switch value := exp.Value.(type) {
// 			// case ColumnRef:
// 			// 	if len(value.Fields) > 0 {
// 			// 		buffer.WriteString(fmt.Sprintf(" %s", value.Fields[0]))
// 			// 	}
// 			// default:
// 			// 	return "", nil, fmt.Errorf("unknown type %q", exp.Value)
// 			// }
//
// 			if exp.Alias != "" {
// 				buffer.WriteString(fmt.Sprintf(" AS %s", exp.Alias))
// 			}
//
// 			if i < len(r.Returning)-1 {
// 				buffer.WriteString(",")
// 			}
// 		}
// 	}
//
// 	return buffer.String(), nil, nil
// }
