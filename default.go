package squl

// Default describes a placeholder for DEFAULT constant value.
type Default struct{}

func (r *Default) dump(counter *ordinalMarker) (string, error) {
	return "DEFAULT", nil
}
