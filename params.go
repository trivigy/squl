package squl

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

type arg struct {
	index int
	value interface{}
}

// Params object for generating ordinal parameter markers for the provided
// parameters.
type Params struct {
	entries map[string]arg
}

// Mark is used to mark a specific key and associated value as a parameter
func (r *Params) Mark(key string, value interface{}) string {
	entry, ok := r.entries[key]
	if !ok {
		entry = arg{index: len(r.entries) + 1, value: value}
		r.entries[key] = entry
	}
	return fmt.Sprintf("$%d", entry.index)
}

// Args prints a base64 encoded encoding/gob version of the parameters. This
// exists for internal usage so I really wouldn't recommend using it unless you
// know what you are doing and need it for some reason.
func (r *Params) Args() (string, error) {
	entries := make(map[int]interface{})
	for _, entry := range r.entries {
		entries[entry.index] = entry.value
	}

	keys := make([]int, 0)
	for key := range entries {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	args := make([]interface{}, 0)
	for _, key := range keys {
		args = append(args, entries[key])
	}

	buffer := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buffer).Encode(args); err != nil {
		return "", errors.WithStack(err)
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
