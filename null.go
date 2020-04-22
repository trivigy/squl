package squl

// Null defines a placehold for specifying NULL constant value.
type Null struct{}

func (r *Null) dump(_ *ordinalMarker) (string, error) {
	return "NULL", nil
}
