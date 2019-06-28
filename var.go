package squl

// Var defines an object for marking parameters in the query. Values which are
// placed inside of Var are ultimately replaced with `$%d` enumerated number
// and the argument is returned in the appropriate position when the quory is
// done building.
type Var struct {
	Value interface{}
}

func (r *Var) dump(counter *ordinalMarker) (string, error) {
	return counter.mark(r.Value), nil
}
