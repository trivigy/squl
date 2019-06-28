package squl

// Command describes top level commands like INSERT, SELECT, UPDATE, DELETE, etc.
type Command interface {
	build() (string, []interface{}, error)
}
