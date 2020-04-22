package squl

// Into describes the INTO clause of SELELCT INTO command.
type Into struct {
	Relation *RangeVar `json:"relation"` /* target relation name */
	Columns  List      `json:"columns"`  /* column names to assign, or NIL */
	// Options        List           `json:"options"`        /* options from WITH clause */
	// OnCommit       OnCommitAction `json:"onCommit"`       /* what do we do at COMMIT? */
	// TableSpaceName *string        `json:"tableSpaceName"` /* table space to use, or NULL */
	// ViewQuery      Node           `json:"viewQuery"`      /* materialized view's SELECT query */
	// SkipData       bool           `json:"skipData"`       /* true for WITH NO DATA */
}

func (r *Into) dump(_ *ordinalMarker) (string, error) {
	return "", nil
}
