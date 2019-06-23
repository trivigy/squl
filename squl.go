package squl

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"
	"sync"
	"text/template"

	"github.com/pkg/errors"
)

// Builder defines the primary concurrency safe sql query builder object.
type Builder struct {
	cache sync.Map
	ctx   context
}

func (r *Builder) context() *context {
	return &r.ctx
}

// Build parses the provided go template and produces a query string and
// arguments slice consumable by db.Exec().
func (r *Builder) Build(text string, data interface{}) (string, []interface{}, error) {
	var driver *template.Template
	if value, ok := r.cache.Load(text); !ok {
		var err error
		driver, err = template.New("main").
			Funcs(template.FuncMap{"squl": r.context}).
			Parse(fmt.Sprintf(`{{$ := squl.Open}}%s;{{$.Params.Args}}{{squl.Close $}}`, text))
		if err != nil {
			return "", nil, errors.WithStack(err)
		}

		r.cache.Store(text, driver)
	} else {
		driver = value.(*template.Template)
	}

	tpl, err := driver.Clone()
	if err != nil {
		return "", nil, errors.WithStack(err)
	}

	buffer := bytes.NewBuffer(nil)
	if err := tpl.Execute(buffer, data); err != nil {
		return "", nil, errors.WithStack(err)
	}

	delim := strings.LastIndex(buffer.String(), ";")
	query := buffer.String()[0:delim]
	decoded, err := base64.StdEncoding.DecodeString(buffer.String()[delim+1:])
	if err != nil {
		return "", nil, errors.WithStack(err)
	}

	args := make([]interface{}, 0)
	buffer = bytes.NewBuffer(decoded)
	if err = gob.NewDecoder(buffer).Decode(&args); err != nil {
		return "", nil, errors.WithStack(err)
	}
	return query, args, nil
}
