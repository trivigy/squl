package squl

// Node describes the abstract element which can be built.
type Node interface {
	dump(counter *ordinalMarker) (string, error)
}
