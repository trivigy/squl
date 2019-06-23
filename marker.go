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

type marker struct {
	entries map[string]arg
}

func (r *marker) Mark(key string, value interface{}) string {
	entry, ok := r.entries[key]
	if !ok {
		entry = arg{index: len(r.entries) + 1, value: value}
		r.entries[key] = entry
	}
	return fmt.Sprintf("$%d", entry.index)
}

func (r *marker) GOB() (string, error) {
	entries := make(map[int]interface{})
	for _, entry := range r.entries {
		entries[entry.index] = entry.value
	}

	var keys []int
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
