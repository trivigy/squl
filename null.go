package squl

// Null defines a placehold for specifying NULL constant value.
type Null struct{}

func (r *Null) dump(counter *ordinalMarker) (string, error) {
	return "NULL", nil
}
