package squl

// Build parses the provided go template and produces a query string and
// arguments slice consumable by db.Exec().
func Build(cmd Command) (string, []interface{}, error) {
	return cmd.build()
}
